package pg

// String converts nullable strings to standard strings
func String(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
