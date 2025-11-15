package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "sprint6: ", log.LstdFlags)

	srv := server.NewHTTPServer(logger)
	if err := srv.Start(); err != nil {
		logger.Fatal(err)
	}

}
