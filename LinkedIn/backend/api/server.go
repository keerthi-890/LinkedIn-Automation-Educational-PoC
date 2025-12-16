package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"linkedin-automation-poc/core"
	
)


type Server struct {
	Router *chi.Mux
	Bot    *core.Bot
	DB     *sql.DB
}


func NewServer(bot *core.Bot, db *sql.DB) *Server {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	s := &Server{
		Router: r,
		Bot:    bot,
		DB:     db,
	}

	s.routes()
	return s
}


func (s *Server) routes() {
	s.Router.Post("/start", s.handleStart)
	s.Router.Post("/stop", s.handleStop)
	s.Router.Get("/stats", s.handleStats)
	s.Router.Get("/logs", s.handleLogs)
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.Router)
}

func (s *Server) handleStart(w http.ResponseWriter, r *http.Request) {
	go s.Bot.Start()
	json.NewEncoder(w).Encode(map[string]string{"status": "started"})
}

func (s *Server) handleStop(w http.ResponseWriter, r *http.Request) {
	s.Bot.Stop()
	json.NewEncoder(w).Encode(map[string]string{"status": "stopped"})
}

func (s *Server) handleStats(w http.ResponseWriter, r *http.Request) {
	// Mock stats for display if DB isn't fully populated
	stats := map[string]int{
		"profiles_found": 15,
		"requests_sent":  5,
		"connected":      2,
		"messages_sent":  1,
	}
	// Real stats would come from s.DB.GetStats()
	json.NewEncoder(w).Encode(stats)
}

func (s *Server) handleLogs(w http.ResponseWriter, r *http.Request) {
	logs := []string{
		"2023-10-27 10:00:01 - Started automation",
		"2023-10-27 10:00:05 - Found 4 profiles",
		"2023-10-27 10:00:12 - Sent connection request to Sarah Engineer",
	}
	json.NewEncoder(w).Encode(logs)
}
