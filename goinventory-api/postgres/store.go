package postgres

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {

	m, err := migrate.New(
		"file://migrations",
		string(dataSourceName))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open("postgres", dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("ERROR OPENING DATABASE: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ERROR CONNECTING TO DATABASE: %w", err)
	}

	return &Store{
		InventoryListStore: &InventoryListStore{DB: db},
		InventoryItemStore: &InventoryItemStore{DB: db},
	}, nil
}

type Store struct {
	*InventoryItemStore
	*InventoryListStore
}
