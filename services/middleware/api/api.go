package api

import (
	"net"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadali07/service-rest-go-api/pkg/log"
	"github.com/muhammadali07/service-rest-go-api/services/middleware/app"
	"github.com/sirupsen/logrus"
)

type MiddlewareAPI struct {
	app       app.MiddlewareServicePort
	log       *log.Logger
	address   string
	validator *validator.Validate
}

func (m *MiddlewareAPI) Start() {
	listener, err := net.Listen("tcp", m.address)
	if err != nil {
		m.log.Fatal(logrus.Fields{
			"error":   err.Error(),
			"address": m.address,
		}, "nil", "failed to listen address")
	}

	app := fiber.New()

	if err := app.Listener(listener); err != nil {
		m.log.Fatal(logrus.Fields{
			"error": err.Error(),
		}, "nil", "failed to server grpc API")
	}
}

func InitMiddledAPI(log *log.Logger, app app.MiddlewareServicePort) *MiddlewareAPI {
	return &MiddlewareAPI{
		log:       log,
		app:       app,
		validator: validator.New(),
	}
}
