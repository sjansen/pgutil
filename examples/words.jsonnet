local command(args, after=null) = {
  [if after != null then 'after']: after,
  target: 'sh',
  config: {
    args: args,
  },
};

local HR = '====================';
local label(text, after=null) = {
  [if after != null then 'after']: after,
  target: 'sh',
  config: {
    args: ['echo', HR, text, HR],
  },
};

local query(sql, after=null) = {
  [if after != null then 'after']: after,
  target: 'pg',
  config: {
    sql: sql,
  },
};

{
  targets: {
    pg: { class: 'pg' },
    sh: { class: 'sh' },
  },
  tasks: {
    'before/label': label('BEFORE'),
    'before/data': command(['psql', '-c', 'SELECT * FROM pgutil_demo.words ORDER BY timestamp'], after=['before/label']),
    'after/label': label('AFTER', after=['foo', 'bar', 'baz']),
    'after/data': command(['psql', '-c', 'SELECT * FROM pgutil_demo.words ORDER BY timestamp'], after=['after/label']),
    init: query(importstr 'init.sql'),
    foo: query("INSERT INTO pgutil_demo.words (value) VALUES ('foo')", after=['init', 'before/data']),
    bar: query("INSERT INTO pgutil_demo.words (value) VALUES ('bar')", after=['init', 'before/data']),
    baz: query("INSERT INTO pgutil_demo.words (value) VALUES ('baz')", after=['init', 'before/data']),
  },
}
