package store

// Store provides a way to load/save arbitary object to a storage medium.
// The interface uses an versioning scheme which is used
// as a mechanism for implementing an optimistic offline lock.
// On error the functions return the id that was passed in
type Store interface {
	// Save updates the given version of the object, and returns the new version.
	// If the version does not exist in the store, the object should still be stored.
	// If version <= 0 then the latest version is always overwritten.
	Save(id ID, obj interface{}) (ID, error)
	// Load retrieves a given version of the object.
	// If version <= 0 then the latest version is loaded.
	Load(id ID, obj interface{}) (ID, error)
	// LoadAll retrieves the latest version of all the objects in the store.
	LoadAll(newObj func() interface{}) (map[ID]interface{}, error)
}
