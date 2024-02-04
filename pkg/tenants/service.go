package tenants

import (
	"github.com/robertkohut/go-payments/pkg/payments"
	pb "github.com/robertkohut/go-payments/proto"
)

type Service interface {
	AddTenant(agent *pb.UserAgent, tenant *pb.Tenant) (string, error)
	GetTenantByApiKey(source, apiKey, gateway string) (*pb.Tenant, error)
}

type service struct {
	repo       Repository
	paymentSvc payments.PaymentService
}

func NewService(repo Repository, payments payments.PaymentService) Service {
	return &service{
		repo:       repo,
		paymentSvc: payments,
	}
}

func (s *service) AddTenant(agent *pb.UserAgent, tenant *pb.Tenant) (string, error) {
	tenantExtId, err := s.paymentSvc.CreateAccount(agent, tenant)
	if err != nil {
		return "", err
	}

	return tenantExtId, nil
}

func (s *service) GetTenantByApiKey(source, apiKey, gateway string) (*pb.Tenant, error) {
	tenant, err := s.repo.SelectTenantByApiKey(source, apiKey, gateway)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}
