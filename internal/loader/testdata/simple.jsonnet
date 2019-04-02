local directory = '/tmp/pgutil-simple-example';
{
    queues: {
        pg: import 'databases.json',
    },
    tasks: {
        'create-dir': {
            queue: 'sh',
            config: {
                args: ['mkdir', directory],
            },
        },
        'remove-dir': {
            after: ['delete-old-measurements', 'insert-new-measurements'],
            queue: 'sh',
            config: {
                args: ['rmdir', directory],
            },
        },
        'create-table': {
            queue: 'pg/src',
            config: {
                sql: importstr 'scripts/create.sql',
            },
        },
        'insert-new-measurements': {
            after: ['create-table'],
            queue: 'pg/src',
            config: {
                sql: importstr 'scripts/insert.sql',
            },
        },
        'delete-old-measurements': {
            after: ['insert-new-measurements'],
            queue: 'pg/src',
            config: {
                sql: importstr 'scripts/delete.sql',
            },
        },
    },
}
