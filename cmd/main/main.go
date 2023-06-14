package main

import (
	"flag"
	"github.com/robertkohut/go-payments/internal/config"
	"github.com/robertkohut/go-payments/internal/server"
	"log"
)

var (
	configPath string
)

func init() {
	flag.Parse()
	flag.StringVar(&configPath, "config", "", "Config path")
}

func main() {
	cfg := config.GetConfig(configPath)

	s := server.NewServer(cfg)
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
