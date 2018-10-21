package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/the4thamigo-uk/gameserver/pkg/store"
	"net/http"
)

type handler func(r *request, g *globals) (*response, error)

type request struct {
	r  *http.Request
	p  httprouter.Params
	w  http.ResponseWriter
	rt *route
}

type response struct {
	Game  interface{} `json:"game,omitempty"`
	Games interface{} `json:"games,omitempty"`
	Links links       `json:"_links,omitempty"`
}

type globals struct {
	cfg    *Config
	routes routes
	store  store.Store
}

func rootIndex(r *request, g *globals) (*response, error) {
	return &response{
		Links: linksForRoute(r.rt, nil),
	}, nil
}
