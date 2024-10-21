package SolversErrors

import "fmt"

type ErrAStarSolver struct {
	msg string
}

func NewErrAStarSolver(msg string) *ErrAStarSolver {
	return &ErrAStarSolver{msg: msg}
}

func (e *ErrAStarSolver) Error() string {
	return fmt.Sprintf("AstarSolver error: %s", e.msg)
}
