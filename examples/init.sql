CREATE SCHEMA IF NOT EXISTS pgutil_demo
;
CREATE TABLE IF NOT EXISTS pgutil_demo.words (
        id BIGSERIAL NOT NULL,
        timestamp TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        value VARCHAR(10) NOT NULL
)
;
DELETE FROM pgutil_demo.words
WHERE timestamp < now() - '1 minutes'::interval
;
