
-- +migrate Up
CREATE TABLE IF NOT EXISTS entity.vaccines(
    id BIGSERIAL  NOT NULL PRIMARY KEY,
    name VARCHAR (100) NOT NULL,
    veterinary BIGINT  NOT NULL,
    doctor VARCHAR (100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS entity.vaccines;
