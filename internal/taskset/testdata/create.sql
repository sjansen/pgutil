CREATE TABLE IF NOT EXISTS measurements (
        id BIGSERIAL NOT NULL,
        timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
        value INTEGER NOT NULL
);
