package sensor

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSensot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sensor")
}
