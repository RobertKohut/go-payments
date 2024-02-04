package payments

import (
	"context"
	"errors"
	"github.com/robertkohut/go-payments/internal/config"
	pb "github.com/robertkohut/go-payments/proto"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
	"log"
	"time"
)

type stripeService struct {
	publishableKey string
	client         *client.API
}

func NewStripeService(cfg *config.StripeConfig) PaymentService {
	return &stripeService{
		publishableKey: cfg.PublishableKey,
		client:         client.New(cfg.SecretKey, nil),
	}
}

func (s *stripeService) GetPublishableKey() (string, error) {
	return s.publishableKey, nil
}

func (s *stripeService) validateBusinessProfile(b *pb.BusinessProfile) error {
	if b.GetCountry() == "" {
		return errors.New("country is required")
	}

	businessType := b.GetBusinessType()
	if businessType == "" {
		return errors.New("business type is required")
	}

	if businessType != "individual" && businessType != "company" && businessType != "non_profit" {
		return errors.New("business type must be individual, company or non_profit")
	}

	if b.GetName() == "" {
		return errors.New("company name is required")
	}

	return nil
}

func (s *stripeService) getStripeBusinessType(businessType string) (stripe.AccountBusinessType, error) {
	switch businessType {
	case "individual":
		return stripe.AccountBusinessTypeIndividual, nil
	case "company":
		return stripe.AccountBusinessTypeCompany, nil
	case "non_profit":
		return stripe.AccountBusinessTypeNonProfit, nil
	default:
		return "", errors.New("invalid business type")
	}
}

func (s *stripeService) CreateAccount(agent *pb.UserAgent, tenant *pb.Tenant) (string, error) {
	bp := tenant.GetBusinessProfile()

	err := s.validateBusinessProfile(bp)
	if err != nil {
		return "", err
	}

	if tenant.GetTosAccepted() == false {
		return "", errors.New("terms of service must be accepted")
	}

	businessType, err := s.getStripeBusinessType(bp.GetBusinessType())
	if err != nil {
		return "", err
	}

	params := &stripe.AccountParams{
		Country:      stripe.String(bp.GetCountry()),
		BusinessType: stripe.String(string(businessType)),
		Company: &stripe.AccountCompanyParams{
			Name: stripe.String(tenant.GetLegalName()),
			Address: &stripe.AddressParams{
				Line1:      stripe.String(bp.GetAddress().GetLine1()),
				Line2:      stripe.String(bp.GetAddress().GetLine2()),
				City:       stripe.String(bp.GetAddress().GetCity()),
				State:      stripe.String(bp.GetAddress().GetState()),
				PostalCode: stripe.String(bp.GetAddress().GetPostalCode()),
				Country:    stripe.String(bp.GetCountry()),
			},
			Phone: stripe.String(bp.GetPhone()),
			TaxID: stripe.String(bp.GetTaxId()),
		},
		BusinessProfile: &stripe.AccountBusinessProfileParams{
			Name: stripe.String(bp.Name),
		},
		Type: stripe.String(string(stripe.AccountTypeCustom)),
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
		TOSAcceptance: &stripe.AccountTOSAcceptanceParams{
			Date:      stripe.Int64(time.Now().Unix()),
			IP:        stripe.String(agent.GetIp()),
			UserAgent: stripe.String(agent.GetUserAgent()),
		},
	}

	a, err := s.client.Accounts.New(params)
	if err != nil {
		return "", err
	}

	log.Println("Created stripe account: ", a.ID)

	return a.ID, nil
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

func (s *stripeService) CreateCharge(ctx context.Context, customer *pb.Customer, card *pb.Card, charge *pb.Charge) (*string, error) {
	tenant := ctx.Value("tenant-gateway").(string)

	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(charge.GetAmount()),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethod: stripe.String(card.GetExtId()),
		Description:   stripe.String(charge.GetDescription()),
	}

	if tenant != "" {
		params.ApplicationFeeAmount = stripe.Int64(200)
		params.Params = stripe.Params{
			StripeAccount: stripe.String(tenant),
		}

		params.PaymentMethodTypes = stripe.StringSlice([]string{
			"card",
		})
	} else {
		params.Customer = stripe.String(customer.GetExtId())
	}

	pi, err := s.client.PaymentIntents.New(params)
	if err != nil {
		return nil, err
	}

	log.Println("Created stripe payment intent: ", pi.ID)

	confirmParams := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String(card.GetExtId()), // Card ID.
	}

	if tenant != "" {
		confirmParams.Params = stripe.Params{
			StripeAccount: stripe.String(tenant),
		}
	}

	pi, err = s.client.PaymentIntents.Confirm(pi.ID, confirmParams)
	if err != nil {
		return nil, err
	}

	log.Println("Confirmed stripe payment intent: ", pi.ID)

	return &pi.ID, nil
}
