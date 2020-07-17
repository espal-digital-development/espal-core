package pageactions

import (
	"github.com/espal-digital-development/espal-core/text"
)

// AddCreate adds a default Create PageAction if the UserRight requirements are met.
func (a *PageActions) AddCreate() {
	a.AddCreateWithFieldAndPath("", "")
}

// AddCreateWithPath adds a Create PageAction with a custom path if the UserRight requirements are met.
func (a *PageActions) AddCreateWithPath(path string) {
	a.AddCreateWithFieldAndPath("", path)
}

// AddCreateWithFieldAndPath adds a Create PageAction with a custom field and path if the UserRight requirements are
// met.
func (a *PageActions) AddCreateWithFieldAndPath(field string, path string) {
	if !a.ctx.HasUserRight("Create" + a.subject) {
		return
	}
	if field == "" {
		field = a.subject
	}
	if path == "" {
		path = a.subject + "/Create"
	}
	a.actions = append(a.actions, &pageAction{
		name:       a.ctx.Translate("new") + " " + a.ctx.Translate(text.LowerFirst(field)),
		targetPath: path,
		class:      "create",
	})
}
