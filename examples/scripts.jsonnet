{
  targets: {
    sh: {
      class: 'sh',
      config: {
        env: {
          set: {
            FOO: '42',
          },
        },
      },
    },
  },
  tasks: {
    'debug-env': {
      target: 'sh',
      config: {
        args: ['./scripts/demo-script'],
        env: {
          set: {
            BAR: 'baz',
          },
        },
      },
    },
  },
}
