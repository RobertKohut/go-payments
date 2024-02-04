package charges

import (
	"context"
	"github.com/robertkohut/go-payments/internal/services/hashid"
	"github.com/robertkohut/go-payments/pkg/metadata"
	"github.com/robertkohut/go-payments/pkg/payments"
	pb "github.com/robertkohut/go-payments/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

type Service interface {
	ChargeCustomerPaymentMethod(ctx context.Context, customer *pb.Customer, card *pb.Card, charge *pb.Charge) (*pb.Charge, error)
	GetCustomerCharges(customer *pb.Customer, filter *pb.Filters) ([]*pb.Charge, error)
	ChargeOneTimePayment(ctx context.Context, card *pb.Card, charge *pb.Charge) (*pb.Charge, error)
}

type service struct {
	paymentSvc payments.PaymentService
	repo       Repository
	hd         *hashid.Service
}

func NewService(payments payments.PaymentService, repo Repository, hd *hashid.Service) Service {
	return &service{
		paymentSvc: payments,
		repo:       repo,
		hd:         hd,
	}
}

func (s *service) getCurrencyIdByCode(code string) int64 {
	id, err := s.repo.SelectCurrencyIdByCode(code)
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (s *service) ChargeOneTimePayment(ctx context.Context, card *pb.Card, charge *pb.Charge) (*pb.Charge, error) {

	extId, err := s.paymentSvc.CreateCharge(ctx, nil, card, charge)
	if err != nil {
		charge.Status = "failed"
		_ = s.repo.UpdateCharge(charge)
		return nil, err
	}

	charge.ExtId = *extId

	return charge, nil
}

func (s *service) ChargeCustomerPaymentMethod(ctx context.Context, customer *pb.Customer, card *pb.Card, charge *pb.Charge) (*pb.Charge, error) {
	charge.GatewayId = customer.GatewayId
	charge.CustomerId = customer.Id
	charge.CurrencyId = s.getCurrencyIdByCode(charge.Currency)

	chargeId, err := s.repo.InsertCharge(charge)
	if err != nil {
		return nil, err
	}

	hdInvoiceId, _ := s.hd.Encode([]int64{chargeId, metadata.HDChargeId})
	if charge.Description == "" {
		charge.Description = "Invoice " + hdInvoiceId
	}

	extId, err := s.paymentSvc.CreateCharge(ctx, customer, card, charge)
	if err != nil {
		charge.Status = "failed"
		_ = s.repo.UpdateCharge(charge)
		return nil, err
	}

	charge.ExtId = *extId

	err = s.repo.UpdateCharge(charge)
	if err != nil {
		return nil, err
	}

	charge.Id = chargeId

	return charge, nil
}

func (s *service) GetCustomerCharges(customer *pb.Customer, filter *pb.Filters) ([]*pb.Charge, error) {
	if filter == nil {
		filter = &pb.Filters{}
	}

	filter.Filters = append(filter.GetFilters(), &pb.Filter{
		Column:   "customer_id",
		Operator: "=",
		Value:    structpb.NewNumberValue(float64(customer.Id)),
	})

	return s.repo.SelectCharges(filter)
}
