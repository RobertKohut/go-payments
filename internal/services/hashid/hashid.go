package hashid

import (
	"errors"
	"github.com/robertkohut/go-payments/internal/config"
	"github.com/speps/go-hashids/v2"
)

type Service struct {
	hd *hashids.HashID
}

func New(cfg *config.HashIdConfig) (*Service, error) {
	hd := hashids.NewData()
	hd.Salt = cfg.Salt
	hd.MinLength = cfg.MinLength
	hd.Alphabet = cfg.Alphabet
	hashID, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}
	return &Service{hd: hashID}, nil
}

func (s *Service) Encode(ids []int64) (string, error) {
	hash, err := s.hd.EncodeInt64(ids)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (s *Service) Decode(hash string) ([]int64, error) {
	ids, err := s.hd.DecodeInt64WithError(hash)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("no valid ids found")
	}
	return ids, nil
}
