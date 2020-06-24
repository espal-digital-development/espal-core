package formview

var _ View = &FormView{}

type validator interface {
	RenderErrors() string
	RenderOpen() string
	RenderField(string) string
	RenderCreateUpdateActions(string, string) string
	ContainsSelectSearch() bool
}

// View represents an object that controls form view rendering logic.
type View interface {
	Errors() string
	Open() string
	Field(name string) string
	CreateUpdateActions(fieldName string, url string) string
	ContainsSelectSearch() bool
}

// FormView controls the view data that gets send to the html page.
type FormView struct {
	validator validator
}

// Errors will render the list of form errors.
// Returns an empty string when no errors are present.
func (formView *FormView) Errors() string {
	return formView.validator.RenderErrors()
}

// Open will render the form open tag.
func (formView *FormView) Open() string {
	return formView.validator.RenderOpen()
}

// Field will render the form field and resolve it's rules and presets.
func (formView *FormView) Field(name string) string {
	return formView.validator.RenderField(name)
}

// CreateUpdateActions will render all admin create/update
// actions of an admin module overview page.
func (formView *FormView) CreateUpdateActions(fieldName string, url string) string {
	return formView.validator.RenderCreateUpdateActions(fieldName, url)
}

// ContainsSelectSearch determines if at least one ChoiceType field is
// present with Searchable active.
func (formView *FormView) ContainsSelectSearch() bool {
	return formView.validator.ContainsSelectSearch()
}

// New returns a new intance of FormView.
func New(validator validator) *FormView {
	return &FormView{
		validator: validator,
	}
}
