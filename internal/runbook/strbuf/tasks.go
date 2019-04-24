package strbuf

import "fmt"

type EchoTask struct{}

func (t *EchoTask) Munge(msg string) string {
	fmt.Println(msg)
	return msg
}

func (t *EchoTask) Check() error {
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

func (t *RevTask) Check() error {
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

func (t *Rot13Task) Check() error {
	return nil
}
