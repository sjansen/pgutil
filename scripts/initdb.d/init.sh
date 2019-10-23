#!/bin/bash
set -e

cd $(dirname "$0")

if [ -z "$POSTGRES_USER" ]; then
  POSTGRES_USER="$USER"
fi

psql \
  -v ON_ERROR_STOP=1 \
  --username "$POSTGRES_USER" \
  --dbname "$POSTGRES_DB" <<-EOF
	CREATE DATABASE pgutil_test_basic;
	CREATE DATABASE pgutil_test_complete;
  UPDATE pg_database SET datistemplate=true WHERE datname='pgutil_test_basic';
  UPDATE pg_database SET datistemplate=true WHERE datname='pgutil_test_complete';
EOF

echo "=== Basic Database ==="
psql \
  -v ON_ERROR_STOP=1 \
  --username "$POSTGRES_USER" \
  --dbname "pgutil_test_basic" \
  -f "sql/basic.sql"

echo "=== Complete Database ==="
psql \
  -v ON_ERROR_STOP=1 \
  --username "$POSTGRES_USER" \
  --dbname "pgutil_test_complete" \
  -f "sql/complete.sql"
