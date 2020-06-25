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

	a.perror(out.WriteString(`<div class="actions">`))

	for k := range a.actions {
		a.perror(out.WriteString(`<a href="`))
		a.perror(out.WriteString(a.ctx.AdminURL()))
		a.perror(out.WriteString(`/`))
		a.perror(out.WriteString(a.actions[k].targetPath))
		a.perror(out.WriteString(`"`))

		if a.actions[k].listAction || a.actions[k].class != "" {
			a.perror(out.WriteString(` class="`))
			if a.actions[k].listAction {
				a.perror(out.WriteString(`listAction`))
			}
			if a.actions[k].class != "" {
				a.perror(out.WriteString(` `))
				a.perror(out.WriteString(a.actions[k].class))
			}
			a.perror(out.WriteString(`"`))
		}

		a.perror(out.WriteString(`>`))
		a.perror(out.WriteString(a.actions[k].name))
		a.perror(out.WriteString(`</a>`))
	}

	a.perror(out.WriteString(`</div>`))

	return out.String()
}
