CREATE TABLE IF NOT EXISTS foo (
    id SERIAL NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT now(),
    modified TIMESTAMPTZ NOT NULL DEFAULT now(),
    key VARCHAR(50) UNIQUE NOT NULL,
    value VARCHAR(500)
)
;
CREATE OR REPLACE FUNCTION update_modified_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.modified = now();
    RETURN NEW; 
END;
$$ language 'plpgsql'
;
CREATE TRIGGER update_foo_modified
  BEFORE UPDATE ON foo
  FOR EACH ROW
  EXECUTE PROCEDURE update_modified_column()
;
