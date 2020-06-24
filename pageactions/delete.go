package pageactions

// AddDelete adds a default Create PageAction if the UserRight requirements are met.
func (pageActions *PageActions) AddDelete() {
	pageActions.AddDeleteWithPath("")
}

// AddDeleteWithPath adds a Delete PageAction with a custom path if the
// UserRight requirements are met.
func (pageActions *PageActions) AddDeleteWithPath(path string) {
	if !pageActions.ctx.HasUserRight("Delete" + pageActions.subject) {
		return
	}
	if path == "" {
		path = pageActions.subject + "/Delete"
	}
	pageActions.actions = append(pageActions.actions, &pageAction{
		name:       pageActions.ctx.Translate("deleteSelected"),
		targetPath: path,
		listAction: true,
		class:      "delete",
	})
}
