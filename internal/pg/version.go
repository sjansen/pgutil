package pg

import "context"

// ServerVersion queries the database host for its server version
func (c *Conn) ServerVersion(ctx context.Context) (string, error) {
	c.log.Infow("requesting server version")

	var version string
	c.log.Debugw("executing query", "query", "SELECT VERSION()")
	err := c.conn.QueryRow(ctx, "SELECT VERSION()").Scan(&version)
	if err != nil {
		return "", err
	}

	c.log.Debugf("server version = %q", version)
	return version, nil
}
