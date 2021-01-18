package sqlparser

import (
	"github.com/sjansen/pgutil/internal/schema"
)

func newForeignKey(opts ...*option) *schema.ForeignKey {
	fk := &schema.ForeignKey{}
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
