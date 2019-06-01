
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
  columns = ["id", "created", "modified", "key", "value"]
}
