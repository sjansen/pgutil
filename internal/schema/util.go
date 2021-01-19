package schema

import "strings"

func quoteName(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
