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
func (v *FormView) Errors() string {
	return v.validator.RenderErrors()
}

// Open will render the form open tag.
func (v *FormView) Open() string {
	return v.validator.RenderOpen()
}

// Field will render the form field and resolve it's rules and presets.
func (v *FormView) Field(name string) string {
	return v.validator.RenderField(name)
}

// CreateUpdateActions will render all admin create/update
// actions of an admin module overview page.
func (v *FormView) CreateUpdateActions(fieldName string, url string) string {
	return v.validator.RenderCreateUpdateActions(fieldName, url)
}

// ContainsSelectSearch determines if at least one ChoiceType field is
// present with Searchable active.
func (v *FormView) ContainsSelectSearch() bool {
	return v.validator.ContainsSelectSearch()
}

// New returns a new intance of FormView.
func New(validator validator) *FormView {
	return &FormView{
		validator: validator,
	}
}
