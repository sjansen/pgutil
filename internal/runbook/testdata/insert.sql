SELECT pg_sleep(5 * random())
;

INSERT INTO measurements
    (timestamp, value)
VALUES
    (now(), random())
;
