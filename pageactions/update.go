package pageactions

import (
	"strings"
)

// AddUpdate adds a default Create PageAction if the UserRight requirements are met.
func (a *PageActions) AddUpdate() {
	a.AddUpdateWithFieldAndPath("", "")
}

// AddUpdateWithField adds a Update PageAction with a custom field if the
// UserRight requirements are met.
func (a *PageActions) AddUpdateWithField(field string) {
	a.AddUpdateWithFieldAndPath(field, "")
}

// AddUpdateWithPath adds a Update PageAction with a custom path if the
// UserRight requirements are met.
func (a *PageActions) AddUpdateWithPath(path string) {
	a.AddUpdateWithFieldAndPath("", path)
}

// AddUpdateWithFieldAndPath adds a Update PageAction with a custom field and path if the
// UserRight requirements are met.
func (a *PageActions) AddUpdateWithFieldAndPath(field string, path string) {
	if !a.ctx.HasUserRight("Update" + a.subject) {
		return
	}
	if path == "" {
		path = a.subject + "/Update" + field
	}
	a.actions = append(a.actions, &pageAction{
		name:       a.ctx.Translate("update" + strings.Title(field) + "Selection"),
		targetPath: path,
		listAction: true,
		class:      "toggleActive",
	})
}
