package storesmeta

// StoreObject object.
type StoreObject struct {
	Name string
	Type string
}

// StoresMeta object.
type StoresMeta struct {
	objects []*StoreObject
}

// All returns all Store Objects meta data.
func (s *StoresMeta) All() []*StoreObject {
	return s.objects
}

// New returns a new instance of StoresMeta.
func New() (*StoresMeta, error) {
	return &StoresMeta{
		objects: []*StoreObject{},
	}, nil
}
