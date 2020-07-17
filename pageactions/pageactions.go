package pageactions

var _ Actions = &PageActions{}

type context interface {
	HasUserRight(string) bool
	Translate(string) string
	AdminURL() string
}

// Actions represents a list of multiple action references.
type Actions interface {
	AddCreate()
	AddCreateWithPath(path string)
	AddCreateWithFieldAndPath(field string, path string)
	AddToggle()
	AddToggleWithField(field string)
	AddToggleWithPath(path string)
	AddToggleWithFieldAndPath(field string, path string)
	AddUpdate()
	AddUpdateWithField(field string)
	AddUpdateWithPath(path string)
	AddUpdateWithFieldAndPath(field string, path string)
	AddDelete()
	AddDeleteWithPath(path string)
	IsFilled() bool
	RenderOverviewActions() string
}

// PageActions is a list of multiple PageAction references.
type PageActions struct {
	actions    []*pageAction
	ctx        context
	hasResults bool
	subject    string
}

type pageAction struct {
	name       string
	targetPath string
	// listAction indicates wether the action needs to trigger on the
	// selected items in the result-list or a standalone action.
	listAction bool
	class      string
}

// IsFilled returns if there are any actions present.
func (a *PageActions) IsFilled() bool {
	return len(a.actions) > 0
}

// New returns an instantiated PageActions.
func New(ctx context, subject string, hasResults bool) *PageActions {
	return &PageActions{
		actions:    make([]*pageAction, 0, 3),
		ctx:        ctx,
		hasResults: hasResults,
		subject:    subject,
	}
}
