package builder

import (
	"fmt"
	"strings"
)

//define the indent size
const (
	indentSize = 2
)

type HTMLElement struct {
	name, text string
	element    []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.String()
}

// func NewHtmlBuilder(rootName string) *HTMLElement{
// 	return &HTMLElement{rootName, HTMLElement{
// 		rootName, "", []HTMLElement{}
// 	}}
// }

var sb = strings.Builder{}

func BuildUnorderList() {

	words := []string{"hello", "providus", "bank"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<ul>")
		sb.WriteString(v)
		sb.WriteString("</ul>")
	}
	sb.WriteString("<ul>")

	fmt.Println(sb.String())
}
