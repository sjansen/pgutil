{
  targets: {
    sh: { class: 'sh' },
  },
  tasks: {
    portal: {
      target: 'sh/sleep',
      config: {
        args: ['sleep', '60'],
      },
    },
  },
}
