package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/one-d-plate/go-boilerplate.git/src/config"
	"github.com/one-d-plate/go-boilerplate.git/src/pkg"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type database struct {
	conf config.Config
}

func NewDatabase() *database {
	dsn := config.Config{}
	return &database{
		conf: dsn,
	}
}

func (d *database) Connect() (*bun.DB, error) {
	var (
		driver string
		dsn    string
	)

	dsn = d.conf.GetDSN()

	switch d.conf.DbConfig.Driver {
	case "mysql":
		driver = "mysql"
	case "postgres":
		driver = "pgx"
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", d.conf.DbConfig.Driver)
	}

	sqldb, err := sql.Open(driver, dsn)
	if err != nil {
		pkg.Logger.Error("Connection failed ", err)
		return nil, err
	}

	if err := sqldb.Ping(); err != nil {
		pkg.Logger.Error("Database ping failed ", err)
		return nil, err
	}

	var db *bun.DB
	switch d.conf.DbConfig.Driver {
	case "mysql":
		db = bun.NewDB(sqldb, mysqldialect.New())
	case "postgres":
		db = bun.NewDB(sqldb, pgdialect.New())
	}

	return db, nil
}
