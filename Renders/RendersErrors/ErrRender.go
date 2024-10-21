package RendersErrors

import "fmt"

type ErrRender struct {
	msg string
}

func NewErrRender(msg string) *ErrRender {
	return &ErrRender{msg: msg}
}

func (e *ErrRender) Error() string {
	return fmt.Sprintf("Render error: %s", e.msg)
}
