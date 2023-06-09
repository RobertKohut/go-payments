package payments

import (
	"github.com/robertkohut/go-payments/internal/config"
	pb "github.com/robertkohut/go-payments/proto"
)

type PaymentService interface {
	CreateCustomer(customer *pb.Customer) (string, error)
	DeleteCustomer(customer *pb.Customer) error

	AddCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) (*pb.Card, error)
	RemoveCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) error
	CreateCharge(customer *pb.Customer, card *pb.Card, amount int64) error
}

func NewService(gateway string, cfg *config.Configuration) PaymentService {
	switch gateway {
	case "stripe":
		return NewStripeService(cfg.Stripe.SecretKey)
	default:
		return nil
	}
}
