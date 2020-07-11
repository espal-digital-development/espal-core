package storage

// Storage engine definition.
type Storage interface {
	Exists(key string) bool
	Get(key string) ([]byte, bool, error)
	Iterate(iterator func(key string, value []byte, err error) (keepCycling bool))
}

// Modifyable adds modification capabilities to the default Storage interface.
type Modifyable interface {
	Storage
	Set(key string, value []byte) error
	Delete(key string) error
	LoadAllFromPath(subjectPath string) error
}
