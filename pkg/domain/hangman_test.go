package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHangman_CreateLatin(t *testing.T) {
	g, err := NewHangman("english", 6)
	assert.Nil(t, err)
	assert.NotNil(t, g)
	assert.Equal(t, "ENGLISH", g.Word())
	assert.Equal(t, "       ", g.Current())
}

func TestHangman_CreateNonLatin(t *testing.T) {
	g, err := NewHangman("русский", 6)
	assert.Nil(t, err)
	assert.NotNil(t, g)
	assert.Equal(t, "РУССКИЙ", g.Word())
}

func TestHangman_CreateNonLetterFails(t *testing.T) {
	g, err := NewHangman("a word", 6)
	assert.NotNil(t, err.(ErrInvalidWord))
	assert.Nil(t, g)
}

func TestHangman_CreateBlankWordFails(t *testing.T) {
	g, err := NewHangman("", 6)
	assert.NotNil(t, err.(ErrInvalidWord))
	assert.Nil(t, g)
}

func TestHangman_PlayWordCorrect(t *testing.T) {
	g, _ := NewHangman("english", 6)
	ok, err := g.PlayWord("english")
	assert.True(t, ok)
	assert.Equal(t, 6, g.Turns())
	assert.Nil(t, err)
	assert.Equal(t, Win, g.State())
}

func TestHangman_PlayWordIncorrect(t *testing.T) {
	g, _ := NewHangman("english", 6)
	ok, err := g.PlayWord("german")
	assert.False(t, ok)
	assert.Nil(t, err)
	assert.Equal(t, 5, g.Turns())
	assert.Equal(t, "       ", g.Current())
	assert.Equal(t, Play, g.State())
}

func TestHangman_PlayNonWordFails(t *testing.T) {
	g, _ := NewHangman("english", 6)
	ok, err := g.PlayWord("a space")
	assert.False(t, ok)
	assert.NotNil(t, err.(ErrInvalidWord))
	assert.Equal(t, 6, g.Turns())
	assert.Equal(t, "       ", g.Current())
	assert.Equal(t, Play, g.State())
}

func TestHangman_PlayNonLetterFails(t *testing.T) {
	g, _ := NewHangman("english", 6)
	ok, err := g.PlayLetter(' ')
	assert.False(t, ok)
	assert.NotNil(t, err.(ErrInvalidLetter))
	assert.Equal(t, 6, g.Turns())
	assert.Equal(t, "       ", g.Current())
	assert.Equal(t, Play, g.State())
}

func TestHangman_PlayLetterExists(t *testing.T) {
	g, _ := NewHangman("english", 6)
	ok, err := g.PlayLetter('E')
	assert.True(t, ok)
	assert.Nil(t, err)
	assert.Equal(t, 6, g.Turns())
	assert.Equal(t, "ENGLISH", g.Word())
	assert.Equal(t, "E      ", g.Current())

	ok, err = g.PlayLetter('S')
	assert.True(t, ok)
	assert.Nil(t, err)
	assert.Equal(t, 6, g.Turns())
	assert.Equal(t, "ENGLISH", g.Word())
	assert.Equal(t, "E    S ", g.Current())

	ok, err = g.PlayLetter('I')
	assert.True(t, ok)
	assert.Nil(t, err)
	assert.Equal(t, 6, g.Turns())
	assert.Equal(t, "ENGLISH", g.Word())
	assert.Equal(t, "E   IS ", g.Current())
}

func TestHangman_PlayLetterNotExists(t *testing.T) {
	g, _ := NewHangman("english", 6)
	ok, err := g.PlayLetter('X')
	assert.False(t, ok)
	assert.Nil(t, err)
	assert.Equal(t, 5, g.Turns())
	assert.Equal(t, "ENGLISH", g.Word())
	assert.Equal(t, "       ", g.Current())
}

func TestHangman_PlayLetterAlreadyUsed(t *testing.T) {
	g, _ := NewHangman("english", 6)
	g.PlayLetter('E')
	_, err := g.PlayLetter('E')
	assert.NotNil(t, err.(ErrAlreadyUsed))
	assert.Equal(t, 6, g.Turns())
	assert.Equal(t, "ENGLISH", g.Word())
	assert.Equal(t, "E      ", g.Current())
}

func TestHangman_PlayLetterAfterHangmanOver(t *testing.T) {
	g, _ := NewHangman("english", 1)
	g.PlayLetter('A')
	assert.Equal(t, 0, g.Turns())
	_, err := g.PlayLetter('B')
	assert.NotNil(t, err.(ErrGameOver))
}

func TestHangman_PlayLetterStateLose(t *testing.T) {
	g, _ := NewHangman("c", 2)
	assert.Equal(t, Play, g.State())
	g.PlayLetter('A')
	assert.Equal(t, Play, g.State())
	g.PlayLetter('B')
	assert.Equal(t, Lose, g.State())
	_, err := g.PlayLetter('X')
	assert.NotNil(t, err.(ErrGameOver))
}

func TestHangman_PlayLetterStateWin(t *testing.T) {
	g, _ := NewHangman("cba", 2)
	assert.Equal(t, Play, g.State())
	g.PlayLetter('A')
	assert.Equal(t, Play, g.State())
	g.PlayLetter('B')
	assert.Equal(t, Play, g.State())
	g.PlayLetter('C')
	assert.Equal(t, Win, g.State())
	_, err := g.PlayLetter('X')
	assert.NotNil(t, err.(ErrGameOver))
}
