package ddl

import "strings"

func quoteName(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
