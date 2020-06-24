package pageactions

import (
	"strings"
)

// AddUpdate adds a default Create PageAction if the UserRight requirements are met.
func (pageActions *PageActions) AddUpdate() {
	pageActions.AddUpdateWithFieldAndPath("", "")
}

// AddUpdateWithField adds a Update PageAction with a custom field if the
// UserRight requirements are met.
func (pageActions *PageActions) AddUpdateWithField(field string) {
	pageActions.AddUpdateWithFieldAndPath(field, "")
}

// AddUpdateWithPath adds a Update PageAction with a custom path if the
// UserRight requirements are met.
func (pageActions *PageActions) AddUpdateWithPath(path string) {
	pageActions.AddUpdateWithFieldAndPath("", path)
}

// AddUpdateWithFieldAndPath adds a Update PageAction with a custom field and path if the
// UserRight requirements are met.
func (pageActions *PageActions) AddUpdateWithFieldAndPath(field string, path string) {
	if !pageActions.ctx.HasUserRight("Update" + pageActions.subject) {
		return
	}
	if path == "" {
		path = pageActions.subject + "/Update" + field
	}
	pageActions.actions = append(pageActions.actions, &pageAction{
		name:       pageActions.ctx.Translate("update" + strings.Title(field) + "Selection"),
		targetPath: path,
		listAction: true,
		class:      "toggleActive",
	})
}
