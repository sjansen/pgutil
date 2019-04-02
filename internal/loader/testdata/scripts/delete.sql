SELECT pg_sleep(5 * random())
;

DELETE FROM measurements
WHERE timestamp < now() - interval '5 minutes'
;
