package xtream_codes_go

import "errors"

type Error struct {
	message string
	prev    error
}

func (e *Error) Error() string {
	var err = e.prev
	var msg = e.message

	for {
		msg += ", " + err.Error()

		if prev := errors.Unwrap(err); prev != nil {
			err = prev
		} else {
			break
		}
	}

	return msg
}

func (e *Error) Unwrap() error {
	return e.prev
}
