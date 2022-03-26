package postgres

import (
	"fmt"

	"github.com/Quasaer/goinventory-api/goinventory"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InventoryListStore struct {
	*sqlx.DB
}

func (s *InventoryListStore) GetInventoryListByID(id uuid.UUID) (goinventory.InventoryList, error) {
	var inventoryList goinventory.InventoryList
	if err := s.Get(&inventoryList, `SELECT * FROM inventory_list WHERE id = $1`, id); err != nil {
		return goinventory.InventoryList{}, fmt.Errorf("ERROR RETREIVING INVENTORY LIST NOT FOUND: %w", err)
	}
	return inventoryList, nil
}

func (s *InventoryListStore) GetAllInventoryListsByUserID(id uuid.UUID) ([]goinventory.InventoryList, error) {
	panic("not implemented") // TODO: Implement
}

func (s *InventoryListStore) CreateInventoryList(i *goinventory.InventoryList) error {
	if err := s.Get(i, `INSERT INTO inventory_list VALUES ($1,$2,$3,$4) RETURNING *`,
		i.ID,
		i.Name,
		i.Description,
		i.CreatedAt); err != nil {
		return fmt.Errorf("ERROR CREATING INVENTORY LIST: %w", err)
	}
	return nil
}

func (s *InventoryListStore) UpdateInventoryList(i *goinventory.InventoryList) error {
	if err := s.Get(i, `UPDATE inventory_list SET name= $1 , description = $2, updatedAt = $3 where id = $4  RETURNING *`,
		i.Name,
		i.Description,
		i.UpdatedAt,
		i.ID); err != nil {
		return fmt.Errorf("ERROR UPDATING INVENTORY LIST: %w", err)
	}
	return nil
}

func (s *InventoryListStore) DeleteInventoryList(InventoryListId uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM inventory_list WHERE id = $1`, InventoryListId); err != nil {
		return fmt.Errorf("ERROR DELETING INVENTORY LIST: %w", err)
	}
	return nil
}
