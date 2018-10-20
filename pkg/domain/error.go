package domain

import (
	"github.com/pkg/errors"
)

type ErrInvalidWord error

func errInvalidWord(word string) error {
	return ErrInvalidWord(
		errors.Errorf(`The word '%v' is not valid.`, word))
}

type ErrInvalidTurns error

func errInvalidTurns(turns int) error {
	return ErrInvalidTurns(
		errors.Errorf(`The number of turns '%v' is not valid.`, turns))
}

type ErrAlreadyUsed error

func errAlreadyUsed(r rune) error {
	return ErrInvalidWord(
		errors.Errorf(`You have previously used the letter '%v'.`, r))
}

type ErrGameOver error

func errGameOver() error {
	return ErrGameOver(
		errors.Errorf(`The game is over.`))
}

type ErrInvalidLetter error

func errInvalidLetter(letter rune) error {
	return ErrInvalidLetter(
		errors.Errorf(`The letter '%v' is not valid.`, letter))
}

type ErrWrongVersion error

func errWrongVersion(id string, expVer int, actVer int) error {
	return ErrWrongVersion(
		errors.Errorf("Error saving object '%v'. Expected version %v, actual version %v", id, expVer, actVer))
}
