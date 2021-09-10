package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/ddl"
)

var hasSequencesView = `
SELECT EXISTS (
  SELECT 1
  FROM pg_catalog.pg_class c
  JOIN pg_catalog.pg_namespace n
    ON n.oid = c.relnamespace
  WHERE n.nspname = 'pg_catalog'
    AND c.relname = 'pg_sequence'
)
`

var listSequences = `
SELECT
  n.nspname as "Schema"
, c.relname as "Name"
, pg_catalog.pg_get_userbyid(c.relowner) as "Owner"
, pg_catalog.obj_description(c.oid, 'pg_class') as "Comment"
, pg_catalog.format_type(s.seqtypid, NULL) AS "Type"
, s.seqstart AS "Start"
, s.seqmin AS "Minimum"
, s.seqmax AS "Maximum"
, s.seqincrement AS "Increment"
, s.seqcache AS "Cache"
, s.seqcycle "Cycle"
, n2.nspname
, c2.relname
, a2.attname
FROM pg_catalog.pg_class c
JOIN pg_catalog.pg_namespace n
  ON n.oid = c.relnamespace
JOIN pg_catalog.pg_sequence s
  ON s.seqrelid = c.oid
LEFT JOIN pg_catalog.pg_depend d
  ON d.objid = c.oid
JOIN pg_catalog.pg_class c2
  ON c2.oid = d.refobjid
JOIN pg_catalog.pg_namespace n2
  ON n2.oid = c2.relnamespace
JOIN pg_catalog.pg_attribute a2
  ON (a2.attrelid = c2.oid AND a2.attnum = d.refobjsubid)
WHERE c.relkind = 'S'
  AND n.nspname <> 'pg_catalog'
  AND n.nspname <> 'information_schema'
  AND n.nspname !~ '^pg_toast'
  AND pg_catalog.pg_table_is_visible(c.oid)
  AND d.classid='pg_catalog.pg_class'::pg_catalog.regclass
  AND d.refclassid='pg_catalog.pg_class'::pg_catalog.regclass
  AND d.deptype IN ('a', 'i')
ORDER BY "Schema", "Name"
`

var listSequencesPre10 = `
SELECT
  n.nspname as "Schema"
, c.relname as "Name"
, pg_catalog.pg_get_userbyid(c.relowner) as "Owner"
, pg_catalog.obj_description(c.oid, 'pg_class') as "Comment"
, n2.nspname
, c2.relname
, a2.attname
FROM pg_catalog.pg_class c
JOIN pg_catalog.pg_namespace n
  ON n.oid = c.relnamespace
LEFT JOIN pg_catalog.pg_depend d
  ON d.objid = c.oid
JOIN pg_catalog.pg_class c2
  ON c2.oid = d.refobjid
JOIN pg_catalog.pg_namespace n2
  ON n2.oid = c2.relnamespace
JOIN pg_catalog.pg_attribute a2
  ON (a2.attrelid = c2.oid AND a2.attnum = d.refobjsubid)
WHERE c.relkind = 'S'
  AND n.nspname <> 'pg_catalog'
  AND n.nspname <> 'information_schema'
  AND n.nspname !~ '^pg_toast'
  AND pg_catalog.pg_table_is_visible(c.oid)
  AND d.classid='pg_catalog.pg_class'::pg_catalog.regclass
  AND d.refclassid='pg_catalog.pg_class'::pg_catalog.regclass
  AND d.deptype IN ('a', 'i')
ORDER BY "Schema", "Name"
`

// ListSequences describes the sequences in the database
func (c *Conn) ListSequences(ctx context.Context) ([]*ddl.Sequence, error) {
	c.log.Infow("ListSequences")

	var hasView bool

	c.log.Debugw("executing query", "query", hasSequencesView)
	err := c.conn.QueryRow(ctx, hasSequencesView).Scan(&hasView)
	if err != nil {
		return nil, err
	}
	if hasView {
		return c.listSequences(ctx)
	}
	return c.listSequencesLegacy(ctx)

}

func (c *Conn) listSequences(ctx context.Context) ([]*ddl.Sequence, error) {
	c.log.Debugw("executing query", "query", listSequences)
	rows, err := c.conn.Query(ctx, listSequences)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var sequences []*ddl.Sequence
	for rows.Next() {
		var (
			schema, name, owner, comment              *string
			dataType                                  *string
			start, minimum, maximum, increment, cache int64
			cycle                                     bool
			ownerSchema, ownerName, ownerColumn       *string
		)
		err = rows.Scan(
			&schema, &name, &owner, &comment,
			&dataType,
			&start, &minimum, &maximum, &increment, &cache,
			&cycle,
			&ownerSchema, &ownerName, &ownerColumn,
		)
		if err != nil {
			break
		}
		sequence := &ddl.Sequence{
			Schema:  String(schema),
			Name:    String(name),
			Owner:   String(owner),
			Comment: String(comment),

			DataType:  String(dataType),
			Start:     start,
			Minimum:   minimum,
			Maximum:   maximum,
			Increment: increment,
			Cache:     cache,
			Cycle:     cycle,
			OwnedBy: &ddl.SequenceOwner{
				Schema: String(ownerSchema),
				Table:  String(ownerName),
				Column: String(ownerColumn),
			},
		}
		c.log.Debugw("row scanned", "sequence", sequence)
		sequences = append(sequences, sequence)
	}
	if err != nil {
		return nil, err
	}

	return sequences, nil
}

func (c *Conn) listSequencesLegacy(ctx context.Context) ([]*ddl.Sequence, error) {
	c.log.Debugw("executing query", "query", listSequencesPre10)
	rows, err := c.conn.Query(ctx, listSequencesPre10)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var sequences []*ddl.Sequence
	for rows.Next() {
		var (
			schema, name, owner, comment        *string
			ownerSchema, ownerName, ownerColumn *string
		)
		err = rows.Scan(
			&schema, &name, &owner, &comment,
			&ownerSchema, &ownerName, &ownerColumn,
		)
		if err != nil {
			break
		}

		sequence := &ddl.Sequence{
			Schema:  String(schema),
			Name:    String(name),
			Owner:   String(owner),
			Comment: String(comment),

			DataType: "bigint",
			OwnedBy: &ddl.SequenceOwner{
				Schema: String(ownerSchema),
				Table:  String(ownerName),
				Column: String(ownerColumn),
			},
		}
		c.log.Debugw("row scanned", "sequence", sequence)
		sequences = append(sequences, sequence)
	}
	if err != nil {
		return nil, err
	}

	for _, seq := range sequences {
		query := `
SELECT
  start_value
, min_value
, max_value
, increment_by
, cache_value
, is_cycled
FROM ` + Identifier(seq.Schema, seq.Name)
		c.log.Debugw("executing query", "query", query)
		var (
			start, minimum, maximum, increment, cache int64
			cycle                                     bool
		)
		err := c.conn.QueryRow(ctx, query).Scan(
			&start, &minimum, &maximum, &increment, &cache,
			&cycle,
		)
		if err != nil {
			return nil, err
		}
		seq.Start = start
		seq.Minimum = minimum
		seq.Maximum = maximum
		seq.Increment = increment
		seq.Cache = cache
		seq.Cycle = cycle
	}

	return sequences, nil
}
