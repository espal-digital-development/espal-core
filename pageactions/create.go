package pageactions

import (
	"github.com/espal-digital-development/espal-core/text"
)

// AddCreate adds a default Create PageAction if the UserRight requirements are met.
func (pageActions *PageActions) AddCreate() {
	pageActions.AddCreateWithFieldAndPath("", "")
}

// AddCreateWithPath adds a Create PageAction with a custom path if the
// UserRight requirements are met.
func (pageActions *PageActions) AddCreateWithPath(path string) {
	pageActions.AddCreateWithFieldAndPath("", path)
}

// AddCreateWithFieldAndPath adds a Create PageAction with a custom field and path if the
// UserRight requirements are met.
func (pageActions *PageActions) AddCreateWithFieldAndPath(field string, path string) {
	if !pageActions.ctx.HasUserRight("Create" + pageActions.subject) {
		return
	}
	if field == "" {
		field = pageActions.subject
	}
	if path == "" {
		path = pageActions.subject + "/Create"
	}
	pageActions.actions = append(pageActions.actions, &pageAction{
		name:       pageActions.ctx.Translate("new") + " " + pageActions.ctx.Translate(text.LowerFirst(field)),
		targetPath: path,
		class:      "create",
	})
}
