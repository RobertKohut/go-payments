package charges

import (
	"github.com/jmoiron/sqlx"
	pb "github.com/robertkohut/go-payments/proto"
)

type Repository interface {
	InsertCharge(charge *pb.Charge) (int64, error)
	UpdateCharge(charge *pb.Charge) error

	SelectCurrencyIdByCode(code string) (int64, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) InsertCharge(charge *pb.Charge) (int64, error) {
	stmt := `INSERT INTO charges (
                     gateway_id,
                     ext_id,
                     customer_id,
                     pm_type, 
                     pm_id, 
                     amount, 
                     currency_id, 
                     status)
    		 VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(
		stmt,
		charge.GetGatewayId(),
		charge.GetExtId(),
		charge.GetCustomerId(),
		charge.GetPmType(),
		charge.GetPmId(),
		charge.GetAmount(),
		charge.GetCurrencyId(),
		charge.GetStatus(),
	)
	if err != nil {
		return 0, err
	}

	charge.Id, _ = result.LastInsertId()

	return result.LastInsertId()
}

func (r *repository) UpdateCharge(charge *pb.Charge) error {
	stmt := `UPDATE charges
			 SET status = ?,
			     ext_id = ?,
			     updated_at = CURRENT_TIMESTAMP
			 WHERE id = ?`

	_, err := r.db.Exec(
		stmt,
		charge.GetStatus(),
		charge.GetExtId(),
		charge.GetId(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) SelectCurrencyIdByCode(code string) (int64, error) {
	stmt := `SELECT id FROM currencies WHERE code = ?`

	var id int64
	err := r.db.Get(&id, stmt, code)
	if err != nil {
		return 0, err
	}

	return id, nil
}
