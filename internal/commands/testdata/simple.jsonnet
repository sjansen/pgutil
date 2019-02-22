{
    databases: {},
    tasks: {
        "init": {
            sql: importstr 'init.sql',
        },
        "bar": {
            after: ["foo"],
            sql: "INSERT INTO words (value) VALUES ('bar')",
        },
        "baz": {
            after: ["bar"],
            sql: "INSERT INTO words (value) VALUES ('baz')",
        },
        "foo": {
            after: ["init"],
            sql: "INSERT INTO words (value) VALUES ('foo')",
        },
        "qux": {
            after: ["baz"],
            exec: ["echo", "qux"],
        },
        "quux": {
            after: ["qux"],
            exec: ["echo", "quux"],
        },
    },
}
