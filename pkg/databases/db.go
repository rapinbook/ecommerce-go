package databases

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rapinbook/ecommerce-go/config"
)

func DBConnect(c config.IDbConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", c.Url())
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
	return db
}
