package api

import (
	"fmt"
	"net"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadali07/service-rest-go-api/pkg/log"
	"github.com/muhammadali07/service-rest-go-api/services/middleware/app"
	"github.com/sirupsen/logrus"
)

type MiddlewareAPI struct {
	host      string
	port      string
	app       app.MiddlewareServicePort
	log       *log.Logger
	validator *validator.Validate
}

func (m *MiddlewareAPI) Start() {
	address := fmt.Sprintf("%v:%v", m.host, m.port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		m.log.Fatal(logrus.Fields{
			"error":   err.Error(),
			"address": address,
		}, "nil", "failed to listen address")
	}

	api := fiber.New()
	if err := api.Listener(listener); err != nil {
		m.log.Fatal(logrus.Fields{
			"error": err.Error(),
		}, "nil", "failed to serve Middleware API")
	}
}

func InitMiddledAPI(host string, port string, log *log.Logger, appInstance app.MiddlewareServicePort) *MiddlewareAPI {
	return &MiddlewareAPI{
		host:      host,
		port:      port,
		log:       log,
		app:       appInstance,
		validator: validator.New(),
	}
}
