package server

import (
	"context"
	"github.com/robertkohut/go-payments/pkg/entities"
	payments "github.com/robertkohut/go-payments/proto"
)

func (s *Server) CreateCustomer(ctx context.Context, req *payments.CreateCustomerRequest) (*payments.CreateCustomerResponse, error) {
	// Implement the logic for creating a customer here.

	customer := &entities.Customer{
		OrgId:     req.GetOrgId(),
		AccountId: req.GetAccountId(),
		Name:      req.GetName(),
	}

	extId, err := s.svc.CustomerSvc.AddCustomer(customer)
	if err != nil {
		return nil, err
	}

	resp := &payments.CreateCustomerResponse{
		Customer: &payments.Customer{
			ExtId: *extId,
		},
	}

	return resp, nil
}

func (s *Server) GetCustomerById(ctx context.Context, req *payments.GetCustomerByIdRequest) (*payments.GetCustomerByIdResponse, error) {
	sourceId := req.GetSourceId()
	accountId := req.GetAccountId()

	c, err := s.svc.CustomerSvc.GetCustomerById(sourceId, accountId)
	if err != nil {
		return nil, err
	}

	if c == nil {
		resp := &payments.GetCustomerByIdResponse{
			Customer: nil,
		}

		return resp, nil
	}

	resp := &payments.GetCustomerByIdResponse{
		Customer: &payments.Customer{
			ExtId: c.ExtId,
		},
	}

	return resp, nil
}
