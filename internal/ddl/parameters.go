package ddl

// Parameters describes database-level configuration options
type Parameters struct {
	SearchPath []string `hcl:"search_path,optional"`
	Timezone   string   `hcl:"timezone,optional"`
}
