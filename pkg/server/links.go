package server

import (
	"fmt"
	"strings"
)

type link struct {
	Title  string
	Href   string
	Method string
}

type links map[string]*link

func linksForRoute(r *route, vals map[string]interface{}) links {
	ls := links{
		"self": newLink(r, vals),
	}
	for rel, r := range r.linksRoutes {
		ls[rel] = newLink(r, vals)
	}
	return ls
}

func newLink(r *route, vals map[string]interface{}) *link {
	href := r.uriTmpl
	for k, v := range vals {
		sv := fmt.Sprintf("%v", v)
		href = strings.Replace(href, `{`+k+`}`, sv, -1)
	}
	return &link{
		Href:   href,
		Method: r.method,
		Title:  r.title,
	}
}
