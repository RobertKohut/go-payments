package customers

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/robertkohut/go-payments/pkg/entities"
	"github.com/robertkohut/go-payments/pkg/metadata"
	"log"
)

type Repository interface {
	InsertCustomer(customer *entities.Customer) (int64, error)
	SelectCustomerByAccountId(sourceId, accountId int64) (*entities.Customer, error)
	DeleteCustomerBySourceId(customer *entities.Customer) error

	AddCustomerCard(customer *entities.Customer, card *entities.Card) (int64, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) InsertCustomer(customer *entities.Customer) (int64, error) {
	// TODO: Remove magic gateway_id
	stmt := `INSERT INTO customers (gateway_id, source_id, account_id, ext_id, flags) VALUES (1, ?, ?, ?, ?)`

	customer.Flags = customer.Flags | metadata.FlagsCustomerActive

	result, err := r.db.Exec(stmt, customer.SourceId, customer.AccountId, customer.ExtId, customer.Flags)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.LastInsertId()
}

func (r *repository) DeleteCustomerBySourceId(customer *entities.Customer) error {
	stmt := `UPDATE customers SET flags = flags &~ ? 
                 WHERE account_id = ?
                 AND (flags & ?) = ?`

	_, err := r.db.Exec(stmt, metadata.FlagsCustomerActive, customer.AccountId, metadata.FlagsCustomerActive, metadata.FlagsCustomerActive)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *repository) SelectCustomerByAccountId(sourceId, accountId int64) (*entities.Customer, error) {
	customer := &entities.Customer{}

	stmt := `SELECT id, ext_id FROM customers 
             WHERE source_id = ?
               AND account_id = ?
               AND (flags & ?) = ?`

	row := r.db.QueryRow(stmt, sourceId, accountId, metadata.FlagsCustomerActive, metadata.FlagsCustomerActive)

	switch err := row.Scan(
		&customer.Id,
		&customer.ExtId,
	); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return customer, nil
	default:
		return nil, err
	}
}

func (r *repository) AddCustomerCard(customer *entities.Customer, card *entities.Card) (int64, error) {
	stmt := `INSERT INTO cards (ext_id, customer_id, exp_month, exp_year, last_four)
			 VALUES (?, ?, ?, ?, ?)`

	result, err := r.db.Exec(stmt,
		card.ExtId,
		customer.Id,
		card.ExpMonth,
		card.ExpYear,
		card.LastFour,
	)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.LastInsertId()
}
