package pages

type Pages struct{}

// New returns a new instance of Pages.
func New() (*Pages, error) {
	p := &Pages{}
	return p, nil
}
