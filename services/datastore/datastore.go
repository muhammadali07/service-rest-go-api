package datastore

import (
	"github.com/jmoiron/sqlx"
	"github.com/muhammadali07/service-rest-go-api/pkg/log"
)

type MiddlewareDataStore struct {
	db  *sqlx.DB
	log *log.Logger
}
