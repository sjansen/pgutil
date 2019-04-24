local directory = '/tmp/pgutil-simple-example';
local databases = import 'databases.json';
{
  targets: {
    sh: {
      class: 'sh',
    },
    src: {
      class: 'pg',
      config: databases.src,
    },
    dst: {
      class: 'pg',
      config: databases.dst,
    },
  },
  tasks: {
    'create-dir': {
      target: 'sh',
      config: {
        args: ['mkdir', '-p', directory],
      },
    },
    'remove-dir': {
      after: ['delete-old-measurements', 'insert-new-measurements'],
      target: 'sh',
      config: {
        args: ['rmdir', directory],
      },
    },
    'create-table': {
      target: 'src/exec',
      config: {
        sql: importstr 'scripts/create.sql',
      },
    },
    'insert-new-measurements': {
      after: ['create-table'],
      target: 'src/exec',
      config: {
        sql: importstr 'scripts/insert.sql',
      },
    },
    'delete-old-measurements': {
      after: ['insert-new-measurements'],
      target: 'src/exec',
      config: {
        sql: importstr 'scripts/delete.sql',
      },
    },
  },
}
