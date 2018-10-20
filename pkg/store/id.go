package store

type ID struct {
	ID      string
	Version int
}

func NewID(id string, ver int) ID {
	return ID{
		ID:      id,
		Version: ver,
	}
}

func (id ID) WithVersion(ver int) ID {
	return NewID(id.ID, ver)
}
