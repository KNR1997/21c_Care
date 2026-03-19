-- +goose Up
CREATE TABLE visits (
    id BIGSERIAL PRIMARY KEY,
    patient_id BIGINT NOT NULL,
    raw_input TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_patient
        FOREIGN KEY (patient_id)
        REFERENCES patients(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS visits;
