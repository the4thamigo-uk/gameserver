package server

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/the4thamigo_uk/gameserver/pkg/domain"
	"github.com/the4thamigo_uk/gameserver/pkg/store"
	"github.com/the4thamigo_uk/gameserver/pkg/store/memorystore"
	"net/http"
	"time"
)

type Server struct {
	s *http.Server
	g *globals
}

func NewServer(addr string) *Server {
	g := &globals{
		routes: newRoutes(),
		store:  memorystore.NewMemoryStore(),
	}

	r := httprouter.New()
	for _, route := range g.routes {
		r.Handle(route.method, route.path, rootHandler(route.handler, g))
	}
	return &Server{
		s: &http.Server{
			Addr:    addr,
			Handler: r,
		},
		g: g,
	}
}

func (s *Server) WithStore(store store.Store) *Server {
	s.g.store = store
	return s
}

func (s *Server) WithDictionary(dict *domain.Dictionary) *Server {
	s.g.dict = dict
	return s
}

func (s *Server) ListenAndServe() error {
	return s.s.ListenAndServe()
}

func (s *Server) Shutdown() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.s.Shutdown(ctx)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.s.Handler.ServeHTTP(w, r)
}

func rootHandler(h handler, g *globals) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := &request{
			r: r,
			w: w,
			p: p,
		}
		rsp, err := h(ctx, g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		b, err := json.Marshal(rsp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/hal+json")
		_, err = w.Write(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
