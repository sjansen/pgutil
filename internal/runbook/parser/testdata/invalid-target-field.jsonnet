{
  targets: {
    sh: {
      class: 'sh',
      config: {
        shell: 'bash',
      },
    },
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
