{
  debug: true,
  targets: {
    sh: { class: 'sh' },
  },
  tasks: {
    timestamp: {
      target: 'sh',
      config: {
        args: ['ddate'],
      },
    },
  },
}
