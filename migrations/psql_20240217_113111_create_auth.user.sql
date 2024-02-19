
-- +migrate Up
CREATE TABLE IF NOT EXISTS auth.user(
    id BIGSERIAL  NOT NULL PRIMARY KEY,
    "name" VARCHAR (50) NOT NULL,
    lastname VARCHAR (50) NOT NULL,
    email VARCHAR (100) NOT NULL,
    cellphone VARCHAR (50) NOT NULL,
    password VARCHAR (255) NOT NULL,
    age INTEGER  NOT NULL,
    city VARCHAR (100) NOT NULL,
    department VARCHAR (100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS auth.user;
