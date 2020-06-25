package filters

import (
	"strconv"
	"strings"
)

// RenderOverviewFilter will render the filter actions of an
// admin module overview page.
func (f *filter) RenderOverviewFilter(ctx Context) string {
	out := strings.Builder{}

	f.perror(out.WriteString(`<div class="filter">`))

	f.perror(out.WriteString(`<div class="filterInfo">`))
	f.perror(out.WriteString(strconv.FormatUint(uint64(f.totalResults), 10)))
	f.perror(out.WriteString(` `))
	if f.totalResults == 1 {
		f.perror(out.WriteString(ctx.Translate("result")))
	} else {
		f.perror(out.WriteString(ctx.TranslatePlural("result")))
	}
	f.perror(out.WriteString(`</div>`))

	f.perror(out.WriteString(`<div>`))
	if f.ShouldShowSearch() {
		f.perror(out.WriteString(`<input id="filterSearch" placeholder="`))
		f.perror(out.WriteString(ctx.Translate("search")))
		f.perror(out.WriteString(`"`))
		if f.search != "" {
			f.perror(out.WriteString(`value="`))
			f.perror(out.WriteString(f.search))
			f.perror(out.WriteString(`"`))
		}
		f.perror(out.WriteString(`>`))
	}
	f.perror(out.WriteString(`<div class="filters"></div>`))
	f.perror(out.WriteString(`</div>`))

	if f.totalPages > 1 {
		f.perror(out.WriteString(`<div class="pagination disable-select">`))

		f.perror(out.WriteString(`<select>`))
		for _, r := range []uint{5, 10, 25, 50, 100} {
			f.perror(out.WriteString(`<option value="`))
			f.perror(out.WriteString(strconv.FormatUint(uint64(r), 10)))
			f.perror(out.WriteString(`"`))
			if r == f.limit {
				f.perror(out.WriteString(` selected`))
			}
			f.perror(out.WriteString(`>`))
			f.perror(out.WriteString(strconv.FormatUint(uint64(r), 10)))
			f.perror(out.WriteString(`</option>`))
		}
		f.perror(out.WriteString(`</select>`))

		for _, page := range f.PaginationBlocks() {
			f.perror(out.WriteString(`<p `))
			if page == 0 || f.CurrentPage() == page {
				f.perror(out.WriteString(`class="`))
				if f.CurrentPage() == page {
					f.perror(out.WriteString(`current`))
				} else if page == 0 {
					f.perror(out.WriteString(`separator`))
				}
				f.perror(out.WriteString(`"`))
			}
			f.perror(out.WriteString(`>`))
			if page == 0 {
				f.perror(out.WriteString(`...`))
			} else {
				f.perror(out.WriteString(strconv.FormatUint(uint64(page), 10)))
			}
			f.perror(out.WriteString(`</p>`))
		}

		f.perror(out.WriteString(`</div>`))
	}

	f.perror(out.WriteString(`</div>`))

	return out.String()
}
