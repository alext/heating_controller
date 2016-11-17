package thermostat

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestThermostat(t *testing.T) {
	RegisterFailHandler(Fail)

	log.SetOutput(ioutil.Discard)

	RunSpecs(t, "Thermostat")
}

func temperatureHandler(value uint) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{"temperature": value}
		Expect(json.NewEncoder(w).Encode(&data)).To(Succeed())
	})
}

var _ = Describe("A Thermostat", func() {
	var (
		t      *thermostat
		server *httptest.Server
	)
	AfterEach(func() {
		if t != nil {
			t.Close()
		}
		if server != nil {
			server.Close()
		}
	})

	Describe("constructing a thermostat", func() {
		BeforeEach(func() {
			server = httptest.NewServer(temperatureHandler(19000))
		})

		It("builds one correctly", func() {
			t = New("something", server.URL, 19000, func(b bool) {}).(*thermostat)
			Expect(t.id).To(Equal("something"))
			Expect(t.url).To(Equal(server.URL))
			Expect(t.target).To(BeNumerically("==", 19000))
		})

		It("starts as active when current temp is within the threshold", func() {
			t = New("something", server.URL, 19000, func(b bool) {}).(*thermostat)
			Expect(t.active).To(BeTrue())
		})
	})

	Describe("setting the target temperature", func() {
		BeforeEach(func() {
			t = &thermostat{
				current: 18000,
				target:  17000,
			}
		})

		It("Updates the target for the thermostat", func() {
			t.Set(16000)
			Expect(t.target).To(BeNumerically("==", 16000))
		})

		It("triggers the thermostat to update the demand state", func() {
			t.Set(19000)
			Expect(t.active).To(BeTrue())
		})
	})

	Describe("reading the temperature from the source", func() {

		Context("happy path", func() {
			BeforeEach(func() {
				server = httptest.NewServer(temperatureHandler(19000))
				t = &thermostat{
					url:     server.URL,
					current: 17000,
					target:  18000,
					active:  true,
				}
			})

			It("updates the current temperature", func() {
				t.readTemp()
				Expect(t.current).To(BeNumerically("==", 19000))
			})

			It("triggers the thermostat to update the demand state", func() {
				t.readTemp()
				Expect(t.active).To(BeFalse())
			})
		})

		Describe("error handling", func() {
			BeforeEach(func() {
				server = httptest.NewServer(temperatureHandler(19000))
				t = &thermostat{
					url:     server.URL,
					current: 18000,
				}
			})

			Context("when the network connection fails", func() {
				It("doesn't update the current temperature", func() {
					server.Close()
					t.readTemp()
					Expect(t.current).To(BeNumerically("==", 18000))
				})
			})

			Context("when the HTTP returns non-200 response", func() {
				It("doesn't update the current temperature", func() {
					server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.WriteHeader(http.StatusInternalServerError)
						data := map[string]interface{}{"temperature": 19000}
						Expect(json.NewEncoder(w).Encode(&data)).To(Succeed())
					})
					t.readTemp()
					Expect(t.current).To(BeNumerically("==", 18000))
				})
			})

			Context("when the response isn't JSON", func() {
				It("doesn't update the current temperature", func() {
					server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.Write([]byte("I'm not JSON..."))
					})
					t.readTemp()
					Expect(t.current).To(BeNumerically("==", 18000))
				})
			})

			Context("when the response does not include a temperature field", func() {
				It("doesn't update the current temperature", func() {
					server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						data := map[string]interface{}{"something": "else"}
						Expect(json.NewEncoder(w).Encode(&data)).To(Succeed())
					})
					t.readTemp()
					Expect(t.current).To(BeNumerically("==", 18000))
				})
			})
		})
	})

	type TriggeringCase struct {
		Current            Temperature
		Target             Temperature
		CurrentlyActive    bool
		ExpectedActive     bool
		ExpectDemandCalled bool
	}

	DescribeTable("triggering changes in state",
		func(c TriggeringCase) {
			demandCalled := false
			demandNotify := make(chan struct{})
			t := &thermostat{
				current: c.Current,
				target:  c.Target,
				active:  c.CurrentlyActive,
				demand: func(param bool) {
					defer GinkgoRecover()
					demandCalled = true
					Expect(param).To(Equal(c.ExpectedActive))
					close(demandNotify)
				},
			}

			t.trigger()
			Expect(t.active).To(Equal(c.ExpectedActive))
			if c.ExpectDemandCalled {
				select {
				case <-demandNotify:
				case <-time.After(time.Second):
				}
				Expect(demandCalled).To(BeTrue(), "expected demandFunc to be called")
			}
		},
		Entry("activates when current well below target", TriggeringCase{
			Current: 15000, Target: 18000, CurrentlyActive: false,
			ExpectedActive: true, ExpectDemandCalled: true,
		}),
		Entry("remains active when current well below target", TriggeringCase{
			Current: 15000, Target: 18000, CurrentlyActive: true,
			ExpectedActive: true, ExpectDemandCalled: false,
		}),
		Entry("deactivates when current well above target", TriggeringCase{
			Current: 20000, Target: 18000, CurrentlyActive: true,
			ExpectedActive: false, ExpectDemandCalled: true,
		}),
		Entry("deactivates when current slightly above target", TriggeringCase{
			Current: 18050, Target: 18000, CurrentlyActive: true,
			ExpectedActive: false, ExpectDemandCalled: true,
		}),
		Entry("remains inactive when current well above target", TriggeringCase{
			Current: 20000, Target: 18000, CurrentlyActive: false,
			ExpectedActive: false, ExpectDemandCalled: false,
		}),
		Entry("remains active when current within threhold below target", TriggeringCase{
			Current: 17950, Target: 18000, CurrentlyActive: true,
			ExpectedActive: true, ExpectDemandCalled: false,
		}),
		Entry("remains inactive when current within threhold below target", TriggeringCase{
			Current: 17950, Target: 18000, CurrentlyActive: false,
			ExpectedActive: false, ExpectDemandCalled: false,
		}),
		Entry("remains inactive when current slightly above target", TriggeringCase{
			Current: 18050, Target: 18000, CurrentlyActive: false,
			ExpectedActive: false, ExpectDemandCalled: false,
		}),
	)
})