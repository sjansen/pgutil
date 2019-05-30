parameters {
  search_path = ["$user", "public"]
}

schema "public" {}

function "public" "update_modified_column" {
  returns = "trigger"
  language = "plpgsql"
  definition = <<EOF
BEGIN
  NEW.modified = now();
  RETURN NEW;
END;
EOF
}

table "public" "foo" {
  comment = "A simple test case"

  columns = ["id", "created", "modified", "key", "value"]
}

trigger "public" "foo" "update_foo_modified" {
  function     = "update_modified_column"
  when         = "before"
  update       = true
  for_each_row = true
}
