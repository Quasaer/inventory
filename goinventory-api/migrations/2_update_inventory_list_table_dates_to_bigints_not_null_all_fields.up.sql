ALTER TABLE inventory_list ALTER created_at TYPE BIGINT;
ALTER TABLE inventory_list ALTER updated_at TYPE BIGINT;
ALTER TABLE inventory_list ALTER created_at SET NOT NULL;
ALTER TABLE inventory_list ALTER updated_at SET NOT NULL;

ALTER TABLE inventory_item ALTER created_at TYPE BIGINT;
ALTER TABLE inventory_item ALTER updated_at TYPE BIGINT;
ALTER TABLE inventory_item ALTER created_at SET NOT NULL;
ALTER TABLE inventory_item ALTER updated_at SET NOT NULL;