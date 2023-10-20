package payments

import (
	"github.com/robertkohut/go-payments/internal/config"
	pb "github.com/robertkohut/go-payments/proto"
)

type PaymentService interface {
	GetPublishableKey() (string, error)

	CreateAccount(agent *pb.UserAgent, tenant *pb.Tenant) (string, error)

	CreateCustomer(customer *pb.Customer) (string, error)
	DeleteCustomer(customer *pb.Customer) error

	AddCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) (*pb.Card, error)
	RemoveCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) error
	CreateCharge(customer *pb.Customer, card *pb.Card, charge *pb.Charge) (*string, error)
}

func NewService(gateway string, cfg *config.Configuration) PaymentService {
	switch gateway {
	case "stripe":
		return NewStripeService(cfg.Stripe)
	default:
		return nil
	}
}
