{
    tasks: {
        'begin': {
            after: [],
            exec: ['date'],
        },
        'end': {
            after: ['hostname', 'whoami'],
            exec: ['date'],
        },
        'hostname': {
            after: ['begin'],
            exec: ['hostname'],
        },
        'whoami': {
            after: ['begin'],
            exec: ['whoami'],
        },
    },
}
