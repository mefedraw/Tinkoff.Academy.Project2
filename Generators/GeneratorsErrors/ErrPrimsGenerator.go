package GeneratorsErrors

import "fmt"

type ErrPrimsGenerator struct {
	msg string
}

func NewErrPrimsGenerator(msg string) *ErrEulersGenerator {
	return &ErrEulersGenerator{msg: msg}
}

func (e *ErrPrimsGenerator) Error() string {
	return fmt.Sprintf("PrimsGenerator error: %s", e.msg)
}
