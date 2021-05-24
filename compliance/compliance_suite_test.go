package compliance_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCompliance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Compliance Suite")
}
