package pg

import "errors"

var ErrNoHostForTLS = errors.New("host server name must be provided when TLS is required")
