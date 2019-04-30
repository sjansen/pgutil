local lines = [
  { line: 1, args: ['The world is a vampire, sent to drain'] },
  { line: 2, args: ['Secret destroyers, hold you up to the flames'] },
  { line: 3, args: ['And', 'what', 'do', 'I', 'get,', 'for', 'my', 'pain?'] },
  { line: 4, args: ['Betrayed', 'desires,', 'and', 'a', 'piece', 'of', 'the', 'game'] },
];

{
  targets: {
    sh: { class: 'sh' },
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
  },
}
