CREATE TABLE IF NOT EXISTS measurement (
    id SERIAL PRIMARY KEY,
    created TIMESTAMPTZ NOT NULL DEFAULT now(),
    modified TIMESTAMPTZ NOT NULL DEFAULT now(),
    key VARCHAR(50) UNIQUE NOT NULL,
    value VARCHAR(500),
    CHECK(key ~ '^\d{4}-\d{4}-\d{4}(:[a-z]+)?$')
)
;
