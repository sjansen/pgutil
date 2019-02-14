{
    'databases': import 'databases.json',
    'tasks': {
        'create-table': {
            'sql': importstr 'create.sql',
        },
        'delete-old-measurements': {
            'after': 'create-table',
            'sql': importstr 'delete.sql',
        },
        'insert-new-measurements': {
            'after': 'create-table',
            'sql': importstr 'insert.sql',
        },
    },
}
