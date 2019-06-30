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
    sql = "SELECT now()"
}

task "pg" "insert-new-measurements" {
    target = "src"
    after = ["create-table"]
    sql = "SELECT now() + '1 hour'"
}

task "pg" "delete-old-measurements" {
    target = "src"
    after = ["insert-new-measurements"]
    sql = "SELECT now() + '1 day'"
}
