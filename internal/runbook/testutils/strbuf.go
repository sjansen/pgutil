package testutils

import (
	"errors"
	"fmt"
)

type StrBuf struct {
	Message string
}

func (q *StrBuf) ConcurrencyLimit() int {
	return 1
}

func (q *StrBuf) Start(task interface{}) error {
	if t, ok := task.(StrBufTask); ok {
		q.Message = t.Munge(q.Message)
	}
	return errors.New("invalid task")
}

func (q *StrBuf) VerifyConfig() error {
	if q.Message == "" {
		return errors.New("invalid message")
	}
	return nil
}

func (q *StrBuf) VerifyTask(config interface{}) error {
	if _, ok := config.(StrBufTask); !ok {
		return errors.New("invalid task")
	}
	return nil
}

type StrBufTask interface {
	Munge(string) string
}

type EchoTask struct{}

func (t *EchoTask) Munge(msg string) string {
	fmt.Println(msg)
	return msg
}

func (t *EchoTask) VerifyConfig() error {
	return nil
}

type RevTask struct{}

func (t *RevTask) Munge(msg string) string {
	runes := []rune(msg)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (t *RevTask) VerifyConfig() error {
	return nil
}

type Rot13Task struct{}

func (t *Rot13Task) Munge(msg string) string {
	runes := []rune(msg)
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
	return string(runes)
}

func (t *Rot13Task) VerifyConfig() error {
	return nil
}
