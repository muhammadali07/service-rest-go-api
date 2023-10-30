package datastore

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/muhammadali07/service-rest-go-api/pkg/log"
)

type MiddlewareDatabase struct {
	db  *sqlx.DB
	log *log.Logger
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
		address = fmt.Sprintf("%s//%s:%s@%s:%d/%s", driver, user, password, host, port, database)
	} else if driver == "godror" {
		address = fmt.Sprintf(`user="%s" password="%s" connectString="(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=%d))(CONNECT_DATA=(SID=%s)))"`, user, password, host, port, database)
	}
	db, err := sqlx.Connect(driver, address)
	if err != nil {
		panic(err)
	}
	return &MiddlewareDatabase{
		db:  db,
		log: log,
	}
}
