{
  targets: {
    pg: { class: 'pg' },
  },
  tasks: {
    portal: {
      target: 'pg/broadcast',
      config: {
        msg: 'start',
      },
    },
  },
}
