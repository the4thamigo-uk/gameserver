package domain

// Dictionary is an unsorted immutable ring of words
type Dictionary struct {
	words []string
}

// NewDictionary creates a new dictionary object from a slice of words
func NewDictionary(words []string) *Dictionary {
	return &Dictionary{
		words: words,
	}
}

// GetAt returns a word form the Dictionary based on the value of a given position in the ring
func (d *Dictionary) GetAt(pos int) string {
	i := toIndex(pos, len(d.words))
	return d.words[i]
}

func toIndex(pos int, len int) int {
	return (pos%len + len) % len
}
