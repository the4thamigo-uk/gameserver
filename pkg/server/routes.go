package server

import (
	"fmt"
	"regexp"
	"strings"
)

type route struct {
	path    string
	method  string
	handler handler
	uriTmpl string
}

type routes map[string]*route

var uriTmplRegex = regexp.MustCompile(`:([^/]*)`)

const (
	relIndex             = "index"
	relHangmanList       = "hangman:list"
	relHangmanCreate     = "hangman:create"
	relHangmanPlayLetter = "hangman:play:letter"
	relHangmanPlayWord   = "hangman:play:word"
)

func newRoutes() routes {
	routes := routes{
		relIndex: &route{
			path:    "/",
			method:  "GET",
			handler: rootIndex,
		},
		relHangmanList: &route{
			path:    "/hangman",
			method:  "GET",
			handler: hangmanList,
		},
		relHangmanCreate: &route{
			path:    "/hangman/create",
			method:  "POST",
			handler: hangmanCreate,
		},
		relHangmanPlayLetter: &route{
			path:    "/hangman/:id/:version/letter/:letter",
			method:  "PATCH",
			handler: hangmanPlayLetter,
		},
		relHangmanPlayWord: &route{
			path:    "/hangman/:id/:version/word/:word",
			method:  "PATCH",
			handler: hangmanPlayWord,
		},
	}
	for _, route := range routes {
		route.uriTmpl = uriTmplRegex.ReplaceAllString(route.path, `{$1}`)
	}
	return routes
}

type Link struct {
	Href   string
	Rel    string
	Method string
}

func (rs routes) Links(rels []string, vals map[string]interface{}) map[string]Link {
	links := map[string]Link{}
	for _, rel := range rels {
		r, ok := rs[rel]
		if !ok {
			continue
		}
		href := r.uriTmpl
		for k, v := range vals {
			sv := fmt.Sprintf("%v", v)
			href = strings.Replace(href, `{`+k+`}`, sv, -1)
		}
		links[rel] = Link{
			Href:   href,
			Method: r.method,
		}

	}
	return links
}
