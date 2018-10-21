package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/the4thamigo-uk/gameserver/pkg/store"
	"net/http"
)

type handler func(r *request, g *globals) (*response, error)

type request struct {
	r *http.Request
	p httprouter.Params
	w http.ResponseWriter
}

type response struct {
	State interface{}
	Links map[string]Link `json:"_links"`
}

type globals struct {
	cfg    *Config
	routes routes
	store  store.Store
}

func rootIndex(r *request, g *globals) (*response, error) {
	return &response{
		Links: g.routes.Links(
			[]string{
				relHangmanList,
			},
			nil,
		)}, nil
}
