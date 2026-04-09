package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	serv := server.NewServer(logger)

	logger.Println("server start")

	err := serv.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatal("server dont start", err)
	}

}
