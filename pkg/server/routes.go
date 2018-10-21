package server

import (
	"regexp"
)

type route struct {
	path        string
	uriTmpl     string
	method      string
	handler     handler
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
	}
	rs := routes{
		relIndex: &route{
			path:    "/",
			method:  "GET",
			handler: rootIndex,
			linksRels: []string{
				relHangmanList,
			},
		},
		relHangmanList: &route{
			path:      "/hangman",
			method:    "GET",
			handler:   hangmanList,
			linksRels: append(playRels, relHangmanJoin),
		},
		relHangmanJoin: &route{
			path:      "/hangman/:id/:version",
			method:    "GET",
			handler:   hangmanJoin,
			linksRels: playRels,
		},
		relHangmanCreate: &route{
			path:      "/hangman/create",
			method:    "POST",
			handler:   hangmanCreate,
			linksRels: playRels,
		},
		relHangmanPlayLetter: &route{
			path:      "/hangman/:id/:version/letter/:letter",
			method:    "PATCH",
			handler:   hangmanPlayLetter,
			linksRels: playRels,
		},
		relHangmanPlayWord: &route{
			path:      "/hangman/:id/:version/word/:word",
			method:    "PATCH",
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
