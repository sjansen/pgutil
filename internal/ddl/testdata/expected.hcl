
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

table "public" "foo" {
  comment = "A simple test case"
  owner   = ""
  columns = ["id", "created", "modified", "key", "value"]
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
