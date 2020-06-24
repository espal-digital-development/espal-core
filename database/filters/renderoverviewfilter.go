package filters

import (
	"strconv"
	"strings"
)

// RenderOverviewFilter will render the filter actions of an
// admin module overview page.
func (filter *filter) RenderOverviewFilter(ctx Context) string {
	out := strings.Builder{}

	filter.perror(out.WriteString(`<div class="filter">`))

	filter.perror(out.WriteString(`<div class="filterInfo">`))
	filter.perror(out.WriteString(strconv.FormatUint(uint64(filter.totalResults), 10)))
	filter.perror(out.WriteString(` `))
	if filter.totalResults == 1 {
		filter.perror(out.WriteString(ctx.Translate("result")))
	} else {
		filter.perror(out.WriteString(ctx.TranslatePlural("result")))
	}
	filter.perror(out.WriteString(`</div>`))

	filter.perror(out.WriteString(`<div>`))
	if filter.ShouldShowSearch() {
		filter.perror(out.WriteString(`<input id="filterSearch" placeholder="`))
		filter.perror(out.WriteString(ctx.Translate("search")))
		filter.perror(out.WriteString(`"`))
		if filter.search != "" {
			filter.perror(out.WriteString(`value="`))
			filter.perror(out.WriteString(filter.search))
			filter.perror(out.WriteString(`"`))
		}
		filter.perror(out.WriteString(`>`))
	}
	filter.perror(out.WriteString(`<div class="filters"></div>`))
	filter.perror(out.WriteString(`</div>`))

	if filter.totalPages > 1 {
		filter.perror(out.WriteString(`<div class="pagination disable-select">`))

		filter.perror(out.WriteString(`<select>`))
		for _, r := range []uint{5, 10, 25, 50, 100} {
			filter.perror(out.WriteString(`<option value="`))
			filter.perror(out.WriteString(strconv.FormatUint(uint64(r), 10)))
			filter.perror(out.WriteString(`"`))
			if r == filter.limit {
				filter.perror(out.WriteString(` selected`))
			}
			filter.perror(out.WriteString(`>`))
			filter.perror(out.WriteString(strconv.FormatUint(uint64(r), 10)))
			filter.perror(out.WriteString(`</option>`))
		}
		filter.perror(out.WriteString(`</select>`))

		for _, page := range filter.PaginationBlocks() {
			filter.perror(out.WriteString(`<p `))
			if page == 0 || filter.CurrentPage() == page {
				filter.perror(out.WriteString(`class="`))
				if filter.CurrentPage() == page {
					filter.perror(out.WriteString(`current`))
				} else if page == 0 {
					filter.perror(out.WriteString(`separator`))
				}
				filter.perror(out.WriteString(`"`))
			}
			filter.perror(out.WriteString(`>`))
			if page == 0 {
				filter.perror(out.WriteString(`...`))
			} else {
				filter.perror(out.WriteString(strconv.FormatUint(uint64(page), 10)))
			}
			filter.perror(out.WriteString(`</p>`))
		}

		filter.perror(out.WriteString(`</div>`))
	}

	filter.perror(out.WriteString(`</div>`))

	return out.String()
}
