{
  targets: {
    pg: {
      class: 'pg',
      config: {
        connect_retries: 5,
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
