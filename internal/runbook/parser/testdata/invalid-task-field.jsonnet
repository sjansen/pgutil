{
  targets: {
    sh: { class: 'sh' },
  },
  tasks: {
    timestamp: {
      target: 'sh',
      config: {
        command: 'ddate',
      },
    },
  },
}
