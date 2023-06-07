package customers

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/robertkohut/go-payments/pkg/metadata"
	pb "github.com/robertkohut/go-payments/proto"
	"log"
)

type Repository interface {
	InsertCustomer(customer *pb.Customer) (int64, error)
	SelectCustomerByAccountId(sourceId, accountId int64) (*pb.Customer, error)
	DeleteCustomer(customer *pb.Customer) error

	AddCustomerCard(customer *pb.Customer, card *pb.Card) (int64, error)
	SelectCustomerCards(customer *pb.Customer) ([]*pb.Card, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) InsertCustomer(customer *pb.Customer) (int64, error) {
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

func (r *repository) DeleteCustomer(customer *pb.Customer) error {
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

func (r *repository) SelectCustomerByAccountId(sourceId, accountId int64) (*pb.Customer, error) {
	customer := &pb.Customer{}

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

func (r *repository) AddCustomerCard(customer *pb.Customer, card *pb.Card) (int64, error) {
	stmt := `INSERT INTO cards (ext_id, customer_id, exp_month, exp_year, last_four)
			 VALUES (?, ?, ?, ?, ?)`

	knownCardBrands := map[string]bool{
		"visa":       true,
		"mastercard": true,
		"amex":       true,
		"discover":   true,
		"jcb":        true,
		"diners":     true,
	}

	if _, ok := knownCardBrands[card.Brand]; !ok {
		log.Println("Unknown card brand: ", card.Brand, " for card: ", card.ExtId)
		card.Brand = "unknown"
	}

	result, err := r.db.Exec(stmt,
		card.ExtId,
		customer.Id,
		card.ExpMonth,
		card.ExpYear,
		card.Last4,
	)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.LastInsertId()
}

func (r *repository) SelectCustomerCards(customer *pb.Customer) ([]*pb.Card, error) {
	var cards []*pb.Card

	stmt := `SELECT id, brand, ext_id, exp_month, exp_year, last_four FROM cards 
			 WHERE customer_id = ?
			   AND (flags & ?) = ?`

	rows, err := r.db.Query(stmt, customer.Id, metadata.FlagsCardActive, metadata.FlagsCardActive)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		card := &pb.Card{}

		err := rows.Scan(
			&card.Id,
			&card.Brand,
			&card.ExtId,
			&card.ExpMonth,
			&card.ExpYear,
			&card.Last4,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}
