-- +migrate Up
CREATE TABLE IF NOT EXISTS entity.pet
(
    id         BIGSERIAL   NOT NULL PRIMARY KEY,
    name       VARCHAR(50) NOT NULL,
    category   VARCHAR(50) NOT NULL,
    age        INTEGER     NOT NULL,
    weight     float8      NOT NULL,
    sexo       VARCHAR(20) NOT NULL,
    "user"     BIGINT      NOT NULL,
    type       VARCHAR(50) NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT now(),
    updated_at TIMESTAMP   NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS entity.pet;
