package genmarkdown_test

import (
	. "github.com/garslo/genmarkdown"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Heading", func() {
	var (
		heading Heading
	)

	BeforeEach(func() {
		heading = Heading{
			Text:  "Some Title",
			Level: 1,
		}
	})

	It("should be a node", func() {
		var _ Node = heading
	})

	It("should generate the correct markdown", func() {
		markdown := Slurp(heading.Markdown())
		Expect(markdown).To(Equal("# Some Title\n\n"))
	})

	It("should have the correct heading level", func() {
		heading.Level = 3
		markdown := Slurp(heading.Markdown())
		Expect(markdown).To(HavePrefix("### "))
	})
})
