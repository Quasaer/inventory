package postgres

import (
	"fmt"

	"github.com/Quasaer/goinventory-api/goinventory"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InventoryItemStore struct {
	*sqlx.DB
}

func (s *InventoryItemStore) GetInventoryItem(id uuid.UUID) (goinventory.InventoryItem, error) {
	var inventoryItem goinventory.InventoryItem
	if err := s.Get(&inventoryItem, `SELECT * FROM inventory_item WHERE id = $1`, id); err != nil {
		return goinventory.InventoryItem{}, fmt.Errorf("ERROR RETREIVING INVENTORY ITEM NOT FOUND: %w", err)
	}
	return inventoryItem, nil
}

func (s *InventoryItemStore) GetAllInventoryItemsByInventoryListID(id uuid.UUID) ([]goinventory.InventoryItem, error) {
	var it []goinventory.InventoryItem

	if err := s.Select(&it, `SELECT * FROM inventory_item where inventory_list_id = $1`, id); err != nil {
		return []goinventory.InventoryItem{}, fmt.Errorf("ERROR RETREIVING INVENTORY ITEM NOT FOUND: %w", err)
	}
	return it, nil
}

func (s *InventoryItemStore) CreateInventoryItem(i *goinventory.InventoryItem) error {
	if err := s.Get(i, `INSERT INTO inventory_item VALUES ($1,$2,$3) RETURNING *`,
		i.ID,
		i.Name,
		i.CreatedAt); err != nil {
		return fmt.Errorf("ERROR CREATING INVENTORY ITEM: %w", err)
	}
	return nil
}

func (s *InventoryItemStore) UpdateInventoryItem(i *goinventory.InventoryItem) error {
	if err := s.Get(i, `UPDATE inventory_item SET name= $1 , count = $2 updatedAt = $3 where id = $4  RETURNING *`,
		i.Name,
		i.Count,
		i.UpdatedAt,
		i.ID); err != nil {
		return fmt.Errorf("ERROR UPDATING INVENTORY ITEM: %w", err)
	}
	return nil
}

func (s *InventoryItemStore) DeleteInventoryItem(InventoryItemId uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM inventory_item WHERE id = $1`, InventoryItemId); err != nil {
		return fmt.Errorf("ERROR DELETING INVENTORY ITEM: %w", err)
	}
	return nil
}
