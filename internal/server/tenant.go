package server

import (
	"context"
	"errors"
	pb "github.com/robertkohut/go-payments/proto"
	"log"
)

func (s *Server) CreateTenant(ctx context.Context, req *pb.CreateTenantRequest) (*pb.CreateTenantResponse, error) {
	bp := req.GetBusinessProfile()

	tenant := &pb.Tenant{
		BusinessProfile: bp,
		TosAccepted:     req.GetTosAccepted(),
	}

	userAgent := req.GetUserAgent()
	if userAgent == nil {
		return nil, errors.New("user agent is required")
	}

	_, err := s.svc.TenantSvc.AddTenant(userAgent, tenant)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateTenantResponse{
		Id: 1,
	}

	return resp, nil
}

func (s *Server) ValidateTenantApiKey(apiKey string) bool {
	if apiKey == "cet" {
		return true
	}

	log.Println("Invalid API key provided")

	return false
}

func (s *Server) GetTenantExtId(ctx context.Context) (string, error) {
	return "acct_1OeVDaPtwlB23Hm0", nil
}
