ALTER TABLE inventory_list ALTER created_at TIMESTAMP NOT NULL;
ALTER TABLE inventory_list ALTER updated_at TIMESTAMP;
ALTER TABLE inventory_item ALTER created_at TIMESTAMP NOT NULL;
ALTER TABLE inventory_item ALTER updated_at TIMESTAMP;