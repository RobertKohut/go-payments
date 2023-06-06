package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/robertkohut/go-payments/pkg/customers"
)

type Services struct {
	DB          *sqlx.DB
	CustomerSvc customers.Service
}
