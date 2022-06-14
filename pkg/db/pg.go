package db

import (
	"bitbucket.org/ssinbeti/crm-common/pkg/logger"
	"database/sql"
	"github.com/cenkalti/backoff"
	"log"
	"time"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

const timeoutDuration = "2s"

func InitDBClient(psn string, logger logger.Logger) (db *sql.DB) {
	timeout, err := time.ParseDuration(timeoutDuration)
	checkErr(err)

	operation := func() error {
		db, err = sql.Open("nrpostgres", psn)
		if err != nil {
			logger.Warn(err)
			return err
		}
		if err := db.Ping(); err != nil {
			logger.Warn(err)
			return err
		}
		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = timeout

	err = backoff.Retry(operation, b)
	checkErr(err)

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
