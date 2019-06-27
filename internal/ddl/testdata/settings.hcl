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
