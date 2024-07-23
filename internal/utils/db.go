package utils

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func InitDB(user, password, dbname, host, port string) (*sql.DB, error) {
	p, err := strconv.Atoi(port)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, p, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("сonnected to the database")
	runMigrations("", dbname, db)

	return db, nil
}

func runMigrations(sourceUrl, dbname string, db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create postgres driver: %v", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get work directory: %v", err)
	}

	// Строим абсолютный путь к директории миграций
	migrationsPath := filepath.Join(wd, "migrations")

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsPath, dbname, driver)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("could not run migrations: %v", err)
	}
	log.Println("Migrations ran successfully")
}
