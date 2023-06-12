package payments

import (
	pb "github.com/robertkohut/go-payments/proto"
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

func (s *stripeService) CreateCustomer(customer *pb.Customer) (string, error) {
	params := &stripe.CustomerParams{
		Description: stripe.String(customer.Name),
	}

	c, err := s.client.Customers.New(params)
	if err != nil {
		return "", err
	}

	return c.ID, nil
}

func (s *stripeService) DeleteCustomer(customer *pb.Customer) error {
	c, err := s.client.Customers.Del(customer.ExtId, nil)
	if err != nil {
		return err
	}

	log.Println("Deleted stripe customer: ", c.ID)

	return nil
}

func (s *stripeService) AddCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) (*pb.Card, error) {
	pm, err := s.client.PaymentMethods.Attach(
		card.GetExtId(),
		&stripe.PaymentMethodAttachParams{
			Customer: stripe.String(customer.GetExtId()),
		})

	if err != nil {
		return nil, err
	}

	card = &pb.Card{
		Brand: string(pm.Card.Brand),
		Last4: pm.Card.Last4,
	}

	return card, nil
}

func (s *stripeService) RemoveCustomerPaymentMethod(_ *pb.Customer, card *pb.Card) error {
	_, err := s.client.PaymentMethods.Detach(
		card.GetExtId(),
		&stripe.PaymentMethodDetachParams{},
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *stripeService) CreateCharge(customer *pb.Customer, card *pb.Card, charge *pb.Charge) (*string, error) {
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(charge.GetAmount()),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Customer:      stripe.String(customer.GetExtId()),
		PaymentMethod: stripe.String(card.GetExtId()),
		Description:   stripe.String(charge.GetDescription()),
	}

	pi, err := s.client.PaymentIntents.New(params)
	if err != nil {
		return nil, err
	}

	log.Println("Created stripe payment intent: ", pi.ID)

	confirmParams := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String(card.GetExtId()), // Card ID.
	}

	pi, err = s.client.PaymentIntents.Confirm(pi.ID, confirmParams)
	if err != nil {
		return nil, err
	}

	log.Println("Confirmed stripe payment intent: ", pi.ID)

	return &pi.ID, nil
}
