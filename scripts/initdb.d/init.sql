CREATE OR REPLACE FUNCTION update_modified_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.modified = now();
    RETURN NEW; 
END;
$$ language 'plpgsql'
;
CREATE TABLE IF NOT EXISTS measurement (
    id SERIAL PRIMARY KEY,
    created TIMESTAMPTZ NOT NULL DEFAULT now(),
    modified TIMESTAMPTZ NOT NULL DEFAULT now(),
    key VARCHAR(50) UNIQUE NOT NULL,
    value VARCHAR(500),
    CHECK(key ~ '^\d{4}-\d{4}-\d{4}(:[a-z]+)?$')
)
;
CREATE TRIGGER update_measurement_modified
  BEFORE UPDATE ON measurement
  FOR EACH ROW
  EXECUTE PROCEDURE update_modified_column()
;
CREATE TABLE IF NOT EXISTS observation (
    id SERIAL PRIMARY KEY,
    created TIMESTAMPTZ NOT NULL DEFAULT now(),
    modified TIMESTAMPTZ NOT NULL DEFAULT now(),
    measurement_id INTEGER NOT NULL REFERENCES measurement(id),
    notes VARCHAR(500),
    CONSTRAINT "encourage detailed notes"
    CHECK(length(notes) > 50)
)
;
CREATE TRIGGER update_observation_modified
  BEFORE UPDATE ON observation
  FOR EACH ROW
  EXECUTE PROCEDURE update_modified_column()
;
