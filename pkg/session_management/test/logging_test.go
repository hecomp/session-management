package test_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)


func TestLogging(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logging Suite")
}

var _ = Describe("Logging", func() {

})
