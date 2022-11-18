package err

import (
	"fmt"
)

var (
	ErrTodo = NewTrackable("Implemention needed (TODO)")
)

type Trackable struct {
	msg   string
	cause error
}

func NewTrackable(msg string, args ...any) *Trackable {
	return &Trackable{
		msg: fmt.Sprintf(msg, args...),
	}
}

func (e *Trackable) Error() string {
	if e.cause == nil {
		return e.msg
	}

	return e.msg + "\n\t" + e.cause.Error()
}

func (e Trackable) Track(msg string, args ...any) *Trackable {
	e.cause = NewTrackable(msg, args...)
	return &e
}

func (e Trackable) Wrap(cause error) *Trackable {
	e.cause = cause
	return &e
}

func (e Trackable) WrapTrack(cause error, msg string, args ...any) *Trackable {
	e.cause = NewTrackable(msg, args...).Wrap(cause)
	return &e
}

func (e Trackable) Unwrap() error {
	return e.cause
}
