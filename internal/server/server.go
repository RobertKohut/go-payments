package server

import (
	"context"
	"github.com/robertkohut/go-payments/internal/services/hashid"
	"github.com/robertkohut/go-payments/pkg/charges"
	"github.com/robertkohut/go-payments/pkg/customers"
	"github.com/robertkohut/go-payments/pkg/payments"
	"github.com/robertkohut/go-payments/pkg/tenants"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"

	pb "github.com/robertkohut/go-payments/proto"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
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
	tenantSvc := tenants.NewService(ps)
	customerSvc := customers.NewService(ps, customers.NewRepository(db, hashIdService))
	chargesSvc := charges.NewService(ps, charges.NewRepository(db, hashIdService), hashIdService)

	return &Server{
		config: cfg,
		svc: &services.Services{
			DB:          db,
			HashId:      hashIdService,
			TenantSvc:   tenantSvc,
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

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			s.loggingInterceptor,
			s.apiKeyInterceptor,
		)),
	)

	pb.RegisterPaymentServiceServer(server, s)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return nil
}

func (s *Server) apiKeyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing required metadata")
	}

	apiKeys, ok := md["api-key"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "API key is missing")
	}

	apiKey := apiKeys[0]

	if !s.ValidateTenantApiKey(apiKey) {
		return nil, status.Errorf(codes.PermissionDenied, "invalid API key")
	}

	return handler(ctx, req)
}

func (s *Server) loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("gRPC request: %s", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("gRPC error: %v", err)
		status, _ := status.FromError(err)
		return nil, status.Err()
	}
	return resp, nil
}

func (s *Server) CloseDB() error {
	return s.svc.DB.Close()
}
