package pageactions

import (
	"strings"
)

// RenderOverviewActions renders the HTML output for the internal actions.
func (pageActions *PageActions) RenderOverviewActions() string {
	if !pageActions.IsFilled() {
		return ""
	}

	out := strings.Builder{}

	pageActions.perror(out.WriteString(`<div class="actions">`))

	for k := range pageActions.actions {
		pageActions.perror(out.WriteString(`<a href="`))
		pageActions.perror(out.WriteString(pageActions.ctx.AdminURL()))
		pageActions.perror(out.WriteString(`/`))
		pageActions.perror(out.WriteString(pageActions.actions[k].targetPath))
		pageActions.perror(out.WriteString(`"`))

		if pageActions.actions[k].listAction || pageActions.actions[k].class != "" {
			pageActions.perror(out.WriteString(` class="`))
			if pageActions.actions[k].listAction {
				pageActions.perror(out.WriteString(`listAction`))
			}
			if pageActions.actions[k].class != "" {
				pageActions.perror(out.WriteString(` `))
				pageActions.perror(out.WriteString(pageActions.actions[k].class))
			}
			pageActions.perror(out.WriteString(`"`))
		}

		pageActions.perror(out.WriteString(`>`))
		pageActions.perror(out.WriteString(pageActions.actions[k].name))
		pageActions.perror(out.WriteString(`</a>`))
	}

	pageActions.perror(out.WriteString(`</div>`))

	return out.String()
}
