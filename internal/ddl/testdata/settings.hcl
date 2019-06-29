function "public" "update_mtime_column" {
  returns = "TRIGGER"
  language = "plpgsql"
  definition = <<EOF
BEGIN
  NEW.mtime = now();
  RETURN NEW;
END;
EOF
}

table "" "settings" {
  column "ctime" {
    type     = "TIMESTAMPTZ"
    not_null = true
    default  = "now()"
  }
  column "mtime" {
    type     = "TIMESTAMPTZ"
    not_null = true
    default  = "now()"
  }
  column "key" {
    type     = "VARCHAR(50)"
    not_null = true
  }
  column "value" {
    type     = "VARCHAR(500)"
    not_null = false
    default  = ""
  }
  check {
    expression = "length(key) > 0"
  }
}

index "" "settings" "settings_pkey" {
  primary = true
  unique  = true
  key {
    column = "key"
  }
}

trigger "" "settings" "update_settings_mtime" {
  function           = "update_mtime_column"
  when               = "BEFORE"
  for_each_row       = true
  update             = true
}
