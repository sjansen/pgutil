package demo

import "github.com/sjansen/pgutil/internal/taskset/base"

var _ execer = &Rot13{}

// Rot13 applies the ROT-13 substitution cipher to the target's string when executed
type Rot13 struct {
	base.Task
}

// Ready validates the task's settings
func (x *Rot13) Ready() error {
	return nil
}

func (x *Rot13) exec(t *Target) error {
	t.Lock()
	runes := []rune(t.String)
	for i, r := range runes {
		switch {
		case 'a' <= r && r <= 'm':
			runes[i] = r + 13
		case 'n' <= r && r <= 'z':
			runes[i] = r - 13
		case 'A' <= r && r <= 'M':
			runes[i] = r + 13
		case 'N' <= r && r <= 'Z':
			runes[i] = r - 13
		}
	}
	t.String = string(runes)
	t.Unlock()
	return nil
}
