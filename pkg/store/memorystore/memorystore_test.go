package memorystore

import (
	"github.com/stretchr/testify/assert"
	"github.com/the4thamigo_uk/gameserver/pkg/store"
	"testing"
)

type data struct {
	X int
}

func TestMemoryStore_SaveNew(t *testing.T) {
	id := store.NewID("123", 0)
	s := NewMemoryStore()
	obj := data{X: 1}
	id1, err := s.Save(id, obj)
	assert.Nil(t, err)
	assert.Equal(t, id.WithVersion(1), id1)
}

func TestMemoryStore_SaveNewVersion(t *testing.T) {
	id := store.NewID("123", 1)
	s := NewMemoryStore()
	obj := data{X: 1}
	id1, err := s.Save(id, obj)
	assert.Nil(t, err)
	assert.Equal(t, id.WithVersion(2), id1)
}

func TestMemoryStore_ConsecutiveSave(t *testing.T) {
	id := store.NewID("123", 0)
	s := NewMemoryStore()
	obj := data{X: 1}
	id1, err := s.Save(id, obj)
	assert.Nil(t, err)

	id2, err := s.Save(id1, obj)
	assert.Nil(t, err)
	assert.Equal(t, id.WithVersion(2), id2)
}

func TestMemoryStore_SaveLatest(t *testing.T) {
	id := store.NewID("123", 0)
	s := NewMemoryStore()
	obj := data{X: 1}
	_, err := s.Save(id, obj)
	assert.Nil(t, err)

	id2, err := s.Save(id, obj)
	assert.Nil(t, err)
	assert.Equal(t, id.WithVersion(2), id2)
}

func TestMemoryStore_SaveWrongVersionFails(t *testing.T) {
	id := store.NewID("123", 1)
	s := NewMemoryStore()
	obj := data{X: 1}
	_, err := s.Save(id, obj)
	assert.Nil(t, err)

	id1, err := s.Save(id.WithVersion(1), obj)
	assert.NotNil(t, err.(store.ErrWrongVersion))
	assert.Equal(t, id.WithVersion(1), id1)

	id2, err := s.Save(id.WithVersion(3), obj)
	assert.NotNil(t, err.(store.ErrWrongVersion))
	assert.Equal(t, id.WithVersion(3), id2)
}

func TestMemoryStore_SaveBadDataFails(t *testing.T) {
	id := store.NewID("123", 0)
	s := NewMemoryStore()
	obj := make(chan int, 0)
	_, err := s.Save(id, obj)
	assert.NotNil(t, err.(store.ErrData))
}

func TestMemoryStore_LoadLatest(t *testing.T) {
	id := store.NewID("123", 0)
	s := NewMemoryStore()
	obj1 := data{X: 1}
	id1, err := s.Save(id, obj1)
	assert.Nil(t, err)

	var obj2 data
	id2, err := s.Load(id1, &obj2)
	assert.Nil(t, err)
	assert.Equal(t, obj1, obj2)
	assert.Equal(t, id2, id1)
}

func TestMemoryStore_LoadVersion(t *testing.T) {
	id := store.NewID("123", 0)
	s := NewMemoryStore()
	obj1 := data{X: 1}
	id1, err := s.Save(id, obj1)
	assert.Nil(t, err)

	var obj2 data
	id2, err := s.Load(id1, &obj2)
	assert.Nil(t, err)
	assert.Equal(t, obj1, obj2)
	assert.Equal(t, id2, id1)
}

func TestMemoryStore_LoadWrongVersionFails(t *testing.T) {
	id := store.NewID("123", 0)
	s := NewMemoryStore()
	obj1 := data{X: 1}
	_, err := s.Save(id, obj1)
	assert.Nil(t, err)

	var obj2 data
	id1, err := s.Load(id.WithVersion(2), &obj2)
	assert.NotNil(t, err.(store.ErrWrongVersion))
	assert.Equal(t, id.WithVersion(2), id1)
}

func TestMemoryStore_LoadBadDataFails(t *testing.T) {
	id := store.NewID("123", 0)
	s := NewMemoryStore()
	s.m[id.ID] = &memoryStoreItem{data: "garbage"}

	var obj data
	_, err := s.Load(id, obj)
	assert.NotNil(t, err.(store.ErrData))
}
