package server

import (
	"context"
	pb "github.com/robertkohut/go-payments/proto"
)

func (s *Server) GetPublishableKey(ctx context.Context, req *pb.GetPublishableKeyRequest) (*pb.GetPublishableKeyResponse, error) {
	key, err := s.svc.CustomerSvc.GetPublishableKey()
	if err != nil {
		return nil, err
	}

	resp := &pb.GetPublishableKeyResponse{
		PublishableKey: key,
	}

	return resp, nil
}
