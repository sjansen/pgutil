package demo

import "github.com/sjansen/pgutil/internal/taskset/base"

var _ execer = &Rev{}

// Rev reverses the target's string when executed
type Rev struct {
	base.Task
}

// Ready validates the task's settings
func (x *Rev) Ready() error {
	return nil
}

func (x *Rev) exec(t *Target) error {
	t.Lock()
	runes := []rune(t.String)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	t.String = string(runes)
	t.Unlock()
	return nil
}
