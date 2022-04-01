ALTER TABLE inventory_list 
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE inventory_list 
ADD COLUMN created_at BIGINT NOT NULL,
ADD COLUMN updated_at BIGINT NOT NULL;

ALTER TABLE inventory_item 
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE inventory_item
ADD COLUMN created_at BIGINT NOT NULL,
ADD COLUMN updated_at BIGINT NOT NULL;
