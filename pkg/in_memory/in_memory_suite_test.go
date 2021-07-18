package in_memory_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInMemory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "InMemory Suite")
}
