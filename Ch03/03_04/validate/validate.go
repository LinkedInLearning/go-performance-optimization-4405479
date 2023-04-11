package validate

import (
	"errors"
	"time"
)

type Event struct {
	Time   time.Time
	Login  string
	Action string
	Path   string
}

var (
	initialTime = time.Date(2022, time.July, 4, 11, 22, 33, 44, time.UTC)
	ErrNoLogin  = errors.New("missing login")
	ErrNoAction = errors.New("missing action")
	ErrBadTime  = errors.New("bad time")
)

func (e Event) Validate() error {
	if e.Login == "" {
		return ErrNoLogin
	}

	if e.Action == "" {
		return ErrNoAction
	}

	if e.Time.Before(initialTime) || e.Time.After(time.Now()) {
		return ErrBadTime
	}

	return nil
}

func Validate(events []Event) []bool {
	var out []bool
	for _, e := range events {
		ok := e.Validate() == nil
		out = append(out, ok)
	}

	return out
}
