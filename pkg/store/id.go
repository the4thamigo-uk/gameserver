package store

// ID indicates the identity of the object in the store, comprising a string ID and version number
type ID struct {
	ID      string `json:"id,omitempty"`
	Version int    `json:"version,omitempty"`
}

// NewID creates a new ID instance
func NewID(id string, ver int) ID {
	return ID{
		ID:      id,
		Version: ver,
	}
}

// WithVersion returns a new ID with the specified version number
func (id ID) WithVersion(ver int) ID {
	return NewID(id.ID, ver)
}
