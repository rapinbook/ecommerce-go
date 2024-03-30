package main

import (
	"fmt"
	"os"

	"github.com/rapinbook/ecommerce-go/config"
	"github.com/rapinbook/ecommerce-go/modules/servers"
	"github.com/rapinbook/ecommerce-go/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	fmt.Println(cfg.Db())
	fmt.Println(cfg.JWT())
	fmt.Println(cfg.App())
	db := databases.DBConnect(cfg.Db())
	fmt.Printf("%#v", db)
	s := servers.NewServer(cfg, db)
	s.Start()
}
