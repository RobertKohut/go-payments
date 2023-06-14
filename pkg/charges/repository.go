package charges

import (
	"errors"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jmoiron/sqlx"
	"github.com/robertkohut/go-payments/internal/services/hashid"
	"github.com/robertkohut/go-payments/pkg/metadata"
	pb "github.com/robertkohut/go-payments/proto"
	"time"
)

type Repository interface {
	SelectCharges(filter *pb.Filters) ([]*pb.Charge, error)
	InsertCharge(charge *pb.Charge) (int64, error)
	UpdateCharge(charge *pb.Charge) error

	SelectCurrencyIdByCode(code string) (int64, error)
}

type repository struct {
	db *sqlx.DB
	hd *hashid.Service
}

func NewRepository(db *sqlx.DB, hd *hashid.Service) Repository {
	return &repository{db: db, hd: hd}
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

func (r *repository) SelectCharges(filter *pb.Filters) ([]*pb.Charge, error) {
	var charges []*pb.Charge
	var args []interface{}

	if filter == nil {
		return nil, errors.New("missing filter")
	}

	stmt := `SELECT c.id,
                    c.ext_id,
                    c.customer_id,
                    c.description,
                    c.pm_type,
                    c.pm_id,
                    c.amount,
                    currencies.code AS currency,
                    c.status,
                    c.created_at,
                    c.updated_at
            FROM charges c
            INNER JOIN currencies ON currencies.id = c.currency_id`

	query, filterArgs := buildFilterQuery(filter)

	stmt += ` WHERE ` + query
	stmt += ` ORDER BY created_at DESC `

	args = append(args, filterArgs...)

	if filter.GetLimit() > 0 {
		stmt += ` LIMIT ? OFFSET ?`
		args = append(args, filter.GetLimit(), filter.GetOffset())
	}

	rows, err := r.db.Queryx(stmt, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	charges, err = r.scanCharges(rows)
	if err != nil {
		return nil, err
	}

	return charges, nil
}

func (r *repository) scanCharges(rows *sqlx.Rows) ([]*pb.Charge, error) {
	var charges []*pb.Charge

	for rows.Next() {
		var charge pb.Charge
		var createdAt, updatedAt time.Time

		err := rows.Scan(
			&charge.Id,
			&charge.ExtId,
			&charge.CustomerId,
			&charge.Description,
			&charge.PmType,
			&charge.PmId,
			&charge.Amount,
			&charge.Currency,
			&charge.Status,
			&createdAt,
			&updatedAt,
		)

		charge.CreatedAt = timeToTimestamp(createdAt)
		charge.UpdatedAt = timeToTimestamp(updatedAt)

		if err != nil {
			return nil, err
		}

		charge.IdStr, err = r.hd.Encode([]int64{charge.Id, metadata.HDChargeId})
		if err != nil {
			return nil, err
		}

		charges = append(charges, &charge)
	}

	return charges, nil
}

// Convert a time.Time to a google.protobuf.Timestamp
func timeToTimestamp(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.UnixNano() % 1e9),
	}
}

// Convert a google.protobuf.Timestamp to a time.Time
func timestampToTime(ts *timestamp.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
}

func buildFilterQuery(filter *pb.Filters) (string, []interface{}) {
	var query string
	var args []interface{}

	for i, f := range filter.Filters {
		// Add the column name to the query
		query += f.Column

		// Add the operator to the query
		query += " " + f.Operator + " "

		// Add the value to the query
		switch v := f.Value.GetKind().(type) {
		case *structpb.Value_NullValue:
			query += "NULL"
		case *structpb.Value_NumberValue:
			query += "?"
			args = append(args, v.NumberValue)
		case *structpb.Value_StringValue:
			query += "?"
			args = append(args, v.StringValue)
		case *structpb.Value_BoolValue:
			query += "?"
			args = append(args, v.BoolValue)
		}

		// Add the AND keyword to the query if there are more filters
		if i < len(filter.Filters)-1 {
			query += " AND "
		}
	}

	return query, args
}
