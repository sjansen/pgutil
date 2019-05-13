{
  targets: {
    demo: {
      class: 'demo',
      config: {
        string: '.ravgyniB ehbl xaveq bg rehf rO',
      },
    },
  },
  tasks: {
    decrypted: {
      target: 'demo/echo',
      after: ['reverse', 'rotate'],
    },
    encrypted: {
      target: 'demo/echo',
    },
    reverse: {
      after: ['encrypted'],
      target: 'demo/rev',
    },
    rotate: {
      after: ['encrypted'],
      target: 'demo/rot13',
    },
  },
}
