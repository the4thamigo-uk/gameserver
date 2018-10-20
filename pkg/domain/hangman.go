package domain

import (
	"strings"
	"unicode"
)

type hangman struct {
	Turns   int
	Word    string
	Current string
	Used    string
}

type Hangman struct {
	hangman hangman
}

func NewHangman(word string, turns int) (*Hangman, error) {
	if !isValidWord(word) {
		return nil, errInvalidWord(word)
	}
	if !isValidTurns(turns) {
		return nil, errInvalidTurns(turns)
	}
	word = strings.ToUpper(word)
	current := strings.Repeat(" ", len([]rune(word)))

	return &Hangman{
		hangman: hangman{
			Turns:   int(turns),
			Word:    word,
			Current: current,
		},
	}, nil
}

func (g *Hangman) Turns() int {
	if g.hangman.Turns > 0 {
		return g.hangman.Turns
	}
	return 0
}

func (g *Hangman) Word() string {
	return g.hangman.Word
}

func (g *Hangman) Current() string {
	return g.hangman.Current
}

func (g *Hangman) State() State {
	if g.hangman.Current == g.hangman.Word {
		return Win
	}
	if g.hangman.Turns <= 0 {
		return Lose
	}

	return Play
}

func (g *Hangman) PlayLetter(l rune) (bool, error) {
	l = unicode.ToUpper(l)
	if g.State() != Play {
		return false, errGameOver()
	}
	if g.isAlreadyUsed(l) {
		return false, errAlreadyUsed(l)
	}
	if !isValidLetter(l) {
		return false, errInvalidLetter(l)
	}

	s := string(l)
	g.hangman.Used += s

	if !strings.Contains(g.hangman.Word, s) {
		g.hangman.Turns -= 1
		return false, nil
	}

	g.hangman.Current = strings.Map(
		func(r rune) rune {
			if strings.Contains(g.hangman.Used, string(r)) {
				return r
			}
			return ' '
		},
		g.hangman.Word)

	return true, nil
}

func (g *Hangman) PlayWord(word string) (bool, error) {
	if g.State() != Play {
		return false, errGameOver()
	}
	if !isValidWord(word) {
		return false, errInvalidWord(word)
	}
	word = strings.ToUpper(word)
	if word == g.hangman.Word {
		g.hangman.Current = word
		return true, nil
	}
	g.hangman.Turns -= 1
	return false, nil
}

func (g *Hangman) isAlreadyUsed(letter rune) bool {
	return strings.Contains(g.hangman.Used, string(letter))
}

func isValidWord(w string) bool {
	if w == "" {
		return false
	}
	for _, l := range w {
		if !unicode.IsLetter(l) {
			return false
		}
	}
	return true
}

func isValidLetter(l rune) bool {
	return unicode.IsLetter(l)
}

func isValidTurns(t int) bool {
	return t > 0
}
