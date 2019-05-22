{
  targets: {
    t: { class: 'demo' },
  },
  tasks: {
    step1: {
      target: 't/sleep',
      message: 'Tuning frobnicator...',
    },
    step2: {
      target: 't/sleep',
      after: ['step1'],
      message: 'Resetting twiddler...',
    },
    step3: {
      target: 't/sleep',
      after: ['step2'],
      message: 'Toggling gonkulator...',
    },
  },
}
