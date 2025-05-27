-- +migrate Up

CREATE TABLE IF NOT EXISTS cross_consultations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    prompt VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    status SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_cross_consultations_status ON cross_consultations (status);

-- +migrate Down

DROP INDEX IF EXISTS idx_cross_consultations_status;
DROP TABLE IF EXISTS cross_consultations;
