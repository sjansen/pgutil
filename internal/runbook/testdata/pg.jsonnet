{
  targets: {
    pg: { class: 'pg' },
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
