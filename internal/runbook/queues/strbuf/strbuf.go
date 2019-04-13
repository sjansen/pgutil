package strbuf

import (
	"errors"
)

type StrBuf struct {
	Message string
}

func (q *StrBuf) ConcurrencyLimit() int {
	return 1
}

func (q *StrBuf) Start(task interface{}) error {
	if t, ok := task.(Task); ok {
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
	if _, ok := config.(Task); !ok {
		return errors.New("invalid task")
	}
	return nil
}
