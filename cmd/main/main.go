package main

import (
	"github.com/robertkohut/go-payments/internal/config"
	"github.com/robertkohut/go-payments/internal/server"
	"log"
)

func main() {
	cfg := config.GetConfig("")

	s := server.NewServer(cfg)
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
