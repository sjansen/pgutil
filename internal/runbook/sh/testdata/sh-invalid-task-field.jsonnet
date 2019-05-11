{
  targets: {
    sh: { class: 'sh' },
  },
  tasks: {
    portal: {
      target: 'sh',
      config: {
        cmd: ['date'],
      },
    },
  },
}
