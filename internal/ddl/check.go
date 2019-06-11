package ddl

func (c *Check) setExpression(s string) {
	c.Expression = trim(s)
}
