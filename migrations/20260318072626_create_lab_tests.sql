-- +goose Up
CREATE TABLE lab_tests (
    id BIGSERIAL PRIMARY KEY,
    visit_id BIGINT NOT NULL REFERENCES visits(id) ON DELETE CASCADE,
    test_name VARCHAR(255) NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_lab_tests_visit_id ON lab_tests(visit_id);

-- +goose Down
DROP TABLE IF EXISTS lab_tests;
