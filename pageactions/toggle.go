package pageactions

import (
	"strings"
)

// AddToggle adds a default Create PageAction if the UserRight requirements are met.
func (pageActions *PageActions) AddToggle() {
	pageActions.AddToggleWithFieldAndPath("", "")
}

// AddToggleWithField adds a Toggle PageAction with a custom field if the
// UserRight requirements are met.
func (pageActions *PageActions) AddToggleWithField(field string) {
	pageActions.AddToggleWithFieldAndPath(field, "")
}

// AddToggleWithPath adds a Toggle PageAction with a custom path if the
// UserRight requirements are met.
func (pageActions *PageActions) AddToggleWithPath(path string) {
	pageActions.AddToggleWithFieldAndPath("", path)
}

// AddToggleWithFieldAndPath adds a Toggle PageAction with a custom field and path if the
// UserRight requirements are met.
func (pageActions *PageActions) AddToggleWithFieldAndPath(field string, path string) {
	if !pageActions.ctx.HasUserRight("Update" + pageActions.subject) {
		return
	}
	if field == "" {
		field = "Active"
	}
	if path == "" {
		path = pageActions.subject + "/Toggle" + field
	}
	pageActions.actions = append(pageActions.actions, &pageAction{
		name:       pageActions.ctx.Translate("toggle" + strings.Title(field) + "Selection"),
		targetPath: path,
		listAction: true,
		class:      "toggleActive",
	})
}
