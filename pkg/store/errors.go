package store

import (
	"github.com/pkg/errors"
)

// ErrWrongVersion indicates that a storage operation is invalid as the ID specifies a different version
type ErrWrongVersion struct {
	error
}

// ErrNotFound indicates that the object could not be found in the store
type ErrNotFound struct {
	error
}

// ErrData indicates that there was some error during serialization/deserialization of the object
type ErrData struct {
	error
}

// NewErrWrongVersion constructs a new ErrWrongVersion
func NewErrWrongVersion(id string, expVer int, actVer int) error {
	return ErrWrongVersion{
		errors.Errorf("Error saving object '%v'. Expected version %v, actual version %v", id, expVer, actVer)}
}

// NewErrNotFound constructs a new ErrNotFound
func NewErrNotFound(id string) error {
	return ErrNotFound{
		errors.Errorf(`Failed to find data for key '%v'`, id)}
}

// NewErrData constructs a new ErrData
func NewErrData(id string, err error) error {
	return ErrData{
		errors.Wrapf(err, `Failed to extract data for key '%v'`, id)}
}
