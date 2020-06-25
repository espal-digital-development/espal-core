package filters

import (
	"strings"
)

// RenderOverviewColumnTableHeaders renders the HTML output for the internal headers.
func (f *filter) RenderOverviewColumnTableHeaders(ctx Context) string {
	out := strings.Builder{}
	for _, field := range f.ColumnsInOrder() {
		f.perror(out.WriteString(`<th>`))
		if field.Plural() {
			f.perror(out.WriteString(ctx.TranslatePlural(field.Name())))
		} else {
			f.perror(out.WriteString(ctx.Translate(field.Name())))
		}
		f.perror(out.WriteString(`</th>`))
	}
	return out.String()
}
