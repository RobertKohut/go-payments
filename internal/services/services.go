package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/robertkohut/go-payments/internal/services/hashid"
	"github.com/robertkohut/go-payments/pkg/charges"
	"github.com/robertkohut/go-payments/pkg/customers"
)

type Services struct {
	DB          *sqlx.DB
	HashId      *hashid.Service
	CustomerSvc customers.Service
	ChargeSvc   charges.Service
}
