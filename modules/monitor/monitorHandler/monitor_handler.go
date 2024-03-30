package monitorHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rapinbook/ecommerce-go/config"
	"github.com/rapinbook/ecommerce-go/modules/monitor"
)

type IMonitor interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	cfg config.IConfig
}

func NewMonitorHandler(cfg config.IConfig) IMonitor {
	return &monitorHandler{
		cfg: cfg,
	}

}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.cfg.App().AppName(),
		Version: h.cfg.App().Version(),
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
