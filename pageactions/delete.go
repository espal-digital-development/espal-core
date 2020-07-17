package pageactions

// AddDelete adds a default Create PageAction if the UserRight requirements are met.
func (a *PageActions) AddDelete() {
	a.AddDeleteWithPath("")
}

// AddDeleteWithPath adds a Delete PageAction with a custom path if the UserRight requirements are met.
func (a *PageActions) AddDeleteWithPath(path string) {
	if !a.ctx.HasUserRight("Delete" + a.subject) {
		return
	}
	if path == "" {
		path = a.subject + "/Delete"
	}
	a.actions = append(a.actions, &pageAction{
		name:       a.ctx.Translate("deleteSelected"),
		targetPath: path,
		listAction: true,
		class:      "delete",
	})
}
