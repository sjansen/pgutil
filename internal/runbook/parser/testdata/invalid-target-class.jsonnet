{
  targets: {
    sh: { class: 'bash' },
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
