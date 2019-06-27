package ddl

import "strings"

func quoteName(s string) string {
	return `"` + strings.Replace(s, `"`, `""`, -1) + `"`
}
