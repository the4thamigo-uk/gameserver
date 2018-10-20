package domain

type dictionary struct {
	Words []string
}

// Dictionary is an unsorted immutable ring of words
type Dictionary struct {
	id         string
	dictionary dictionary
}

func NewDictionary(words []string) *Dictionary {
	return &Dictionary{
		dictionary: dictionary{
			Words: words,
		},
	}
}

func (d *Dictionary) ID() string {
	return d.id
}

// Get returns a word form the Dictionary based on the value of a given position in the ring
func (d *Dictionary) GetAt(pos int) string {
	i := toIndex(pos, len(d.dictionary.Words))
	return d.dictionary.Words[i]
}

func toIndex(pos int, len int) int {
	return (pos%len + len) % len
}
