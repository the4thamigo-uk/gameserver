package domain

import (
	"github.com/the4thamigo-uk/gameserver/pkg/store"
)

// HangmanResult is the result of the last action in a game of hangman.
type HangmanResult struct {
	ID      store.ID `json:"id,omitempty"`
	Current string   `json:"current,omitempty"`
	Word    *string  `json:"word,omitempty"`
	Turns   int      `json:"turns,omitempty"`
	State   State    `json:"state,omitempty"`
	Success *bool    `json:"success,omitempty"`
}

// HangmanResults is a slice of HangmanResult
type HangmanResults []*HangmanResult

func newHangmanResult(g *Hangman, id store.ID, success *bool) *HangmanResult {
	res := &HangmanResult{
		ID:      id,
		Current: g.Current(),
		Turns:   g.Turns(),
		State:   g.State(),
		Success: success,
	}
	if g.State() != Play {
		word := g.Word()
		res.Word = &word
	}
	return res
}

// LoadHangman loads a specific game from the store
func LoadHangman(s store.Store, id store.ID) (*Hangman, store.ID, error) {
	var g Hangman
	id, err := s.Load(id, &g.hangman)
	if err != nil {
		return nil, id, err
	}
	return &g, id, err
}
func newHangman() interface{} {
	return &hangman{}
}

// LoadAllHangman loads all hangman games from the store
func LoadAllHangman(s store.Store) (map[store.ID]*Hangman, error) {
	objs, err := s.LoadAll(newHangman)
	if err != nil {
		return nil, err
	}
	gs := map[store.ID]*Hangman{}
	for id, obj := range objs {
		gs[id] = &Hangman{
			hangman: *obj.(*hangman),
		}
	}
	return gs, nil
}

// SaveHangman loads a specific game from the store
func SaveHangman(s store.Store, id store.ID, g *Hangman) (store.ID, error) {
	id, err := s.Save(id, &g.hangman)
	if err != nil {
		return id, err
	}
	return id, nil

}

// CreateHangman creates a new game instance and stores it.
func CreateHangman(stg store.Store, word string, turns int) (*HangmanResult, error) {
	g, err := NewHangman(word, turns)
	if err != nil {
		return nil, err
	}
	id := store.NewID(NewID(), 0)
	id, err = SaveHangman(stg, id, g)
	if err != nil {
		return nil, err
	}
	return newHangmanResult(g, id, nil), nil
}

// ListHangman returns details of all current games of hangman in the store
func ListHangman(stg store.Store) (HangmanResults, error) {
	gs, err := LoadAllHangman(stg)
	if err != nil {
		return nil, err
	}
	rs := HangmanResults{}
	for id, g := range gs {
		r := newHangmanResult(g, id, nil)
		rs = append(rs, r)
	}
	return rs, nil
}

// JoinHangman loads an existing game instance.
func JoinHangman(stg store.Store, id store.ID) (*HangmanResult, error) {
	g, id, err := LoadHangman(stg, id)
	if err != nil {
		return nil, err
	}
	return newHangmanResult(g, id, nil), nil
}

// PlayLetter submits a letter guess to a specified game
func PlayLetter(stg store.Store, id store.ID, letter rune) (*HangmanResult, error) {
	g, id, err := LoadHangman(stg, id)
	if err != nil {
		return nil, err
	}
	ok, err := g.PlayLetter(letter)
	if err != nil {
		return nil, err
	}
	id, err = SaveHangman(stg, id, g)
	if err != nil {
		return nil, err
	}
	return newHangmanResult(g, id, ptrBool(ok)), nil
}

// PlayWord submits a word guess to a specified game.
func PlayWord(stg store.Store, id store.ID, word string) (*HangmanResult, error) {
	g, ver, err := LoadHangman(stg, id)
	if err != nil {
		return nil, err
	}
	ok, err := g.PlayWord(word)
	if err != nil {
		return nil, err
	}
	ver, err = SaveHangman(stg, ver, g)
	if err != nil {
		return nil, err
	}
	return newHangmanResult(g, ver, ptrBool(ok)), nil
}

func ptrBool(b bool) *bool {
	return &b
}
