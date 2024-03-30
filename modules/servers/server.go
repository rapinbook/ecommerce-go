package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rapinbook/ecommerce-go/config"
)

type IServer interface {
	Start()
}

type server struct {
	cfg config.IConfig
	app *fiber.App
	db  *sqlx.DB
}

func (s *server) Start() {
	// v1 := s.app.Group("v1")

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	s.app.Listen(s.cfg.App().Url())
}

func NewServer(cfg config.IConfig, db *sqlx.DB) IServer {
	return &server{
		cfg: cfg,
		app: fiber.New(fiber.Config{
			AppName:      cfg.App().AppName(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeout(),
			WriteTimeout: cfg.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
		db: db,
	}
}
