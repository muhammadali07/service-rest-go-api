package main

import (
	"github.com/muhammadali07/service-rest-go-api/pkg/log"
	"github.com/muhammadali07/service-rest-go-api/pkg/utils"
	"github.com/muhammadali07/service-rest-go-api/services/middleware/api"
	"github.com/muhammadali07/service-rest-go-api/services/middleware/app"
	"github.com/muhammadali07/service-rest-go-api/services/middleware/datastore"
	"github.com/sirupsen/logrus"
)

func main() {
	// Inisialisasi konfigurasi
	cfg, err := utils.InitConfig()

	logger := log.NewLogger(cfg.Service)

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
