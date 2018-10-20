package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary_TwoWords(t *testing.T) {
	d := NewDictionary([]string{"apple", "banana"})
	s := d.GetAt(-3)
	assert.Equal(t, "banana", s)
	s = d.GetAt(-2)
	assert.Equal(t, "apple", s)
	s = d.GetAt(-1)
	assert.Equal(t, "banana", s)
	s = d.GetAt(0)
	assert.Equal(t, "apple", s)
	s = d.GetAt(1)
	assert.Equal(t, "banana", s)
	s = d.GetAt(2)
	assert.Equal(t, "apple", s)
	s = d.GetAt(3)
	assert.Equal(t, "banana", s)
}

func TestDictionary_ThreeWords(t *testing.T) {
	d := NewDictionary([]string{"apple", "banana", "carrot"})
	s := d.GetAt(-3)
	assert.Equal(t, "apple", s)
	s = d.GetAt(-2)
	assert.Equal(t, "banana", s)
	s = d.GetAt(-1)
	assert.Equal(t, "carrot", s)
	s = d.GetAt(0)
	assert.Equal(t, "apple", s)
	s = d.GetAt(1)
	assert.Equal(t, "banana", s)
	s = d.GetAt(2)
	assert.Equal(t, "carrot", s)
	s = d.GetAt(3)
	assert.Equal(t, "apple", s)
}
