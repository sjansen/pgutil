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

table "public" "bar" {
  column "id" {
    type = "integer"
    not_null = true
  }
  column "foo_id" {
    type = "integer"
    not_null = true
  }
  foreign_key "foo" {
    columns = ["foo_id"]
    referenced = ["id"]
  }
}

table "public" "foo" {
  comment = "A simple test case"

  column "id" {
    type = "integer"
    not_null = true
  }
  column "created" {
    type = "timestamp with time zone"
    not_null = true
    default = "now()"
  }
  column "modified" {
    type = "timestamp with time zone"
    not_null = true
    default = "now()"
  }
  column "key" {
    type = "character varying(50)"
    not_null = true
  }
  column "value" {
    type = "character varying(500)"
  }
}

trigger "public" "foo" "update_foo_modified" {
  function     = "update_modified_column"
  when         = "before"
  update       = true
  for_each_row = true
}
