package DataStructuresErrors

import "fmt"

type ErrUnionFind struct {
	msg string
}

func NewErrUnionFind(msg string) *ErrUnionFind {
	return &ErrUnionFind{msg: msg}
}

func (e *ErrUnionFind) Error() string {
	return fmt.Sprintf("UnionFind error: %s", e.msg)
}
