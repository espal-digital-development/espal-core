package storage

// Storage engine definition.
type Storage interface {
	Exists(key string) bool
	Iterate(iterator func(key string, value []byte, err error) (keepCycling bool))
	Set(key string, value []byte) error
	Get(key string) ([]byte, bool, error)
	Delete(key string) error
}
