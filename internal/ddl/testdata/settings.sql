CREATE OR REPLACE FUNCTION update_mtime_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.mtime = now();
  RETURN NEW;
END;
$$ LANGUAGE 'plpgsql'
;
CREATE TABLE "settings" (
  "ctime" TIMESTAMPTZ NOT NULL DEFAULT now()
, "mtime" TIMESTAMPTZ NOT NULL DEFAULT now()
, "key" VARCHAR(50) NOT NULL
, "value" VARCHAR(500)
, CHECK(length(key) > 0)
)
;
ALTER TABLE "settings"
  ADD CONSTRAINT "settings_pkey"
  PRIMARY KEY ("key")
;
CREATE TRIGGER "update_settings_mtime"
  BEFORE UPDATE ON "settings"
  FOR EACH ROW
  EXECUTE PROCEDURE update_mtime_column()
;
