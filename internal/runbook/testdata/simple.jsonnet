local directory = '/tmp/pgutil-simple-example';
{
    databases: import 'databases.json',
    tasks: {
        'create-dir': {
            exec: ['mkdir', directory],
        },
        'remove-dir': {
            after: ['delete-old-measurements', 'insert-new-measurements'],
            exec: ['rmdir', directory],
        },
        'create-table': {
            after: ['create-dir'],
            sql: importstr 'create.sql',
        },
        'delete-old-measurements': {
            after: ['create-table'],
            sql: importstr 'delete.sql',
        },
        'insert-new-measurements': {
            after: ['create-table'],
            sql: importstr 'insert.sql',
        },
    },
}
