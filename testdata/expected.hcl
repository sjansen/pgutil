
schema "public" {
  comment = "standard public schema"
  owner   = "docker"
}

function "public" "update_modified_column" {
  comment    = ""
  owner      = "docker"
  returns    = "trigger"
  language   = "plpgsql"
  definition = "\nBEGIN\n    NEW.modified = now();\n    RETURN NEW; \nEND;\n"
}

table "public" "foo" {
  comment = ""
  owner   = "docker"

  column "id" {
    type     = "integer"
    not_null = true
    default  = "nextval('foo_id_seq'::regclass)"
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

trigger "" "foo" "update_foo_modified" {
  function           = "update_modified_column"
  when               = "BEFORE"
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
