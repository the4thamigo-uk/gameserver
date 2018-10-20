package domain

import (
	"github.com/pkg/errors"
)

// ErrGameOver indicates that the game has finished and no more play can take place.
type ErrGameOver error

func errGameOver() error {
	return ErrGameOver(
		errors.Errorf(`The game is over.`))
}

// ErrInvalidTurns indicates that the specified number of turns is not acceptable.
type ErrInvalidTurns error

func errInvalidTurns(turns int) error {
	return ErrInvalidTurns(
		errors.Errorf(`The number of turns '%v' is not valid.`, turns))
}

// ErrInvalidWord indicates that a word is not valid choice for the game
type ErrInvalidWord error

func errInvalidWord(word string) error {
	return ErrInvalidWord(
		errors.Errorf(`The word '%v' is not valid.`, word))
}

// ErrInvalidLetter indicates that a letter is not valid choice for the game
type ErrInvalidLetter error

func errInvalidLetter(letter rune) error {
	return ErrInvalidLetter(
		errors.Errorf(`The letter '%v' is not valid.`, letter))
}

// ErrAlreadyUsed indicates that the letter choice has previously been used.
type ErrAlreadyUsed error

func errAlreadyUsed(r rune) error {
	return ErrInvalidWord(
		errors.Errorf(`You have previously used the letter '%v'.`, r))
}
