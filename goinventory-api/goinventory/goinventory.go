package goinventory

import (
	"github.com/google/uuid"
)

type InventoryItem struct {
	ID              uuid.UUID `db:"id"`
	InventoryListId uuid.UUID `db:"inventory_list_id"`
	Name            string    `db:"name"`
	Count           int       `db:"count"`
	CreatedAt       int64     `db:"created_at"`
	UpdatedAt       int64     `db:"updated_at"`
}

type InventoryList struct {
	ID          uuid.UUID `db:"id" json:"ID"`
	Name        string    `db:"name" json:"Name" validate:"required"`
	Description string    `db:"description" json:"Description"`
	CreatedAt   int64     `db:"created_at" json:"CreatedAt"`
	UpdatedAt   int64     `db:"updated_at" json:"UpdatedAt"`
}

type InventoryListStore interface {
	GetInventoryListByID(id uuid.UUID) (InventoryList, error)
	GetAllInventoryListsByUserID(id uuid.UUID) ([]InventoryList, error)
	CreateInventoryList(i *InventoryList) error
	UpdateInventoryList(i *InventoryList) error
	DeleteInventoryList(InventoryListId uuid.UUID) error
	ValidateInventoryListOnCreate(i *InventoryList) error
	ValidateInventoryListOnUpdate(i *InventoryList) error
}

type InventoryItemStore interface {
	GetInventoryItem(id uuid.UUID) (InventoryItem, error)
	GetAllInventoryItemsByInventoryListID(id uuid.UUID) ([]InventoryItem, error)
	CreateInventoryItem(i *InventoryItem) error
	UpdateInventoryItem(i *InventoryItem) error
	DeleteInventoryItem(InventoryItemId uuid.UUID) error
}

type Store interface {
	InventoryListStore
	InventoryItemStore
}
