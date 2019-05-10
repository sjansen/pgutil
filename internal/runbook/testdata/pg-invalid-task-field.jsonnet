{
  targets: {
    pg: { class: 'pg' },
  },
  tasks: {
    timestamp: {
      target: 'pg/exec',
      config: {
        query: ['SELECT now()'],
      },
    },
  },
}
