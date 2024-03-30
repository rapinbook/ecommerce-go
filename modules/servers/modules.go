package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rapinbook/ecommerce-go/modules/monitor/monitorHandler"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r fiber.Router
	s *server
}

func NewModule(r fiber.Router, s *server) IModuleFactory {
	return &moduleFactory{
		r: r,
		s: s,
	}
}

// routers for code cleaniness
func (m *moduleFactory) MonitorModule() {
	handler := monitorHandler.NewMonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}
