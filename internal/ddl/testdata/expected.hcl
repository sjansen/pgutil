
parameters {
  search_path = ["$user", "public"]
}

schema "public" {
  comment = ""
  owner   = ""
}

function "public" "update_modified_column" {
  comment    = ""
  owner      = ""
  returns    = "trigger"
  language   = "plpgsql"
  definition = "BEGIN\n  NEW.modified = now();\n  RETURN NEW;\nEND;\n"
}

table "public" "bar" {
  comment = ""
  owner   = ""

  column "id" {
    type     = "integer"
    not_null = true
    default  = ""
  }
  column "foo_id" {
    type     = "integer"
    not_null = true
    default  = ""
  }

  foreign_key "foo" {
    name               = ""
    columns            = ["foo_id"]
    referenced         = ["id"]
    match              = ""
    on_delete          = ""
    on_update          = ""
    deferrable         = false
    initially_deferred = false
  }
}
table "public" "foo" {
  comment = "A simple test case"
  owner   = ""

  column "id" {
    type     = "integer"
    not_null = true
    default  = ""
  }
  column "created" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
  }
  column "modified" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
  }
  column "key" {
    type     = "character varying(50)"
    not_null = true
    default  = ""
  }
  column "value" {
    type     = "character varying(500)"
    not_null = false
    default  = ""
  }
}

trigger "public" "foo" "update_foo_modified" {
  function           = "update_modified_column"
  when               = "before"
  constraint         = false
  deferrable         = false
  initially_deferred = false
  for_each_row       = true
  delete             = false
  insert             = false
  truncate           = false
  update             = true
  columns            = null
}
