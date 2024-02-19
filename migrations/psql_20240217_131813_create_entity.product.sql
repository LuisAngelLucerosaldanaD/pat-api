-- +migrate Up
CREATE TABLE IF NOT EXISTS entity.product
(
    id           BIGSERIAL    NOT NULL PRIMARY KEY,
    name         VARCHAR(50)  NOT NULL,
    description  VARCHAR(500) NOT NULL,
    stock        INTEGER      NOT NULL,
    veterinary   BIGINT       NOT NULL,
    category     VARCHAR(50)  NOT NULL,
    price        float8       NOT NULL,
    type_product VARCHAR(50)  NOT NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at   TIMESTAMP    NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS entity.product;
