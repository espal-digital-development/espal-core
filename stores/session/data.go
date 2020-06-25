package session

// @synthesize-ignore

// DataEntries represents the data dictionary.
type DataEntries interface {
	Get(key uint8) (bool, string)
	Set(key uint8, value string)
	All() map[uint8]string
}

type dataEntries struct {
	entries map[uint8]string
}

// Get returns the data entry for the given key.
func (d *dataEntries) Get(key uint8) (bool, string) {
	if value, ok := d.entries[key]; ok {
		return true, value
	}
	return false, ""
}

// Set sets the data entry for the given key.
func (d *dataEntries) Set(key uint8, value string) {
	d.entries[key] = value
}

// All returns all known entries.
func (d *dataEntries) All() map[uint8]string {
	return d.entries
}

func newDataEntries() *dataEntries {
	return &dataEntries{
		entries: map[uint8]string{},
	}
}

// NewDataEntries returns a new instance of DataEntries.
func NewDataEntries() DataEntries {
	return newDataEntries()
}
