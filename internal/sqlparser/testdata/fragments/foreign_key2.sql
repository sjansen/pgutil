FOREIGN KEY (bar, baz)
REFERENCES foo(qux, quux)
MATCH FULL
ON UPDATE SET DEFAULT
ON DELETE CASCADE
DEFERRABLE INITIALLY DEFERRED