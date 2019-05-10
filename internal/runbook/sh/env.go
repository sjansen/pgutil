package sh

import (
	"strings"
)

// Env describes how to filter environment variables
type Env struct {
	Allow []string
	Deny  []string
	Set   map[string]string
}

// Apply returns a filtered copy of strings representing environment
// variables, using the same format as os.Environ()
func (e *Env) Apply(environ []string) []string {
	allow := map[string]struct{}{}
	for _, key := range e.Allow {
		allow[key] = struct{}{}
	}

	deny := map[string]struct{}{}
	for _, key := range e.Deny {
		deny[key] = struct{}{}
	}

	set := map[string]string{}
	for key, val := range e.Set {
		set[key] = key + "=" + val
	}

	tmp := make([]string, 0, len(environ))
	updated := map[string]struct{}{}
	for _, kv := range environ {
		key := kv
		if idx := strings.Index(kv, "="); idx > 0 {
			key = kv[:idx]
		}

		newkv, ok := set[key]
		switch {
		case ok:
			tmp = append(tmp, newkv)
			updated[key] = struct{}{}
		case len(allow) > 0:
			if _, ok := allow[key]; ok {
				tmp = append(tmp, kv)
			}
		case len(deny) > 0:
			if _, ok := deny[key]; !ok {
				tmp = append(tmp, kv)
			}
		default:
			tmp = append(tmp, kv)
		}
	}

	for key, kv := range set {
		if _, ok := updated[key]; !ok {
			tmp = append(tmp, kv)
		}
	}
	return tmp
}
