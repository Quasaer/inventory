CREATE TABLE inventory_list (
    id UUID PRIMARY KEY,
    name text NOT NULL,
    description TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE inventory_item (
    id UUID PRIMARY KEY,
    inventory_list_id UUID NOT NULL REFERENCES inventory_list(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    count INT DEFAULT 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);