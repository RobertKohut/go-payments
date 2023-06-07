package customers

import (
	"errors"
	"github.com/robertkohut/go-payments/pkg/payments"
	pb "github.com/robertkohut/go-payments/proto"
	"log"
)

type Service interface {
	AddCustomer(customer *pb.Customer) (*string, error)
	AddCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) (*pb.Card, error)

	GetCustomerById(sourceId, accountId int64) (*pb.Customer, error)

	DeleteCustomer(customer *pb.Customer) error
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

func (s *service) GetCustomerById(sourceId, accountId int64) (*pb.Customer, error) {
	customer, err := s.repo.SelectCustomerByAccountId(sourceId, accountId)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *service) AddCustomer(customer *pb.Customer) (*string, error) {
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

func (s *service) DeleteCustomer(customer *pb.Customer) error {
	err := s.repo.DeleteCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) AddCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) (*pb.Card, error) {
	log.Println("AddCustomerCard", card)

	_, err := s.paymentSvc.AddCustomerPaymentMethod(customer, card)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.AddCustomerCard(customer, card)
	if err != nil {
		return nil, err
	}

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
