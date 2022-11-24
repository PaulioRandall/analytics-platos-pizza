package err

import (
	"fmt"
)

var (
	ErrTodo = Track("Implemention needed (TODO)")
)

type trackable struct {
	msg   string
	cause error
}

func Track(msg string, args ...any) *trackable {
	return &trackable{
		msg: fmt.Sprintf(msg, args...),
	}
}

func Wrap(cause error, msg string, args ...any) *trackable {
	return &trackable{
		msg:   fmt.Sprintf(msg, args...),
		cause: cause,
	}
}

func (e *trackable) Error() string {
	if e.cause == nil {
		return e.msg
	}

	return e.msg + "\n\t" + e.cause.Error()
}

func (e trackable) Unwrap() error {
	return e.cause
}

func (e trackable) Is(other error) bool {
	if oe, ok := other.(*trackable); ok {
		return e.msg == oe.msg
	}

	return false
}

func (e *trackable) Wrap(causes ...error) *trackable {
	var parent *trackable = &(*e)

	for _, c := range causes {
		child, ok := c.(*trackable)
		if !ok {
			child = Wrap(c, "[Trackable wrapper]")
		}

		parent.cause = child
		parent = child
	}

	return e
}

func (e trackable) Trace(msg string, args ...any) *trackable {
	e.cause = Track(msg, args...)
	return &e
}

func (e trackable) TraceWrap(cause error, msg string, args ...any) *trackable {
	e.cause = Wrap(cause, msg, args...)
	return &e
}
