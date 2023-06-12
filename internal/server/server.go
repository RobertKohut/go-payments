package server

import (
	"github.com/robertkohut/go-payments/internal/services/hashid"
	"github.com/robertkohut/go-payments/pkg/charges"
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

	hashIdService, _ := hashid.New(&cfg.HashId)

	ps := payments.NewService("stripe", cfg)
	customerSvc := customers.NewService(ps, customers.NewRepository(db))
	chargesSvc := charges.NewService(ps, charges.NewRepository(db), hashIdService)

	return &Server{
		config: cfg,
		svc: &services.Services{
			DB:          db,
			HashId:      hashIdService,
			CustomerSvc: customerSvc,
			ChargeSvc:   chargesSvc,
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
