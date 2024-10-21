package DataStructuresErrors

import "fmt"

type ErrPriorityQueue struct {
	msg string
}

func NewErrPriorityQueue(msg string) *ErrPriorityQueue {
	return &ErrPriorityQueue{msg: msg}
}

func (e *ErrPriorityQueue) Error() string {
	return fmt.Sprintf("PriorityQueue error: %s", e.msg)
}
