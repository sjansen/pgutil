package strbuf

import (
	"errors"
	"fmt"
	"time"
)

var _ munger = &Echo{}

type Echo struct{}

func (x *Echo) munge(t *Target) error {
	_, err := fmt.Fprintln(t.stdout, t.Data)
	return err
}

func (x *Echo) Check() error {
	return nil
}

var _ munger = &Fail{}

type Fail struct{}

var ErrFail = errors.New("fail task executed")

func (x *Fail) munge(t *Target) error {
	return ErrFail
}

func (x *Fail) Check() error {
	return nil
}

var _ munger = &Rev{}

type Rev struct{}

func (x *Rev) munge(t *Target) error {
	runes := []rune(t.Data)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	t.Data = string(runes)
	return nil
}

func (x *Rev) Check() error {
	return nil
}

var _ munger = &Rot13{}

type Rot13 struct{}

func (x *Rot13) munge(t *Target) error {
	runes := []rune(t.Data)
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
	t.Data = string(runes)
	return nil
}

func (x *Rot13) Check() error {
	return nil
}

var _ munger = &Sleep{}

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
