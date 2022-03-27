CREATE TABLE inventory_list (
    id UUID PRIMARY KEY,
    name text NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE inventory_item (
    id UUID PRIMARY KEY,
    inventory_list_id UUID NOT NULL REFERENCES inventory_list(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    count INT NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP
);