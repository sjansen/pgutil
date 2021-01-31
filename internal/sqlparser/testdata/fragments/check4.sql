CHECK (
CASE
    WHEN key::text = 'begin'::text THEN 0 <= value
    ELSE 0 < value
END)
