CREATE TRIGGER update_foo_modified
BEFORE UPDATE ON foo
FOR EACH ROW
EXECUTE PROCEDURE update_modified_column()
