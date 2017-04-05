package gallifrey_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGallifrey(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gallifrey Suite")
}
