package server

import (
	"context"
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

	log.Println("AddCustomerPaymentMethod", c)

	resp := &pb.AddCustomerPaymentMethodResponse{
		Success: true,
	}

	return resp, nil
}
