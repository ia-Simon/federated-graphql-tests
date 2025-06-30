CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "user" (
    id uuid NOT NULL PRIMARY KEY,
    username VARCHAR(96) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    name VARCHAR(128) NOT NULL,
    access_type VARCHAR(32) NOT NULL
);

CREATE TABLE IF NOT EXISTS "todo" (
    id uuid NOT NULL PRIMARY KEY,
    text_data text NOT NULL,
    is_done bool NOT NULL DEFAULT FALSE,
    user_id uuid NOT NULL REFERENCES "user" (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS "product" (
    sku uuid NOT NULL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    price BIGINT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "review" (
    id uuid NOT NULL PRIMARY KEY,
    rating SMALLINT NOT NULL,
    comment text NOT NULL,
    product_sku uuid NOT NULL REFERENCES "product" (sku) ON DELETE CASCADE ON UPDATE CASCADE,
    user_id uuid NOT NULL REFERENCES "user" (id) ON DELETE CASCADE ON UPDATE CASCADE
);
