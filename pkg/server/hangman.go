package server

import (
	"github.com/pkg/errors"
	"github.com/the4thamigo_uk/gameserver/pkg/domain"
	"github.com/the4thamigo_uk/gameserver/pkg/store"
	"math/rand"
	"strconv"
)

const hangmanDefaultTurns = 6

func hangmanList(r *request, g *globals) (*response, error) {
	return nil, nil
}

func hangmanCreate(r *request, g *globals) (*response, error) {
	cfg := g.cfg.Hangman
	word := cfg.dict.GetAt(rand.Int())
	res, err := domain.CreateHangman(g.store, word, hangmanDefaultTurns)
	if err != nil {
		return nil, err
	}
	return &response{
		State: res,
		Links: g.routes.Links(
			[]string{
				relHangmanPlayLetter,
				relHangmanPlayWord,
			},
			map[string]interface{}{
				"id":      res.ID.ID,
				"version": res.ID.Version,
			},
		)}, nil
}

func gameID(r *request) (*store.ID, error) {
	id := r.p.ByName("id")
	ver, err := strconv.Atoi(r.p.ByName("version"))
	if err != nil {
		return nil, errors.Wrap(err, "Version must be an integer")
	}
	return &store.ID{ID: id, Version: ver}, nil
}

func hangmanPlayLetter(r *request, g *globals) (*response, error) {
	gid, err := gameID(r)
	if err != nil {
		return nil, err
	}
	l := []rune(r.p.ByName("letter"))
	if len(l) != 1 {
		return nil, errors.New("Letter must be a single character")
	}
	res, err := domain.PlayLetter(g.store, *gid, l[0])
	if err != nil {
		return nil, err
	}
	return &response{
		State: res,
		Links: g.routes.Links(
			[]string{
				relHangmanPlayLetter,
				relHangmanPlayWord,
			},
			map[string]interface{}{
				"id":      res.ID.ID,
				"version": res.ID.Version,
			},
		)}, nil
}

func hangmanPlayWord(r *request, g *globals) (*response, error) {
	gid, err := gameID(r)
	if err != nil {
		return nil, err
	}
	word := r.p.ByName("word")
	res, err := domain.PlayWord(g.store, *gid, word)
	if err != nil {
		return nil, err
	}
	return &response{
		State: res,
		Links: g.routes.Links(
			[]string{
				relHangmanPlayLetter,
				relHangmanPlayWord,
			},
			map[string]interface{}{
				"id":      res.ID.ID,
				"version": res.ID.Version,
			},
		)}, nil
}
