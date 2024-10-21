package GeneratorsErrors

import "fmt"

type ErrEulersGenerator struct {
	msg string
}

func NewErrEulersGenerator(msg string) *ErrEulersGenerator {
	return &ErrEulersGenerator{msg: msg}
}

func (e *ErrEulersGenerator) Error() string {
	return fmt.Sprintf("EulersGenerator error: %s", e.msg)
}
