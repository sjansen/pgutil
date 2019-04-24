{
  targets: {
    strbuf: {
      class: 'strbuf',
      config: {
        data: '.ravgyniB ehbl xaveq bg rehf rO',
      },
    },
  },
  tasks: {
    decrypted: {
      target: 'strbuf/echo',
      after: ['reverse', 'rotate'],
    },
    encrypted: {
      target: 'strbuf/echo',
    },
    reverse: {
      after: ['encrypted'],
      target: 'strbuf/rev',
    },
    rotate: {
      after: ['encrypted'],
      target: 'strbuf/rot13',
    },
  },
}
