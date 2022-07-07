package common

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"notes/infra"
	"os"
)

var DbConn *sql.DB

func runMigrations(DbDriver string) {
	err := goose.SetDialect(DbDriver)

	workdir, _ := os.Getwd()

	err = goose.Up(DbConn, fmt.Sprintf("%s/sql", workdir))

	if err != nil {
		panic(err)
	}
}

func InitDatabase(cfg *infra.Config) {
	var err error
	DbConn, err = sql.Open(cfg.DbDriver, cfg.DbUrl)

	if err != nil {
		panic(err)
	}

	runMigrations(infra.Cfg.DbDriver)

	fmt.Println("Connected to the repositories")
}
