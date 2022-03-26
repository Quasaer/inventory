package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
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
