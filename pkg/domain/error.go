package domain

import (
	"github.com/pkg/errors"
)

// ErrGameOver indicates that the game has finished and no more play can take place.
type ErrGameOver struct {
	error
}

func errGameOver() error {
	return ErrGameOver{
		errors.Errorf(`The game is over.`)}
}

// ErrInvalidTurns indicates that the specified number of turns is not acceptable.
type ErrInvalidTurns struct {
	error
}

func errInvalidTurns(turns int) error {
	return ErrInvalidTurns{
		errors.Errorf(`The number of turns '%v' is not valid.`, turns)}
}

// ErrInvalidWord indicates that a word is not valid choice for the game
type ErrInvalidWord struct {
	error
}

func errInvalidWord(word string) error {
	return ErrInvalidWord{
		errors.Errorf(`The word '%v' is not valid.`, word)}
}

// ErrInvalidLetter indicates that a letter is not valid choice for the game
type ErrInvalidLetter struct {
	error
}

func errInvalidLetter(letter rune) error {
	return ErrInvalidLetter{
		errors.Errorf(`The letter '%v' is not valid.`, letter)}
}

// ErrAlreadyUsed indicates that the letter choice has previously been used.
type ErrAlreadyUsed struct {
	error
}

func errAlreadyUsed(r rune) error {
	return ErrAlreadyUsed{
		errors.Errorf(`You have previously used the letter '%c'.`, r)}
}
