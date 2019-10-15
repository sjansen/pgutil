
parameters {
  search_path = ["$user", "public"]
  timezone    = ""
}

schema "public" {
  owner   = ""
  comment = ""
}

function "public" "update_modified_column" {
  owner      = ""
  comment    = ""
  returns    = "trigger"
  language   = "plpgsql"
  definition = "BEGIN\n  NEW.modified = now();\n  RETURN NEW;\nEND;\n"
}

sequence "public" "bar_id_seq" {
  owner     = ""
  comment   = ""
  data_type = ""
  start     = 1
  minimum   = 1
  maximum   = 2147483647
  increment = 1
  cache     = 1
  cycle     = false

  owned_by {
    schema = "public"
    table  = "bar"
    column = "id"
  }
}
sequence "public" "foo_id_seq" {
  owner     = ""
  comment   = ""
  data_type = ""
  start     = 0
  minimum   = 0
  maximum   = 2147483647
  increment = 0
  cache     = 0
  cycle     = false
}

table "public" "bar" {
  owner   = ""
  comment = ""

  column "id" {
    type     = "integer"
    not_null = true
    default  = "nextval('bar_id_seq'::regclass)"
    comment  = ""
  }
  column "foo_id" {
    type     = "integer"
    not_null = true
    default  = ""
    comment  = ""
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
  owner   = ""
  comment = "A simple test case"

  column "id" {
    type     = "integer"
    not_null = true
    default  = "nextval('foo_id_seq'::regclass)"
    comment  = ""
  }
  column "created" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
    comment  = ""
  }
  column "modified" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
    comment  = ""
  }
  column "key" {
    type     = "character varying(50)"
    not_null = true
    default  = ""
    comment  = ""
  }
  column "value" {
    type     = "character varying(500)"
    not_null = false
    default  = ""
    comment  = ""
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
