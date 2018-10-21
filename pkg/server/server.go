package server

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/the4thamigo-uk/gameserver/pkg/store"
	"github.com/the4thamigo-uk/gameserver/pkg/store/memorystore"
	"net/http"
	"time"
)

// Server is an instance of a REST gameserver.
type Server struct {
	s *http.Server
	g *globals
}

// NewServer creates a new instance of Server with default in-memory storage,
func NewServer(cfg *Config) *Server {
	g := &globals{
		cfg:    cfg,
		routes: newRoutes(),
		store:  memorystore.New(),
	}

	r := httprouter.New()
	for _, route := range g.routes {
		r.Handle(route.method, route.path, rootHandler(route.handler, g))
	}
	return &Server{
		s: &http.Server{
			Addr:    cfg.Address,
			Handler: r,
		},
		g: g,
	}
}

// WithStore overrides the default in-memory storage with a client-specified storage.
func (s *Server) WithStore(store store.Store) *Server {
	s.g.store = store
	return s
}

// ListenAndServe starts the server. This function blocks until either Shutdown is called or the process is terminated.
func (s *Server) ListenAndServe() error {
	return s.s.ListenAndServe()
}

// Shutdown attempts a graceful shutdown of the server, timing out after 5 seconds.
func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.s.Shutdown(ctx)
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
