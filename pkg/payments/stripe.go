package payments

import (
	"github.com/robertkohut/go-payments/pkg/entities"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
	"log"
)

type stripeService struct {
	client *client.API
}

func NewStripeService(key string) PaymentService {
	return &stripeService{
		client: client.New(key, nil),
	}
}

func (s *stripeService) CreateCustomer(customer *entities.Customer) (string, error) {
	params := &stripe.CustomerParams{
		Description: stripe.String(customer.Name),
	}

	c, err := s.client.Customers.New(params)
	if err != nil {
		return "", err
	}

	return c.ID, nil
}

func (s *stripeService) DeleteCustomer(customer *entities.Customer) error {
	c, err := s.client.Customers.Del(customer.ExtId, nil)
	if err != nil {
		return err
	}

	log.Println("Deleted stripe customer: ", c.ID)

	return nil
}
