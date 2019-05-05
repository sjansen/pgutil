local ns = std.range(1, 9);

{
  targets: {
    ['t' + n]: {
      class: 'sh',
    }
    for n in ns
  },
  tasks: {
    begin: {
      target: 't1',
      config: {
        args: ['date'],
      },
    },
    end: {
      after: ['sleep/' + n for n in ns],
      target: 't1',
      config: {
        args: ['date'],
      },
    },
  } + {
    ['echo/' + n]: {
      after: ['begin'],
      target: 't' + n,
      config: {
        args: ['echo', '#' + n, 'sleeping for 1 second'],
      },
    }
    for n in ns
  } + {
    ['sleep/' + n]: {
      after: ['echo/' + n],
      target: 't' + n,
      config: {
        args: ['sleep', '1'],
      },
    }
    for n in ns
  },
}
