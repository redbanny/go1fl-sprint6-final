package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "[morse]", log.LstdFlags|log.Lmsgprefix)

	srv := server.New(logger)

	dir, err := os.Getwd()
	if err != nil {
		return
	}

	logger.Println(dir)

	logger.Println("Server started at", srv.HTTPServer.Addr)
	err = srv.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
