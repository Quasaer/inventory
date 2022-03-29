CREATE TABLE inventory_list (
    id UUID PRIMARY KEY NOT NULL,
    name text NOT NULL,
    description TEXT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT
);

CREATE TABLE inventory_item (
    id UUID PRIMARY KEY,
    inventory_list_id UUID NOT NULL REFERENCES inventory_list(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    count INT NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);