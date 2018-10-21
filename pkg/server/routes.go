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
	relHangmanJoin       = "hangman:join"
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
		relHangmanJoin: &route{
			path:    "/hangman/:id/:version",
			method:  "GET",
			handler: hangmanJoin,
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

// Link represents a HAL-like (http://stateless.co/hal_specification.html) structure to specify associated links in a REST-response.
type Link struct {
	Href   string
	Method string
}

// Links generates the requested route name (rels) as Links. Uri template replacements are made with any values provided in the vals map.
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
