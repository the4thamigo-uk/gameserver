package store

// ErrWrongVersion indicates that a storage operation is invalid as the ID specifies a different version
type ErrWrongVersion error

// ErrNotFound indicates that the object could not be found in the store
type ErrNotFound error

// ErrData indicates that there was some error during serialization/deserialization of the object
type ErrData error
