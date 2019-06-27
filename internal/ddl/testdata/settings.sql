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
