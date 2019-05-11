local lines = [
  { line: 1, args: ['The world is a vampire, sent to drain'] },
  { line: 2, args: ['Secret destroyers, hold you up to the flames'] },
  { line: 3, args: ['And', 'what', 'do', 'I', 'get,', 'for', 'my', 'pain?'] },
  { line: 4, args: ['Betrayed', 'desires,', 'and', 'a', 'piece', 'of', 'the', 'game\n'] },
];

{
  targets: {
    sh: {
      class: 'sh',
      config: {
        env: {
          allow: ['PGUTIL_CHORUS'],
          set: {
            PGUTIL_LINE5: "Even though I know, I suppose I'll show",
          },
        },
      },
    },
  },
  tasks: {
    ['line' + x.line]: {
      target: 'sh',
      [if x.line > 1 then 'after']: ['line' + (x.line - 1)],
      config: {
        args: ['echo'] + x.args,
      },
    }
    for x in lines
  } + {
    line5: {
      after: ['line4'],
      target: 'sh',
      config: {
        args: ['printenv', 'PGUTIL_LINE5'],
      },
    },
    line6: {
      after: ['line5'],
      target: 'sh',
      config: {
        args: ['printenv', 'PGUTIL_LINE6'],
        env: {
          set: {
            PGUTIL_LINE6: 'All my cool and cold--like old Job',
          },
        },
      },
    },
    line7: {
      after: ['line6'],
      target: 'sh',
      config: {
        args: ['printenv', 'PGUTIL_CHORUS'],
      },
    },
  },
}
