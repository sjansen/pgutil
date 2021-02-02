package sqlparser

import "github.com/sjansen/pgutil/internal/ddl"

func newForeignKey(opts ...*option) *ddl.ForeignKey {
	fk := &ddl.ForeignKey{}
	for _, opt := range opts {
		switch opt.Name {
		case "columns":
			fk.Columns = opt.Value.([]string)
		case "deferrable":
			fk.Deferrable = opt.Value.(bool)
		case "initially_deferred":
			fk.InitiallyDeferred = opt.Value.(bool)
		case "match":
			fk.Match = opt.Value.(string)
		case "on_delete":
			fk.OnDelete = opt.Value.(string)
		case "on_update":
			fk.OnUpdate = opt.Value.(string)
		case "referenced":
			fk.Referenced = opt.Value.([]string)
		case "table":
			fk.Table = opt.Value.(string)
		}
	}
	return fk
}

func newTrigger(opts ...*option) *ddl.Trigger {
	t := &ddl.Trigger{}
	for _, opt := range opts {
		switch opt.Name {
		case "delete":
			t.Delete = opt.Value.(bool)
		case "insert":
			t.Insert = opt.Value.(bool)
		case "truncate":
			t.Truncate = opt.Value.(bool)
		case "update":
			t.Update = opt.Value.(bool)
		case "columns":
			t.Columns = opt.Value.([]string)
		}
	}
	return t
}
