package memorystore

import (
	"encoding/json"
	"github.com/the4thamigo-uk/gameserver/pkg/store"
	"sync"
)

// Store is an in-memory object store
type Store struct {
	mtx sync.RWMutex
	m   map[string]*storeItem
}

type storeItem struct {
	ver  int
	data []byte
}

// New creates a new instance of an memorystore.Store.
func New() *Store {
	return &Store{
		m: map[string]*storeItem{},
	}
}

// Save updates the given version of the object, and returns the new version.
func (s *Store) Save(id store.ID, obj interface{}) (store.ID, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	item, ok := s.m[id.ID]
	if ok {
		if id.Version > 0 && item.ver != id.Version {
			return id, store.NewErrWrongVersion(id.ID, id.Version, item.ver)
		}
		id.Version = item.ver
	}

	data, err := json.Marshal(obj)
	if err != nil {
		return id, store.NewErrData(id.ID, err)
	}

	id.Version++
	s.m[id.ID] = &storeItem{
		ver:  id.Version,
		data: data,
	}
	return id, nil
}

// Load retrieves a given version of the object.
func (s *Store) Load(id store.ID, obj interface{}) (store.ID, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	item, ok := s.m[id.ID]
	if !ok {
		return id, store.NewErrNotFound(id.ID)
	}
	if id.Version > 0 && item.ver != id.Version {
		return id, store.NewErrWrongVersion(id.ID, id.Version, item.ver)
	}

	err := json.Unmarshal(item.data, obj)
	if err != nil {
		return id, store.NewErrData(id.ID, err)
	}
	id.Version = item.ver
	return id, err
}

// LoadAll retrieves the latest version of all the objects in the store.
func (s *Store) LoadAll(newData func() interface{}) (map[store.ID]interface{}, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	objs := map[store.ID]interface{}{}
	for id, item := range s.m {
		obj := newData()
		err := json.Unmarshal(item.data, obj)
		if err != nil {
			return nil, err
		}
		objs[store.NewID(id, item.ver)] = obj
	}

	return objs, nil
}
