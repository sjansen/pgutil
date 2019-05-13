local target(s) = {
  class: 'demo',
  config: {
    data: s,
  },
};

{
  targets: {
    t1: target('foo'),
    t2: target('rab'),
    t3: target('mno'),
  },
  tasks: {
    // t1
    't1/1': {
      target: 't1/echo',
    },
    't1/2a': {
      after: ['t1/1'],
      target: 't1/rot13',
    },
    't1/2b': {
      after: ['t1/1'],
      target: 't1/rev',
    },
    't1/3': {
      after: ['t1/2a', 't1/2b'],
      target: 't1/sleep',
    },
    't1/4': {
      after: ['t1/3', 't2/5'],
      target: 't1/echo',
    },
    // t2
    't2/1': {
      target: 't2/rev',
    },
    't2/2': {
      after: ['t1/1', 't2/1'],
      target: 't2/echo',
    },
    't2/3': {
      after: ['t3/2'],
      target: 't2/rot13',
    },
    't2/4': {
      after: ['t3/2'],
      target: 't2/fail',
    },
    't2/5': {
      after: ['t2/4'],
      target: 't2/echo',
    },
    // t3
    't3/1a': {
      target: 't3/rot13',
    },
    't3/1b': {
      target: 't3/rev',
    },
    't3/2': {
      after: ['t2/2', 't3/1a', 't3/1b'],
      target: 't3/echo',
    },
    't3/3': {
      after: ['t3/2'],
      target: 't3/sleep',
    },
    't3/4': {
      after: ['t2/5', 't3/3'],
      target: 't3/echo',
    },
  },
}
