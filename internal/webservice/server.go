package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jpogah/testforgeui/internal/config"
)

type Server struct {
	cfg    config.Config
	router *chi.Mux
}

func NewServer(cfg config.Config) *Server {
	router := chi.NewRouter()

	s := &Server{
		cfg:    cfg,
		router: router,
	}

	s.registerRoutes()

	return s
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	return http.ListenAndServe(addr, s.router)
}

func (s *Server) registerRoutes() {
	s.router.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":      "ok",
			"environment": s.cfg.Environment,
		})
	})
}
