package themes

var _ Viewable = &View{}

// Viewable represents an object that provides visual rendering processing.
type Viewable interface {
	ID() int
	Code() string
	Render() []byte
}

// View manages visual processing.
type View struct {
	id   int
	code string
}

// ID returns the unique View identifier.
func (v *View) ID() int {
	return v.id
}

// Code returns the unique View code.
func (v *View) Code() string {
	return v.code
}

// Render returns the View's final output.
func (v *View) Render() []byte {
	return nil
}

// NewView returns a new instance of View.
func NewView() *View {
	return &View{}
}
