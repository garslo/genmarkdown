package genmarkdown_test

import (
	. "github.com/garslo/genmarkdown"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Table", func() {
	var (
		table Table
	)

	BeforeEach(func() {
		table = Table{
			Columns: []string{"Col1", "Col2", "Col3"},
			Rows: []Row{
				Row{[]string{"x11", "x12", "x13"}},
				Row{[]string{"x21", "x22", "x23"}},
				Row{[]string{"x31", "x32", "x33"}},
				Row{[]string{"x41", "x42", "x43 longer than the rest"}},
			},
		}
	})

	It("should be a node", func() {
		var _ Node = table
	})

	It("should generate the correct markdown", func() {
		markdown := Slurp(table.Markdown())
		expected := `| Col1 | Col2 | Col3                     |
| ---  | ---  | ---                      |
| x11  | x12  | x13                      |
| x21  | x22  | x23                      |
| x31  | x32  | x33                      |
| x41  | x42  | x43 longer than the rest |

`
		Expect(markdown).To(Equal(expected))
	})
})
