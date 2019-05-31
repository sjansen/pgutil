package demo

import (
	"errors"
	"fmt"
	"time"
)

var _ munger = &Echo{}

// Echo prints the target's string when executed
type Echo struct{}

func (x *Echo) munge(t *Target) error {
	_, err := fmt.Fprintln(t.stdout, t.String)
	return err
}

func (x *Echo) Check() error {
	return nil
}

var _ munger = &Fail{}

// Fail always returns an error when executed
type Fail struct{}

var ErrFail = errors.New("fail task executed")

func (x *Fail) munge(t *Target) error {
	return ErrFail
}

func (x *Fail) Check() error {
	return nil
}

var _ munger = &Rev{}

// Rev reverses the target's string when executed
type Rev struct{}

func (x *Rev) munge(t *Target) error {
	runes := []rune(t.String)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	t.String = string(runes)
	return nil
}

func (x *Rev) Check() error {
	return nil
}

var _ munger = &Rot13{}

// Rot13 applies the ROT-13 substitution cipher to the target's string when executed
type Rot13 struct{}

func (x *Rot13) munge(t *Target) error {
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
	return nil
}

func (x *Rot13) Check() error {
	return nil
}

var _ munger = &Sleep{}

// Sleep pauses the current target for Seconds seconds when executed
type Sleep struct {
	Seconds int
}

func (x *Sleep) munge(t *Target) error {
	time.Sleep(
		time.Duration(x.Seconds) * time.Second,
	)
	return nil
}

func (x *Sleep) Check() error {
	if x.Seconds < 1 {
		x.Seconds = 1
	}
	return nil
}
