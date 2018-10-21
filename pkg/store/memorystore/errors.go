package memorystore

import (
	"github.com/pkg/errors"
	"github.com/the4thamigo-uk/gameserver/pkg/store"
)

func errWrongVersion(id string, expVer int, actVer int) error {
	return store.ErrWrongVersion(
		errors.Errorf("Error saving object '%v'. Expected version %v, actual version %v", id, expVer, actVer))
}

func errNotFound(id string) error {
	return store.ErrNotFound(
		errors.Errorf(`Failed to find data for key '%v'`, id))
}

func errData(id string, err error) error {
	return store.ErrData(
		errors.Wrapf(err, `Failed to extract data for key '%v'`, id))
}
