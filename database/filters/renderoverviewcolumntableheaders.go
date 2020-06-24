package filters

import (
	"strings"
)

// RenderOverviewColumnTableHeaders renders the HTML output for the internal headers.
func (filter *filter) RenderOverviewColumnTableHeaders(ctx Context) string {
	out := strings.Builder{}
	for _, field := range filter.ColumnsInOrder() {
		filter.perror(out.WriteString(`<th>`))
		if field.Plural() {
			filter.perror(out.WriteString(ctx.TranslatePlural(field.Name())))
		} else {
			filter.perror(out.WriteString(ctx.Translate(field.Name())))
		}
		filter.perror(out.WriteString(`</th>`))
	}
	return out.String()
}
