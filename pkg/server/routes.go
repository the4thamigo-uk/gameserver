package server

import (
	"net/http"
	"regexp"
)

type route struct {
	path        string
	uriTmpl     string
	method      string
	handler     handler
	title       string
	linksRels   []string
	linksRoutes routes
}

type routes map[string]*route

var uriTmplRegex = regexp.MustCompile(`:([^/]*)`)

const (
	relIndex             = "index"
	relHangmanList       = "hangman:list"
	relHangmanJoin       = "hangman:join"
	relHangmanCreate     = "hangman:create"
	relHangmanPlayLetter = "hangman:play:letter"
	relHangmanPlayWord   = "hangman:play:word"
)

func newRoutes() routes {
	playRels := []string{
		relHangmanPlayLetter,
		relHangmanPlayWord,
		relHangmanCreate,
		relHangmanList,
	}
	rs := routes{
		relIndex: &route{
			title:   "Game server",
			path:    "/",
			method:  http.MethodGet,
			handler: rootIndex,
			linksRels: []string{
				relHangmanList,
			},
		},
		relHangmanList: &route{
			title:     "Hangman game list",
			path:      "/hangman",
			method:    http.MethodGet,
			handler:   hangmanList,
			linksRels: append(playRels, relHangmanJoin),
		},
		relHangmanJoin: &route{
			title:     "Join hangman game",
			path:      "/hangman/:id/:version",
			method:    http.MethodGet,
			handler:   hangmanJoin,
			linksRels: playRels,
		},
		relHangmanCreate: &route{
			title:     "Create hangman game",
			path:      "/hangman/create",
			method:    http.MethodPost,
			handler:   hangmanCreate,
			linksRels: playRels,
		},
		relHangmanPlayLetter: &route{
			title:     "Guess a letter",
			path:      "/hangman/:id/:version/letter/:letter",
			method:    http.MethodPatch,
			handler:   hangmanPlayLetter,
			linksRels: playRels,
		},
		relHangmanPlayWord: &route{
			title:     "Guess the word",
			path:      "/hangman/:id/:version/word/:word",
			method:    http.MethodPatch,
			handler:   hangmanPlayWord,
			linksRels: playRels,
		},
	}
	for _, r := range rs {
		r.init(rs)
	}
	return rs
}

func (r *route) init(rs routes) {
	r.uriTmpl = uriTmplRegex.ReplaceAllString(r.path, `{$1}`)

	lrs := routes{}
	for _, lrel := range r.linksRels {
		lr, ok := rs[lrel]
		if !ok {
			continue
		}
		lrs[lrel] = lr
	}
	r.linksRoutes = lrs
}
