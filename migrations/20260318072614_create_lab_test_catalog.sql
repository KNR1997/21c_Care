-- +goose Up
CREATE TABLE lab_test_catalogs (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    default_price NUMERIC(10,2) NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE IF EXISTS lab_test_catalogs;
