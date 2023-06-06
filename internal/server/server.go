package server

import (
	"github.com/robertkohut/go-payments/pkg/customers"
	"github.com/robertkohut/go-payments/pkg/payments"
	"log"
	"net"

	pb "github.com/robertkohut/go-payments/proto"

	"github.com/robertkohut/go-payments/internal/config"
	"github.com/robertkohut/go-payments/internal/services"
	"github.com/robertkohut/go-payments/internal/services/repository"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Configuration
	svc    *services.Services
	pb.UnimplementedPaymentServiceServer
}

func NewServer(cfg *config.Configuration) *Server {
	db, err := repository.DBConnect(cfg.DB)
	if err != nil {
		log.Panic("Unable to connect to database")
	}

	ps := payments.NewService("stripe", cfg)
	cs := customers.NewService(ps, customers.NewRepository(db))

	return &Server{
		config: cfg,
		svc: &services.Services{
			DB:          db,
			CustomerSvc: cs,
		},
	}
}

func (s *Server) Run() error {
	log.Println("Starting server on port", s.config.App.Addr)

	listener, err := net.Listen("tcp", s.config.App.Addr)
	if err != nil {
		log.Fatalf("Unable to listen on port %s: %v", s.config.App.Addr, err)
	}

	server := grpc.NewServer()

	pb.RegisterPaymentServiceServer(server, s)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return nil
}

func (s *Server) CloseDB() error {
	return s.svc.DB.Close()
}
