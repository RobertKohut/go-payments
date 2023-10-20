package tenants

import (
	"github.com/robertkohut/go-payments/pkg/payments"
	pb "github.com/robertkohut/go-payments/proto"
)

type Service interface {
	AddTenant(agent *pb.UserAgent, tenant *pb.Tenant) (string, error)
}

type service struct {
	paymentSvc payments.PaymentService
}

func NewService(payments payments.PaymentService) Service {
	return &service{
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
