package repositories

type Repositories struct{}

// New returns a new instance of Repositories.
func New() (*Repositories, error) {
	r := &Repositories{}
	return r, nil
}
