package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/robertkohut/go-payments/internal/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect(cfg *config.DBConfig) (*sqlx.DB, error) {
	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := sqlx.Open("mysql", connection)

	if err != nil {
		log.Println(err)
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
		time.Sleep(time.Duration(5) * time.Second)
		return DBConnect(cfg)
	}

	return db, nil
}
