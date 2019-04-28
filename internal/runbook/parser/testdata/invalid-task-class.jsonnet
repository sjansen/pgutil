{
  targets: {
    sh: { class: 'sh' },
  },
  tasks: {
    timestamp: {
      target: 'sh/time',
      config: {
        args: ['ddate'],
      },
    },
  },
}
