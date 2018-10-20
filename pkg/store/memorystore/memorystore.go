package memorystore

import (
	"github.com/the4thamigo_uk/gameserver/pkg/store"
	"sync"
)

// MemoryStore very simple storage of objects in memory.
// Assumes Last-write-wins. Safe to call from multiple goroutines.
type MemoryStore struct {
	mtx sync.Mutex
	m   map[string]*memoryStoreItem
}

type memoryStoreItem struct {
	ver  int
	data string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		m: map[string]*memoryStoreItem{},
	}
}

func (s *MemoryStore) Save(id store.ID, obj interface{}) (store.ID, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	item, ok := s.m[id.ID]
	if ok {
		if id.Version > 0 && item.ver != id.Version {
			return id, errWrongVersion(id.ID, id.Version, item.ver)
		}
		id.Version = item.ver
	}

	data, err := toJSONBase64(obj)
	if err != nil {
		return id, errData(id.ID, err)
	}

	id.Version += 1
	s.m[id.ID] = &memoryStoreItem{
		ver:  id.Version,
		data: data,
	}
	return id, nil
}

func (s *MemoryStore) Load(id store.ID, obj interface{}) (store.ID, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	item, ok := s.m[id.ID]
	if !ok {
		return id, errNotFound(id.ID)
	}
	if id.Version > 0 && item.ver != id.Version {
		return id, errWrongVersion(id.ID, id.Version, item.ver)
	}

	err := fromJSONBase64(item.data, obj)
	if err != nil {
		return id, errData(id.ID, err)
	}
	id.Version = item.ver
	return id, err
}

func (s *MemoryStore) LoadAll(newData func() interface{}) (map[string]interface{}, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	objs := map[string]interface{}{}
	for id, item := range s.m {
		obj := newData()
		err := fromJSONBase64(item.data, obj)
		if err != nil {
			return nil, err
		}
		objs[id] = obj
	}

	return objs, nil
}
