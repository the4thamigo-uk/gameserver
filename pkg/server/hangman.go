package server

import (
	"github.com/pkg/errors"
	"github.com/the4thamigo-uk/gameserver/pkg/domain"
	"github.com/the4thamigo-uk/gameserver/pkg/store"
	"math/rand"
	"strconv"
)

const hangmanDefaultTurns = 6

func hangmanCreate(r *request, g *globals) (*response, error) {
	cfg := g.cfg.Hangman
	word := cfg.dict.GetAt(rand.Int())
	res, err := domain.CreateHangman(g.store, word, hangmanDefaultTurns)
	if err != nil {
		return nil, err
	}
	return hangmanResponse(r.rt, res), nil
}

func hangmanList(r *request, g *globals) (*response, error) {
	res, err := domain.ListHangman(g.store)
	if err != nil {
		return nil, err
	}
	return &response{
		Games: res,
		Links: linksForRoute(
			r.rt,
			nil,
		)}, nil
}

func hangmanJoin(r *request, g *globals) (*response, error) {
	gid, err := gameID(r)
	if err != nil {
		return nil, err
	}
	res, err := domain.JoinHangman(g.store, *gid)
	if err != nil {
		return nil, err
	}
	return hangmanResponse(r.rt, res), nil
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
	return hangmanResponse(r.rt, res), nil
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
	return hangmanResponse(r.rt, res), nil
}

func gameID(r *request) (*store.ID, error) {
	id := r.p.ByName("id")

	ver := 0
	if s := r.p.ByName("version"); s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.Wrap(err, "Version must be an integer")
		}
		ver = i
	}
	gid := store.NewID(id, ver)
	return &gid, nil
}

func hangmanResponse(rt *route, res *domain.HangmanResult) *response {
	ls := linksForRoute(
		rt,
		map[string]interface{}{
			"id":      res.ID.ID,
			"version": res.ID.Version,
		},
	)
	if res.State != domain.Play {
		delete(ls, relHangmanPlayLetter)
		delete(ls, relHangmanPlayWord)
	}
	return &response{
		Game:  res,
		Links: ls,
	}
}
