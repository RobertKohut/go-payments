package customers

import (
	"github.com/robertkohut/go-payments/pkg/payments"
	pb "github.com/robertkohut/go-payments/proto"
	"log"
)

type Service interface {
	AddCustomer(customer *pb.Customer) (*string, error)
	GetCustomerById(sourceId, accountId int64) (*pb.Customer, error)
	DeleteCustomer(customer *pb.Customer) error

	AddCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) (*pb.Card, error)
	GetCustomerPaymentMethod(customer *pb.Customer, cardId int64) (*pb.Card, error)
	SetCustomerPrimaryPaymentMethod(customer *pb.Customer, card *pb.Card) error
	RemoveCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) error
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

	customer.Cards, err = s.repo.SelectCustomerCards(customer)

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

	cardId, err := s.repo.AddCustomerCard(customer, card)
	if err != nil {
		return nil, err
	}

	card.Id = cardId

	return card, nil
}

func (s *service) GetCustomerPaymentMethod(customer *pb.Customer, cardId int64) (*pb.Card, error) {
	log.Println("GetCustomerCard", cardId)

	card, err := s.repo.SelectCustomerCard(customer, cardId)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (s *service) SetCustomerPrimaryPaymentMethod(customer *pb.Customer, card *pb.Card) error {
	log.Println("SetCustomerPrimaryPaymentMethod", card)

	err := s.repo.UpdateCustomerPrimaryCard(customer, card)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) RemoveCustomerPaymentMethod(customer *pb.Customer, card *pb.Card) error {
	log.Println("RemoveCustomerCard", card)

	err := s.paymentSvc.RemoveCustomerPaymentMethod(customer, card)
	if err != nil {
		return err
	}

	err = s.repo.DeleteCustomerCard(customer, card)
	if err != nil {
		return err
	}

	return nil
}
