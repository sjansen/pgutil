CREATE CONSTRAINT TRIGGER trigger1
AFTER INSERT OR DELETE ON table1
DEFERRABLE INITIALLY DEFERRED
FOR EACH ROW
EXECUTE FUNCTION fn1()
