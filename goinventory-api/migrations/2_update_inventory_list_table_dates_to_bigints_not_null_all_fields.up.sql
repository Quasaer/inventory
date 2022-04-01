ALTER TABLE inventory_list ALTER created_at BIGINT NOT NULL;
ALTER TABLE inventory_list ALTER updated_at BIGINT NOT NULL;
ALTER TABLE inventory_item ALTER created_at BIGINT NOT NULL;
ALTER TABLE inventory_item ALTER updated_at BIGINT NOT NULL;