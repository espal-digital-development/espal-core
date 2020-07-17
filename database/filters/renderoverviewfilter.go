package filters

import (
	"io"
	"strconv"
	"strings"
)

// RenderOverviewFilter will render the filter actions of an admin module overview page.
func (f *filter) RenderOverviewFilter(ctx Context) string {
	out := &strings.Builder{}

	out.WriteString(`<div class="filter">`)

	out.WriteString(`<div class="filterInfo">`)
	out.WriteString(strconv.FormatUint(uint64(f.totalResults), 10))
	out.WriteString(` `)
	if f.totalResults == 1 {
		out.WriteString(ctx.Translate("result"))
	} else {
		out.WriteString(ctx.TranslatePlural("result"))
	}
	out.WriteString(`</div>`)

	out.WriteString(`<div>`)
	if f.ShouldShowSearch() {
		out.WriteString(`<input id="filterSearch" placeholder="`)
		out.WriteString(ctx.Translate("search"))
		out.WriteString(`"`)
		if f.search != "" {
			out.WriteString(`value="`)
			out.WriteString(f.search)
			out.WriteString(`"`)
		}
		out.WriteString(`>`)
	}
	out.WriteString(`<div class="filters"></div>`)
	out.WriteString(`</div>`)

	if f.totalPages > 1 {
		f.renderPagination(out)
	}

	out.WriteString(`</div>`)

	return out.String()
}

func (f *filter) renderPagination(out io.StringWriter) {
	out.WriteString(`<div class="pagination disable-select">`)

	out.WriteString(`<select>`)
	for _, r := range []uint{5, 10, 25, 50, 100} {
		out.WriteString(`<option value="`)
		out.WriteString(strconv.FormatUint(uint64(r), 10))
		out.WriteString(`"`)
		if r == f.limit {
			out.WriteString(` selected`)
		}
		out.WriteString(`>`)
		out.WriteString(strconv.FormatUint(uint64(r), 10))
		out.WriteString(`</option>`)
	}
	out.WriteString(`</select>`)

	for _, page := range f.PaginationBlocks() {
		out.WriteString(`<p `)
		if page == 0 || f.CurrentPage() == page {
			out.WriteString(`class="`)
			if f.CurrentPage() == page {
				out.WriteString(`current`)
			} else if page == 0 {
				out.WriteString(`separator`)
			}
			out.WriteString(`"`)
		}
		out.WriteString(`>`)
		if page == 0 {
			out.WriteString(`...`)
		} else {
			out.WriteString(strconv.FormatUint(uint64(page), 10))
		}
		out.WriteString(`</p>`)
	}

	out.WriteString(`</div>`)
}
