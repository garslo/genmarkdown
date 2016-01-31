package genmarkdown_test

import (
	. "github.com/garslo/genmarkdown"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CodeBlock", func() {
	var (
		codeBlock CodeBlock
	)

	BeforeEach(func() {
		codeBlock = CodeBlock{
			Language: "clojure",
			Code: `(let [x 4]
  (print x))`,
		}
	})

	It("should be a node", func() {
		var _ Node = codeBlock
	})

	It("should generate the correct markdown", func() {
		markdown := Slurp(codeBlock.Markdown())
		expected := "```clojure\n(let [x 4]\n  (print x))\n```\n\n"
		Expect(markdown).To(Equal(expected))
	})

	It("should not print the language if none is given", func() {
		codeBlock.Language = ""
		markdown := Slurp(codeBlock.Markdown())
		expected := "```\n(let [x 4]\n  (print x))\n```\n\n"
		Expect(markdown).To(Equal(expected))
	})
})
