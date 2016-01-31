package genmarkdown_test

import (
	. "github.com/garslo/genmarkdown"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Paragraph", func() {
	var (
		paragraph Paragraph
	)

	BeforeEach(func() {
		paragraph = Paragraph{
			Text: "Some text that may span multiple lines",
		}
	})

	It("should be a node", func() {
		var _ Node = paragraph
	})

	It("should generate the correct markdown", func() {
		markdown := Slurp(paragraph.Markdown())
		Expect(markdown).To(Equal("Some text that may span multiple lines\n\n"))
	})

	It("should have a maximum line length of 80", func() {
		// 90 chars long
		paragraph.Text = "xx xx xx xxx xx xx xxx xx xx xxx xx xx xxx xx xx xxx xx xx xxx xx xx xxx xx xx xxx xx xx x"
		expected := "xx xx xx xxx xx xx xxx xx xx xxx xx xx xxx xx xx xxx xx xx xxx xx xx xxx xx xx\nxxx xx xx x\n\n"
		markdown := Slurp(paragraph.Markdown())
		Expect(markdown).To(Equal(expected))
	})
})
