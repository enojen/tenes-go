package main

import (
	"log"
	"tenes-go/internal/server"
	"tenes-go/pkg/config"
)

func main() {
	cfg := config.Load()
	srv := server.NewServer()

	if err := srv.Start(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
