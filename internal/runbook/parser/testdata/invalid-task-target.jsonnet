{
  targets: {
    sh: { class: 'sh' },
  },
  tasks: {
    timestamp: {
      target: 'bash',
      config: {
        args: ['ddate'],
      },
    },
  },
}
