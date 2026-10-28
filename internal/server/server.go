package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	HTTPServer *http.Server
	Logger     *log.Logger
}

func New(logger *log.Logger) *Server {
	if logger == nil {
		logger = log.Default()
	}

	router := http.NewServeMux()

	router.HandleFunc("/", handlers.GetForm)
	router.HandleFunc("/upload", handlers.UploadFile)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: httpServer,
	}
}
