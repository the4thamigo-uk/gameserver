package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/the4thamigo-uk/gameserver/pkg/store"
	"github.com/the4thamigo-uk/gameserver/pkg/store/memorystore"
	"testing"
)

func TestService_SaveLoad(t *testing.T) {
	s := memorystore.New()
	g1, _ := NewHangman("word", 6)
	id1, err := SaveHangman(s, store.NewID("123", 0), g1)
	assert.Nil(t, err)
	g2, id2, err := LoadHangman(s, id1)
	assert.Nil(t, err)
	assert.Equal(t, g1, g2)
	assert.Equal(t, id1, id2)
}

func TestService_VersionIncrements(t *testing.T) {
	s := memorystore.New()
	r1, err := CreateHangman(s, "word", 5)
	assert.Nil(t, err)
	assert.Equal(t, 1, r1.ID.Version)
	r2, err := PlayLetter(s, r1.ID, 'X')
	assert.Nil(t, err)
	assert.Equal(t, 2, r2.ID.Version)
	r3, err := PlayWord(s, r2.ID, "WRONG")
	assert.Nil(t, err)
	assert.Equal(t, 3, r3.ID.Version)
}

func TestService_VersionConflict(t *testing.T) {
	s := memorystore.New()
	r1, err := CreateHangman(s, "word", 5) // user 1 creates game
	assert.Nil(t, err)
	r2, err := JoinHangman(s, r1.ID) // user 2 joins game
	assert.Nil(t, err)
	_, err = PlayLetter(s, r1.ID, 'X') // user 1 plays
	assert.Nil(t, err)
	_, err = PlayLetter(s, r2.ID, 'Y') // user 2 plays
	assert.NotNil(t, err.(store.ErrWrongVersion))
}

func TestService_CreateHangman(t *testing.T) {
	s := memorystore.New()
	r, err := CreateHangman(s, "word", 5)
	assert.Nil(t, err)
	assert.NotEmpty(t, r.ID)
	assert.Equal(t, "    ", r.Current)
	assert.Equal(t, 5, r.Turns)
	assert.Equal(t, Play, r.State)
}

func TestService_JoinHangman(t *testing.T) {
	s := memorystore.New()
	r1, err := CreateHangman(s, "word", 5)
	assert.Nil(t, err)
	r2, err := JoinHangman(s, r1.ID)
	assert.Equal(t, r1, r2)
}

func TestService_CreateUniqueIDs(t *testing.T) {
	s := memorystore.New()
	r1, err := CreateHangman(s, "word", 5)
	assert.Nil(t, err)
	r2, err := CreateHangman(s, "word", 5)
	assert.Nil(t, err)
	assert.NotEqual(t, r1.ID, r2.ID)
}

func TestService_PlayWordCorrect(t *testing.T) {
	s := memorystore.New()
	r, _ := CreateHangman(s, "word", 5)
	r, err := PlayWord(s, r.ID, "word")
	assert.Nil(t, err)
	assert.NotEmpty(t, r.ID)
	assert.Equal(t, "WORD", r.Current)
	assert.Equal(t, 5, r.Turns)
	assert.Equal(t, Win, r.State)
	assert.True(t, r.Success)
}

func TestService_PlayWordIncorrect(t *testing.T) {
	s := memorystore.New()
	r, _ := CreateHangman(s, "word", 5)
	r, err := PlayWord(s, r.ID, "notword")
	assert.Nil(t, err)
	assert.NotEmpty(t, r.ID)
	assert.Equal(t, "    ", r.Current)
	assert.Equal(t, 4, r.Turns)
	assert.Equal(t, Play, r.State)
	assert.False(t, r.Success)
}

func TestService_PlayLetterCorrect(t *testing.T) {
	s := memorystore.New()
	r, _ := CreateHangman(s, "word", 5)
	r, err := PlayLetter(s, r.ID, 'R')
	assert.Nil(t, err)
	assert.NotEmpty(t, r.ID)
	assert.Equal(t, "  R ", r.Current)
	assert.Equal(t, 5, r.Turns)
	assert.Equal(t, Play, r.State)
	assert.True(t, r.Success)
}

func TestService_PlayLetterIncorrect(t *testing.T) {
	s := memorystore.New()
	r, _ := CreateHangman(s, "word", 5)
	r, err := PlayLetter(s, r.ID, 'X')
	assert.Nil(t, err)
	assert.NotEmpty(t, r.ID)
	assert.Equal(t, "    ", r.Current)
	assert.Equal(t, 4, r.Turns)
	assert.Equal(t, Play, r.State)
	assert.False(t, r.Success)
}

func TestService_ListGames(t *testing.T) {
	s := memorystore.New()
	r1, err := CreateHangman(s, "apple", 5)
	assert.Nil(t, err)
	r2, err := CreateHangman(s, "banana", 6)
	assert.Nil(t, err)
	rs, err := ListHangman(s)
	assert.Nil(t, err)

	t.Log(rs)

	assert.Len(t, rs, 2)
	assert.Equal(t, r1, rs[r1.ID.ID])
	assert.Equal(t, r2, rs[r2.ID.ID])
}
