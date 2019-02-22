CREATE TABLE IF NOT EXISTS words (
        id BIGSERIAL NOT NULL,
        timestamp TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        value VARCHAR(10) NOT NULL
)
;
DELETE FROM words
WHERE timestamp < now() - interval '5 minutes'
;
