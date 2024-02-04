package tenants

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/robertkohut/go-payments/internal/services/hashid"
	pb "github.com/robertkohut/go-payments/proto"
	"google.golang.org/grpc/status"
)

type Repository interface {
	SelectTenantByApiKey(source, key, gateway string) (*pb.Tenant, error)
}

type repository struct {
	db *sqlx.DB
	hd *hashid.Service
}

func NewRepository(db *sqlx.DB, hd *hashid.Service) Repository {
	return &repository{
		db: db,
		hd: hd,
	}
}

func (r *repository) SelectTenantByApiKey(source, key, gateway string) (*pb.Tenant, error) {
	t := &pb.Tenant{}

	stmt := `SELECT t.id, t.name, g.ext_id
			 FROM tenants t
			 INNER JOIN sources s ON s.tenant_id = t.id
			 INNER JOIN api_keys a ON a.source_id = s.id
			 INNER JOIN gateways g ON g.tenant_id = t.id
			 WHERE s.name = ?
			  AND a.hash = ?
			  AND g.gateway = ?`

	row := r.db.QueryRowx(stmt, source, key, gateway)

	switch err := row.Scan(
		&t.Id,
		&t.LegalName,
		&t.ExternalId,
	); {
	case errors.Is(err, sql.ErrNoRows):
		st := status.New(404, "tenant not found")
		return nil, st.Err()
	case err == nil:
		return t, nil
	default:
		return nil, err
	}
}
