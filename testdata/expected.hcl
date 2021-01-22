
parameters {
  search_path = ["\"$user\"", " public"]
  timezone    = "UTC"
}

schema "public" {
  owner   = "docker"
  comment = "standard public schema"
}

function "public" "update_modified_column" {
  owner      = "docker"
  comment    = ""
  returns    = "trigger"
  language   = "plpgsql"
  definition = "\nBEGIN\n    NEW.modified = now();\n    RETURN NEW; \nEND;\n"
}

sequence "public" "measurement_id_seq" {
  owner     = "docker"
  comment   = ""
  data_type = "bigint"
  start     = 1
  minimum   = 1
  maximum   = 9223372036854775807
  increment = 1
  cache     = 1
  cycle     = false

  owned_by {
    schema = "public"
    table  = "measurement"
    column = "id"
  }
}
sequence "public" "observation_id_seq" {
  owner     = "docker"
  comment   = ""
  data_type = "bigint"
  start     = 1
  minimum   = 1
  maximum   = 9223372036854775807
  increment = 1
  cache     = 1
  cycle     = false

  owned_by {
    schema = "public"
    table  = "observation"
    column = "id"
  }
}

table "public" "measurement" {
  owner   = "docker"
  comment = ""

  column "created" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
    comment  = ""
  }
  column "id" {
    type     = "bigint"
    not_null = true
    default  = "nextval('measurement_id_seq'::regclass)"
    comment  = ""
  }
  column "key" {
    type     = "character varying(50)"
    not_null = true
    default  = ""
    comment  = ""
  }
  column "modified" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
    comment  = ""
  }
  column "value" {
    type     = "character varying(500)"
    not_null = false
    default  = ""
    comment  = ""
  }

  check {
    name               = "measurement_key_check"
    expression         = "key::text ~ '^\\d{4}-\\d{4}-\\d{4}(:[a-z]+)?$'::text"
    deferrable         = false
    initially_deferred = false
  }

  storage_parameters {
    fillfactor         = 75
    autovacuum_enabled = false
  }
}
table "public" "observation" {
  owner   = "docker"
  comment = ""

  column "created" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
    comment  = ""
  }
  column "id" {
    type     = "bigint"
    not_null = true
    default  = "nextval('observation_id_seq'::regclass)"
    comment  = ""
  }
  column "measurement_id" {
    type     = "integer"
    not_null = true
    default  = ""
    comment  = ""
  }
  column "modified" {
    type     = "timestamp with time zone"
    not_null = true
    default  = "now()"
    comment  = ""
  }
  column "notes" {
    type     = "character varying(500)"
    not_null = false
    default  = ""
    comment  = ""
  }

  check {
    name               = "encourage detailed notes"
    expression         = "length(notes::text) > 50"
    deferrable         = false
    initially_deferred = false
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

  storage_parameters {
  }
}

index "public" "measurement" "measurement_pkey" {
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
index "public" "measurement" "measurement_key_key" {
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
index "public" "observation" "observation_pkey" {
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

trigger "public" "measurement" "update_modified_column" {
  from               = ""
  function           = "update_modified_column"
  timing             = "BEFORE"
  when               = ""
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
trigger "public" "observation" "update_modified_column" {
  from               = ""
  function           = "update_modified_column"
  timing             = "BEFORE"
  when               = ""
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
