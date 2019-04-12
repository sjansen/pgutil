{
    queues: {
        strbuf: {
            '': {
                message: ".ravgyniB ehbl xaveq bg rehf rO",
            },
        },
    },
    tasks: {
        'encrypted': {
            queue: 'strbuf',
            type: 'strbuf/echo',
        },
        'decrypted': {
            after: ['reverse', 'rotate'],
            queue: 'strbuf',
            type: 'strbuf/echo',
        },
        'reverse': {
            after: ['encrypted'],
            queue: 'strbuf',
            type: 'strbuf/rev',
        },
        'rotate': {
            after: ['encrypted'],
            queue: 'strbuf',
            type: 'strbuf/rot13',
        },
    }
}
