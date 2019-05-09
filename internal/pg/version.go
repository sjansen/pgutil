package pg

func (c *Conn) ServerVersion() (string, error) {
	c.log.Infow("requesting server version")

	var version string
	c.log.Debugw("executing query", "query", "SELECT VERSION()")
	err := c.conn.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "", err
	}

	c.log.Debugf("server version = %q", version)
	return version, nil
}
