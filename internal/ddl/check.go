package ddl

// A Check constraint limits column values
type Check struct {
	Name              string `hcl:"name,optional"`
	Expression        string `hcl:"expression,attr"`
	Deferrable        bool   `hcl:"deferrable,optional"`
	InitiallyDeferred bool   `hcl:"initially_deferred,optional"`
}

func (c *Check) setExpression(s string) {
	c.Expression = trim(s)
}
