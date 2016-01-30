# genmarkdown

Generate markdown from structured data in Go.

# Example

Code:

```go
package main

import "github.com/garslo/genmarkdown"

func main() {
	tbl := genmarkdown.Table{
		Columns: []string{"Col1", "Col2"},
		Rows: []genmarkdown.Row{
			genmarkdown.Row{[]string{"row11", "row12"}},
			genmarkdown.Row{[]string{"row21", "row22"}},
		},
	}
	e := genmarkdown.Emitter{
		Nodes: genmarkdown.Nodes{
			genmarkdown.Heading{
				Text:  "Heading 1",
				Level: 1,
			},
			genmarkdown.Paragraph{
				Text: "This is some text in a paragraph.",
			},
			genmarkdown.Heading{
				Text:  "Heading 2",
				Level: 2,
			},
			genmarkdown.CodeBlock{
				Language: "go",
				Code:     "fmt.Sprintf()",
			},
			genmarkdown.Heading{
				Text:  "Heading 4",
				Level: 4,
			},
			tbl,
		},
	}
	e.Emit()
}
```

Output:

<pre>```
# Heading 1

This is some text in a paragraph.

## Heading 2

```go
fmt.Sprintf()
```

#### Heading 4

| Col1  | Col2  |
| ---   | ---   |
| row11 | row12 |
| row21 | row22 |
```</pre>
