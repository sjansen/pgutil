
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

table "public" "measurement" {
  comment = ""
  owner   = "docker"

  column "created" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
  }
  column "id" {
    type     = "integer"
    not_null = true
    default  = "nextval('measurement_id_seq'::regclass)"
  }
  column "key" {
    type     = "character varying(50)"
    not_null = true
    default  = ""
  }
  column "modified" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
  }
  column "value" {
    type     = "character varying(500)"
    not_null = false
    default  = ""
  }
}
table "public" "observation" {
  comment = ""
  owner   = "docker"

  column "created" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
  }
  column "id" {
    type     = "integer"
    not_null = true
    default  = "nextval('observation_id_seq'::regclass)"
  }
  column "measurement_id" {
    type     = "integer"
    not_null = true
    default  = ""
  }
  column "modified" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
  }
  column "notes" {
    type     = "character varying(500)"
    not_null = false
    default  = ""
  }

  foreign_key "measurement" {
    name               = "observation_measurement_id_fkey"
    columns            = ["measurement_id"]
    referenced         = ["id"]
    match              = ""
    on_delete          = ""
    on_update          = ""
    deferrable         = false
    initially_deferred = false
  }
}

index "" "measurement" "measurement_pkey" {
  primary = true
  unique  = true
  using   = "btree"
  where   = ""

  key {
    column     = "id"
    expression = ""
    opclass    = ""
    descending = false
  }
}
index "" "measurement" "measurement_key_key" {
  primary = false
  unique  = true
  using   = "btree"
  where   = ""

  key {
    column     = "key"
    expression = ""
    opclass    = ""
    descending = false
  }
}
index "" "observation" "observation_pkey" {
  primary = true
  unique  = true
  using   = "btree"
  where   = ""

  key {
    column     = "id"
    expression = ""
    opclass    = ""
    descending = false
  }
}

trigger "" "measurement" "update_measurement_modified" {
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
trigger "" "observation" "update_observation_modified" {
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
