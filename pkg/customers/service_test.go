package customers

import (
	"github.com/robertkohut/go-payments/internal/config"
	db "github.com/robertkohut/go-payments/internal/services/repository"
	"github.com/robertkohut/go-payments/pkg/metadata"
	"github.com/robertkohut/go-payments/pkg/payments"
	pb "github.com/robertkohut/go-payments/proto"
	"testing"
)

func setupServices() (Service, error) {
	conf := config.GetConfig("../..")
	db, err := db.DBConnect(conf.DB)
	if err != nil {
		return nil, err
	}

	s := NewService(
		payments.NewStripeService(conf.Stripe),
		NewRepository(db, nil),
	)

	return s, err
}

func TestAddCustomer(t *testing.T) {
	service, err := setupServices()
	if err != nil {
		t.Fatalf("Could not setup services: %v", err)
	}

	customer := &pb.Customer{
		SourceId:  metadata.PaymentSourceStripe,
		AccountId: 55,
		Name:      "Test Customer",
	}

	_, err = service.AddCustomer(customer)
	if err != nil {
		t.Fatalf("Could not add customer: %v", err)
	}

	t.Logf("customer id: %v", customer.ExtId)
}

func TestDeleteCustomer(t *testing.T) {
	service, err := setupServices()
	if err != nil {
		t.Fatalf("Could not setup services: %v", err)
	}

	customer := &pb.Customer{
		SourceId:  1,
		AccountId: 55,
		ExtId:     "cus_O0vtZdCIxvw98R",
	}

	err = service.DeleteCustomer(customer)
	if err != nil {
		t.Fatalf("Could not delete customer: %v", err)
	}

}
