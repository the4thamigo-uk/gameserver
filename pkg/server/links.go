package server

import (
	"fmt"
	"strings"
)

type link struct {
	Href   string
	Method string
}

type links map[string]*link

func newLinks(rs routes, vals map[string]interface{}) links {
	if len(rs) == 0 {
		return nil
	}
	ls := links{}
	for rel, r := range rs {
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
	}
}
