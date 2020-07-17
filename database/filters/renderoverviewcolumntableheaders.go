package filters

import (
	"strings"
)

// RenderOverviewColumnTableHeaders renders the HTML output for the internal headers.
func (f *filter) RenderOverviewColumnTableHeaders(ctx Context) string {
	out := strings.Builder{}
	for _, field := range f.ColumnsInOrder() {
		out.WriteString(`<th>`)
		if field.Plural() {
			out.WriteString(ctx.TranslatePlural(field.Name()))
		} else {
			out.WriteString(ctx.Translate(field.Name()))
		}
		out.WriteString(`</th>`)
	}
	return out.String()
}
