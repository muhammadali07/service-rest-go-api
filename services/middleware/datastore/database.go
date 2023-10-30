package datastore

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/muhammadali07/service-rest-go-api/pkg/log"
	"github.com/sirupsen/logrus"
)

type MiddlewareDatabase struct {
	db  *sqlx.DB
	log *log.Logger
}

func (m *MiddlewareDatabase) Begin() (tx *sqlx.Tx, err error) {
	tx, err = m.db.Beginx()
	if err != nil {
		remark := "failed to start transaction"
		m.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
	}
	return
}

func (f *MiddlewareDatabase) Rollback(tx *sqlx.Tx) {
	err := tx.Rollback()
	if err != nil {
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, "failed to rollback transaction")
	}
}

func (f *MiddlewareDatabase) Commit(tx *sqlx.Tx) {
	err := tx.Commit()
	if err != nil {
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, "failed to rollback transaction")
	}
}

func InitMiddleDatastore(
	driver,
	host,
	user,
	password,
	database string,
	port int,
	log *log.Logger) *MiddlewareDatabase {
	var address string

	if driver == "postgres" {
		address = fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable", user, password, host, port, database)
	} else if driver == "godror" {
		address = fmt.Sprintf(`user="%s" password="%s" connectString="(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=%d))(CONNECT_DATA=(SID=%s)))"`, user, password, host, port, database)
	} else if driver == "postgresql" {
		address = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, database)
	}

	// if using ssl database
	if driver == "postgresql" {
		db, err := sqlx.Open(driver, address)
		if err != nil {
			panic(err)
		}

		return &MiddlewareDatabase{
			db:  db,
			log: log,
		}

	} else {
		db, err := sqlx.Connect(driver, address)
		if err != nil {
			panic(err)
		}

		return &MiddlewareDatabase{
			db:  db,
			log: log,
		}
	}
}
