package configs

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(cfg *EnvConfig) {
	dsn := fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s", cfg.DBType, cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	log.Printf("Running migrations with DSN: %s", dsn)

	m, err := migrate.New(
		"file:///home/fnp/Assignments/A1_Go_CRUD_SQL_Basic_Kaffka/backend/migrations",
		dsn,
	)
	if err != nil {
		log.Fatalf("Migration init failed: %v", err)
	}
	defer m.Close()

	version, dirty, _ := m.Version()
	log.Printf("Current migration version: %v, dirty: %v", version, dirty)

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations applied successfully")
}
