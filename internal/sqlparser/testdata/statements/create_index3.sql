CREATE INDEX foo__uniq__key__null_value
ON foo ( key )
WHERE value IS NULL
