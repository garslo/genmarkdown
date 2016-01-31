package genmarkdown_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGenmarkdown(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Genmarkdown Suite")
}
