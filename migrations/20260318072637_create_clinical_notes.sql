-- +goose Up
CREATE TABLE clinical_notes (
    id BIGSERIAL PRIMARY KEY,
    visit_id BIGINT NOT NULL REFERENCES visits(id) ON DELETE CASCADE,
    note TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_clinical_notes_visit_id ON clinical_notes(visit_id);

-- +goose Down
DROP TABLE IF EXISTS clinical_notes;
