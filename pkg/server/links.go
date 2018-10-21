package server

import (
	"fmt"
	"strings"
)

// Link represents a HAL-like (http://stateless.co/hal_specification.html) structure to specify associated links in a REST-response.
type link struct {
	Href   string
	Method string
}

type links map[string]link

// Links generates the requested route name (rels) as Links. Uri template replacements are made with any values provided in the vals map.
func newLinks(rs routes, vals map[string]interface{}) links {
	links := links{}
	for rel, r := range rs {
		href := r.uriTmpl
		for k, v := range vals {
			sv := fmt.Sprintf("%v", v)
			href = strings.Replace(href, `{`+k+`}`, sv, -1)
		}
		links[rel] = link{
			Href:   href,
			Method: r.method,
		}

	}
	return links
}
