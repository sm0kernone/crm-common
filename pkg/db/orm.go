package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"time"

	// DB adapter
	_ "github.com/lib/pq"
)

type DB interface {
	orm.DB
	//RunInTransaction(fn func(*pg.Tx) error) error
}

// New creates new database connection to a postgres database
// Function panics if it can't connect to database
func NewGoPG(psn string, logger pg.QueryHook) (*pg.DB, error) {
	u, err := pg.ParseURL(psn)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(u).WithTimeout(time.Second * 50)

	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	if logger != nil {
		db.AddQueryHook(logger)
	}

	return db, nil
}
