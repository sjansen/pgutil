{
  targets: {
    pg: {
      class: 'pg',
      config: {
        host: 'invalid-hostname',
        database: 'invalid-database',

        connect_retries: 1,
        sslmode: 'allow',
      },
    },
  },
  tasks: {
    timestamp: {
      target: 'pg/exec',
      config: {
        sql: 'SELECT now()',
      },
    },
  },
}
