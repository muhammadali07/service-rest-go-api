package app

import "github.com/muhammadali07/service-rest-go-api/pkg/log"

type MiddlewareApplication struct {
	datastore MiddlewareDataStorePort
	log       *log.Logger
}

func InitMiddleApplication(datastore MiddlewareDataStorePort, log *log.Logger) *MiddlewareApplication {
	return &MiddlewareApplication{
		datastore: datastore,
		log:       log,
	}
}
