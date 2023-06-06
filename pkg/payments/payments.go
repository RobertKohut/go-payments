package payments

import (
	"github.com/robertkohut/go-payments/internal/config"
	"github.com/robertkohut/go-payments/pkg/entities"
)

type PaymentService interface {
	CreateCustomer(customer *entities.Customer) (string, error)
	DeleteCustomer(customer *entities.Customer) error
}

func NewService(gateway string, cfg *config.Configuration) PaymentService {
	switch gateway {
	case "stripe":
		return NewStripeService(cfg.Stripe.SecretKey)
	default:
		return nil
	}
}
