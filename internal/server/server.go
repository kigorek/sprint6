package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type HTTPServer struct {
	logs *log.Logger
	hs   http.Server
}

func NewHTTPServer(logs *log.Logger) *HTTPServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	SERV := HTTPServer{
		hs: http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     logs,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
		logs: logs,
	}

	return &SERV
}
func (s *HTTPServer) Start() error {
	s.logs.Println("HTTPServer starting on", s.hs.Addr)
	return s.hs.ListenAndServe()
}
