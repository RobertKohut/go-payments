package customers

import (
	"errors"
	"github.com/robertkohut/go-payments/pkg/entities"
	"github.com/robertkohut/go-payments/pkg/payments"
)

type Service interface {
	GetCustomerById(id int64) (*entities.Customer, error)
	AddCustomer(customer *entities.Customer) (*string, error)
	DeleteCustomer(customer *entities.Customer) error
}

type service struct {
	paymentSvc payments.PaymentService
	repo       Repository
}

func NewService(payments payments.PaymentService, repo Repository) Service {
	return &service{
		paymentSvc: payments,
		repo:       repo,
	}
}

func (s *service) GetCustomerById(id int64) (*entities.Customer, error) {
	customer, err := s.repo.SelectCustomerByAccountId(id)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *service) AddCustomer(customer *entities.Customer) (*string, error) {
	customerExtId, err := s.paymentSvc.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	customer.ExtId = customerExtId

	customer.Id, err = s.repo.InsertCustomer(customer)
	if err != nil {
		_ = s.paymentSvc.DeleteCustomer(customer)
		return nil, err
	}

	return &customerExtId, nil
}

func (s *service) DeleteCustomer(customer *entities.Customer) error {
	err := s.repo.DeleteCustomerBySourceId(customer)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) AddCustomerCard(customer *entities.Customer) (*entities.Card, error) {
	return nil, errors.New("not implemented")
	//card, err := s.paymentSvc.AddCustomerCard(customer)
	//if err != nil {
	//	return nil, err
	//}
	//
	//card.Id, err = s.repo.AddCustomerCard(customer.Id, card)
	//if err != nil {
	//	_ = s.paymentSvc.DeleteCustomerCard(customer, card)
	//	return nil, err
	//}
	//
	//return card, nil
}
