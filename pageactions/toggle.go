package pageactions

import (
	"strings"
)

// AddToggle adds a default Create PageAction if the UserRight requirements are met.
func (a *PageActions) AddToggle() {
	a.AddToggleWithFieldAndPath("", "")
}

// AddToggleWithField adds a Toggle PageAction with a custom field if the UserRight requirements are met.
func (a *PageActions) AddToggleWithField(field string) {
	a.AddToggleWithFieldAndPath(field, "")
}

// AddToggleWithPath adds a Toggle PageAction with a custom path if the UserRight requirements are met.
func (a *PageActions) AddToggleWithPath(path string) {
	a.AddToggleWithFieldAndPath("", path)
}

// AddToggleWithFieldAndPath adds a Toggle PageAction with a custom field and path if the UserRight requirements are
// met.
func (a *PageActions) AddToggleWithFieldAndPath(field string, path string) {
	if !a.ctx.HasUserRight("Update" + a.subject) {
		return
	}
	if field == "" {
		field = "Active"
	}
	if path == "" {
		path = a.subject + "/Toggle" + field
	}
	a.actions = append(a.actions, &pageAction{
		name:       a.ctx.Translate("toggle" + strings.Title(field) + "Selection"),
		targetPath: path,
		listAction: true,
		class:      "toggleActive",
	})
}
