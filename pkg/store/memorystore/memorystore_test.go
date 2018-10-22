package memorystore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/the4thamigo-uk/gameserver/pkg/store"
	"testing"
)

type data struct {
	X int
}

func newData() interface{} { return &data{} }

func TestMemoryStore_SaveNew(t *testing.T) {
	id := store.NewID("123", 0)
	s := New()
	obj := data{X: 1}
	id1, err := s.Save(id, obj)
	require.Nil(t, err)
	assert.Equal(t, id.WithVersion(1), id1)
}

func TestMemoryStore_SaveNewVersion(t *testing.T) {
	id := store.NewID("123", 1)
	s := New()
	obj := data{X: 1}
	id1, err := s.Save(id, obj)
	require.Nil(t, err)
	assert.Equal(t, id.WithVersion(2), id1)
}

func TestMemoryStore_ConsecutiveSave(t *testing.T) {
	id := store.NewID("123", 0)
	s := New()
	obj := data{X: 1}
	id1, err := s.Save(id, obj)
	require.Nil(t, err)

	id2, err := s.Save(id1, obj)
	require.Nil(t, err)
	assert.Equal(t, id.WithVersion(2), id2)
}

func TestMemoryStore_SaveLatest(t *testing.T) {
	id := store.NewID("123", 0)
	s := New()
	obj := data{X: 1}
	_, err := s.Save(id, obj)
	require.Nil(t, err)

	id2, err := s.Save(id, obj)
	require.Nil(t, err)
	assert.Equal(t, id.WithVersion(2), id2)
}

func TestMemoryStore_SaveWrongVersionFails(t *testing.T) {
	id := store.NewID("123", 1)
	s := New()
	obj := data{X: 1}
	_, err := s.Save(id, obj)
	require.Nil(t, err)

	id1, err := s.Save(id.WithVersion(1), obj)
	assert.NotNil(t, err.(store.ErrWrongVersion))
	assert.Equal(t, id.WithVersion(1), id1)

	id2, err := s.Save(id.WithVersion(3), obj)
	assert.NotNil(t, err.(store.ErrWrongVersion))
	assert.Equal(t, id.WithVersion(3), id2)
}

func TestMemoryStore_SaveBadDataFails(t *testing.T) {
	id := store.NewID("123", 0)
	s := New()
	obj := make(chan int)
	_, err := s.Save(id, obj)
	assert.NotNil(t, err.(store.ErrData))
}

func TestMemoryStore_LoadLatest(t *testing.T) {
	id := store.NewID("123", 0)
	s := New()
	obj1 := data{X: 1}
	id1, err := s.Save(id, obj1)
	require.Nil(t, err)

	var obj2 data
	id2, err := s.Load(id1, &obj2)
	require.Nil(t, err)
	assert.Equal(t, obj1, obj2)
	assert.Equal(t, id2, id1)
}

func TestMemoryStore_LoadVersion(t *testing.T) {
	id := store.NewID("123", 0)
	s := New()
	obj1 := data{X: 1}
	id1, err := s.Save(id, obj1)
	require.Nil(t, err)

	var obj2 data
	id2, err := s.Load(id1, &obj2)
	require.Nil(t, err)
	assert.Equal(t, obj1, obj2)
	assert.Equal(t, id2, id1)
}

func TestMemoryStore_LoadWrongVersionFails(t *testing.T) {
	id := store.NewID("123", 0)
	s := New()
	obj1 := data{X: 1}
	_, err := s.Save(id, obj1)
	require.Nil(t, err)

	var obj2 data
	id1, err := s.Load(id.WithVersion(2), &obj2)
	assert.NotNil(t, err.(store.ErrWrongVersion))
	assert.Equal(t, id.WithVersion(2), id1)
}

func TestMemoryStore_LoadBadDataFails(t *testing.T) {
	// TODO : this test is bound to implementation
	id := store.NewID("123", 0)
	s := New()
	s.m[id.ID] = &storeItem{data: []byte("garbage")}

	var obj data
	_, err := s.Load(id, obj)
	assert.NotNil(t, err.(store.ErrData))
}

func TestMemoryStore_LoadAll(t *testing.T) {
	s := New()
	obj1 := data{X: 1}
	id1 := store.NewID("item1", 1)
	_, err := s.Save(id1, obj1)
	require.Nil(t, err)
	obj2 := data{X: 2}
	id2 := store.NewID("item21", 2)
	_, err = s.Save(id2, obj2)
	require.Nil(t, err)

	objs, err := s.LoadAll(newData)
	require.Nil(t, err)
	assert.Len(t, objs, 2)
	assert.Equal(t, obj1, *objs[id1.WithVersion(2)].(*data))
	assert.Equal(t, obj2, *objs[id2.WithVersion(3)].(*data))
}
