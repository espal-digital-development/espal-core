package pageactions

import (
	"strings"
)

// RenderOverviewActions renders the HTML output for the internal actions.
func (a *PageActions) RenderOverviewActions() string {
	if !a.IsFilled() {
		return ""
	}

	out := strings.Builder{}

	out.WriteString(`<div class="actions">`)

	for k := range a.actions {
		out.WriteString(`<a href="`)
		out.WriteString(a.ctx.AdminURL())
		out.WriteString(`/`)
		out.WriteString(a.actions[k].targetPath)
		out.WriteString(`"`)

		if a.actions[k].listAction || a.actions[k].class != "" {
			out.WriteString(` class="`)
			if a.actions[k].listAction {
				out.WriteString(`listAction`)
			}
			if a.actions[k].class != "" {
				out.WriteString(` `)
				out.WriteString(a.actions[k].class)
			}
			out.WriteString(`"`)
		}

		out.WriteString(`>`)
		out.WriteString(a.actions[k].name)
		out.WriteString(`</a>`)
	}

	out.WriteString(`</div>`)

	return out.String()
}
