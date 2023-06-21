package server

import (
	"context"
	"fmt"
	pb "github.com/robertkohut/go-payments/proto"
	"log"
)

func (s *Server) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	customer := &pb.Customer{
		SourceId:  req.GetSourceId(),
		AccountId: req.GetAccountId(),
		Name:      req.GetName(),
	}

	extId, err := s.svc.CustomerSvc.AddCustomer(customer)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateCustomerResponse{
		Customer: &pb.Customer{
			ExtId: *extId,
		},
	}

	return resp, nil
}

func (s *Server) GetCustomerById(ctx context.Context, req *pb.GetCustomerByIdRequest) (*pb.GetCustomerByIdResponse, error) {
	sourceId := req.GetSourceId()
	accountId := req.GetAccountId()

	c, err := s.svc.CustomerSvc.GetCustomerById(sourceId, accountId)
	if err != nil {
		return nil, err
	}

	if c == nil {
		resp := &pb.GetCustomerByIdResponse{
			Customer: nil,
		}

		return resp, nil
	}

	resp := &pb.GetCustomerByIdResponse{
		Customer: c,
	}

	return resp, nil
}

func (s *Server) AddCustomerPaymentMethod(ctx context.Context, req *pb.AddCustomerPaymentMethodRequest) (*pb.AddCustomerPaymentMethodResponse, error) {
	sourceId := req.GetSourceId()
	accountId := req.GetAccountId()
	card := req.GetCard()

	customer, err := s.svc.CustomerSvc.GetCustomerById(sourceId, accountId)
	if err != nil {
		return nil, err
	}

	log.Println("AddCustomerPaymentMethod", sourceId, accountId, card)

	c, err := s.svc.CustomerSvc.AddCustomerPaymentMethod(customer, card)
	if err != nil {
		return nil, err
	}

	if len(customer.Cards) == 0 {
		err = s.svc.CustomerSvc.SetCustomerPrimaryPaymentMethod(customer, c)
		if err != nil {
			log.Println("Customer -> AddCustomerPaymentMethod():", err)
		}
	}

	log.Println("AddCustomerPaymentMethod", c)

	resp := &pb.AddCustomerPaymentMethodResponse{
		Success: true,
		Card:    c,
	}

	return resp, nil
}

func (s *Server) RemoveCustomerPaymentMethod(ctx context.Context, req *pb.RemoveCustomerPaymentMethodRequest) (*pb.RemoveCustomerPaymentMethodResponse, error) {
	sourceId := req.GetSourceId()
	accountId := req.GetAccountId()
	cardId := req.GetCardId()

	customer, err := s.svc.CustomerSvc.GetCustomerById(sourceId, accountId)
	if err != nil {
		return nil, err
	}

	card, err := s.svc.CustomerSvc.GetCustomerPaymentMethod(customer, cardId)
	if err != nil {
		return nil, err
	}

	err = s.svc.CustomerSvc.RemoveCustomerPaymentMethod(customer, card)
	if err != nil {
		return nil, err
	}

	resp := &pb.RemoveCustomerPaymentMethodResponse{
		Success: true,
	}

	return resp, nil
}

func (s *Server) SetCustomerPrimaryPaymentMethod(ctx context.Context, req *pb.SetCustomerPrimaryPaymentMethodRequest) (*pb.SetCustomerPrimaryPaymentMethodResponse, error) {
	sourceId := req.GetSourceId()
	accountId := req.GetAccountId()
	cardId := req.GetCardId()

	customer, err := s.svc.CustomerSvc.GetCustomerById(sourceId, accountId)
	if err != nil {
		return nil, err
	}

	card, err := s.svc.CustomerSvc.GetCustomerPaymentMethod(customer, cardId)
	if err != nil {
		return nil, err
	}

	err = s.svc.CustomerSvc.SetCustomerPrimaryPaymentMethod(customer, card)
	if err != nil {
		return nil, err
	}

	resp := &pb.SetCustomerPrimaryPaymentMethodResponse{
		Success: true,
	}

	return resp, nil
}

func (s *Server) RetrieveCustomerCharges(ctx context.Context, req *pb.RetrieveCustomerChargesRequest) (*pb.RetrieveCustomerChargesResponse, error) {
	sourceId := req.GetSourceId()
	accountId := req.GetAccountId()
	filters := req.GetFilters()

	customer, err := s.svc.CustomerSvc.GetCustomerById(sourceId, accountId)
	if err != nil {
		return nil, err
	}

	charges, err := s.svc.ChargeSvc.GetCustomerCharges(customer, filters)
	if err != nil {
		return nil, err
	}

	resp := &pb.RetrieveCustomerChargesResponse{
		Charges: charges,
	}

	return resp, nil
}

func (s *Server) CreateCharge(ctx context.Context, req *pb.CreateChargeRequest) (*pb.CreateChargeResponse, error) {
	const (
		errInvalidInput      = "invalid input"
		errSourceIDRequired  = "source id is required"
		errAccountIDRequired = "account id is required"
		errCurrencyRequired  = "currency is required"
		errNoPrimaryCardSet  = "no primary card set"
	)

	if req.GetSourceId() == 0 {
		return nil, fmt.Errorf("%w: %s", errInvalidInput, errSourceIDRequired)
	}

	if req.GetAccountId() == 0 {
		return nil, fmt.Errorf("%w: %s", errInvalidInput, errAccountIDRequired)
	}

	if req.GetCharge().Currency == "" {
		return nil, fmt.Errorf("%w: %s", errInvalidInput, errCurrencyRequired)
	}

	charge := req.GetCharge()
	cardId := charge.GetPmId()

	customer, err := s.svc.CustomerSvc.GetCustomerById(req.GetSourceId(), req.GetAccountId())
	if err != nil {
		return nil, err
	}

	// A default card was not set. Use the primary card
	if cardId == 0 {
		cardId = customer.GetPrimaryCardId()

		// No primary card was set. Use the first card
		if cardId == 0 {
			return nil, fmt.Errorf("%w: %s", errInvalidInput, errNoPrimaryCardSet)
		}
	}

	card, err := s.svc.CustomerSvc.GetCustomerPaymentMethod(customer, cardId)
	if err != nil {
		return nil, err
	}

	charge, err = s.svc.ChargeSvc.ChargeCustomerPaymentMethod(customer, card, charge)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateChargeResponse{
		Charge: charge,
	}

	return resp, nil
}
