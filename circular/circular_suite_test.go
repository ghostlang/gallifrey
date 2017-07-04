package circular_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCircular(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Circular Suite")
}
