package SolversErrors

import "fmt"

type ErrBfsSolver struct {
	msg string
}

func NewErrBfsSolver(msg string) *ErrBfsSolver {
	return &ErrBfsSolver{msg: msg}
}

func (e *ErrBfsSolver) Error() string {
	return fmt.Sprintf("AstarSolver error: %s", e.msg)
}
