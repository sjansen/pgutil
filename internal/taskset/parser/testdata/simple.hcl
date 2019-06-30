target "pg" "src" {}

task "sh" "create-dir" {
    args = ["mkdir", "-p", "/tmp/pgutil-simple-example"]
}

task "sh" "remove-dir" {
    after = ["delete-old-measurements", "insert-new-measurements"]
    args = ["rmdir", "/tmp/pgutil-simple-example"]
}

task "pg" "create-table" {
    target = "src"
    sql = <<EOF
CREATE TABLE IF NOT EXISTS measurements (
  id BIGSERIAL NOT NULL,
  timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
  value DOUBLE PRECISION NOT NULL
)
EOF
}

task "pg" "insert-new-measurements" {
    target = "src"
    after = ["create-table"]
    sql = <<EOF
INSERT INTO measurements
  (timestamp, value)
VALUES
  (now(), random())
EOF
}

task "pg" "delete-old-measurements" {
    target = "src"
    after = ["insert-new-measurements"]
    sql = <<EOF
DELETE FROM measurements
 WHERE timestamp < now() - interval '5 minutes'
EOF
}
