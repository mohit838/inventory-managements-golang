// Make package name main
package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	"github.com/mohit838/inventory-managements-golang/pkg/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}
	direction := os.Args[1] // "up" or "down"

	// 1) Load YAML config instead of .env
	cfg, err := config.AppConfig()
	if err != nil {
		log.Fatalf("failed to load YAML config: %v", err)
	}

	// 2) Build a MySQL DSN from cfg.Database
	dbCfg := cfg.Database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DbName,
		dbCfg.Charset,
		dbCfg.ParseTime,
		dbCfg.Loc,
	)

	// 3) Open a *sqlx.DB so we can hand off its underlying *sql.DB
	dbx, err := sqlx.Open(dbCfg.Driver, dsn)
	if err != nil {
		log.Fatalf("sqlx.Open error: %v", err)
	}
	defer dbx.Close()

	// Optional: ping to verify the connection
	if err := dbx.Ping(); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	// 4) Extract *sql.DB from *sqlx.DB
	stdDB := dbx.DB

	// 5) Create a golang-migrate MySQL driver
	driver, err := mysql.WithInstance(stdDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("mysql.WithInstance error: %v", err)
	}

	// 6) Point migrate at your migrations folder.
	//    Adjust "file://cmd/migrate/migrations" if your .up.sql/.down.sql live elsewhere.
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("migrate.NewWithDatabaseInstance error: %v", err)
	}

	// 7) Run up or down
	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migrate Up error: %v", err)
		}
		log.Println("Migrations applied (up).")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migrate Down error: %v", err)
		}
		log.Println("Migrations rolled back (down).")
	default:
		log.Fatal("Invalid direction! Use 'up' or 'down'")
	}
}
