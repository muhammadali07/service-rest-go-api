package main

import (
	"github.com/muhammadali07/service-rest-go-api/pkg/log"
	"github.com/muhammadali07/service-rest-go-api/pkg/utils"
	"github.com/muhammadali07/service-rest-go-api/services/api"
	"github.com/muhammadali07/service-rest-go-api/services/app"
	"github.com/muhammadali07/service-rest-go-api/services/datastore"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := log.NewLogger("middleware-rest-api")

	// Inisialisasi konfigurasi
	cfg, err := utils.InitConfig()
	if err != nil {
		logger.Fatal(
			logrus.Fields{
				"err": err,
			}, "nil", err.Error())
	}

	ds := datastore.InitMiddleDatastore(
		cfg.Driver,
		cfg.DatabaseHost,
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.Database,
		cfg.DatabasePort, logger)
	app := app.InitMiddleApplication(ds, logger)
	api := api.InitMiddledAPI(logger, app)
	api.Start()
}
