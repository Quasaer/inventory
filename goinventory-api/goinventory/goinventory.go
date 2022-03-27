package goinventory

import "github.com/google/uuid"

type InventoryItem struct {
	ID              uuid.UUID `db:"id"`
	InventoryListId uuid.UUID `db:"inventory_list_id"`
	Name            string    `db:"name"`
	Count           int       `db:"count"`
	CreatedAt       int       `db:"created_at"`
	UpdatedAt       int       `db:"updated_at"`
}

type InventoryList struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   int       `db:"created_at"`
	UpdatedAt   int       `db:"updated_at"`
}

type InventoryListStore interface {
	GetInventoryListByID(id uuid.UUID) (InventoryList, error)
	GetAllInventoryListsByUserID(id uuid.UUID) ([]InventoryList, error)
	CreateInventoryList(i *InventoryList) error
	UpdateInventoryList(i *InventoryList) error
	DeleteInventoryList(InventoryListId uuid.UUID) error
}

type InventoryItemStore interface {
	GetInventoryItem(id uuid.UUID) (InventoryItem, error)
	GetAllInventoryItemsByInventoryListID(id uuid.UUID) ([]InventoryItem, error)
	CreateInventoryItem(i *InventoryItem) error
	UpdateInventoryItem(i *InventoryItem) error
	DeleteInventoryItem(InventoryItemId uuid.UUID) error
}

type Store interface {
	InventoryItemStore
	InventoryItemStore
}
