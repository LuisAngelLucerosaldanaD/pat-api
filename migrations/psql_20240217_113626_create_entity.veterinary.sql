-- +migrate Up
CREATE TABLE IF NOT EXISTS entity.veterinary
(
    id          BIGSERIAL     NOT NULL PRIMARY KEY,
    name        VARCHAR(50)   NOT NULL,
    description VARCHAR(1000) NOT NULL,
    email       VARCHAR(100)  NOT NULL,
    address     VARCHAR(255)  NOT NULL,
    cellphone   VARCHAR(10)   NOT NULL,
    "user"      BIGINT        NOT NULL,
    web_page    VARCHAR(500)  NOT NULL,
    created_at  TIMESTAMP     NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP     NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS entity.veterinary;
