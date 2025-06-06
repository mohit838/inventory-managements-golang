package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mohit838/inventory-managements-golang/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

func DbInitialized(cfg config.Database) (*sqlx.DB, error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.Charset,
		cfg.ParseTime,
		cfg.Loc,
	)

	db, err := sqlx.Connect(cfg.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database (driver=%q, dsn=%q): %w", cfg.Driver, dsn, err)
	}

	// Apply connection‐pool settings
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return db, nil
}
