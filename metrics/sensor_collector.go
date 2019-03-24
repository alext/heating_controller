package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

func newDensorDesc() *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName("house", "", "temperature_celcius"),
		"Current temperature in degrees Celcius",
		[]string{"name", "device_id"},
		nil,
	)
}

func (m *Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.sensorDesc
}

func (m *Metrics) Collect(ch chan<- prometheus.Metric) {
	for name, s := range m.ctrl.SensorsByName {
		temp, ts := s.Read()

		metric, err := prometheus.NewConstMetric(m.sensorDesc, prometheus.GaugeValue, temp.Float(), name, s.DeviceId())
		if err != nil {
			log.Printf("[metrics] Error constructing sensor metric for %s: %s", name, err.Error())
			continue
		}
		metric = prometheus.NewMetricWithTimestamp(ts, metric)
		ch <- metric
	}
}
