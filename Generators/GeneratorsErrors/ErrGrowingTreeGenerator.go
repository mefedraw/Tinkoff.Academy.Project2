package GeneratorsErrors

import "fmt"

type ErrGrowingTreeGenerator struct {
	msg string
}

func NewErrGrowingTreeGenerator(msg string) *ErrEulersGenerator {
	return &ErrEulersGenerator{msg: msg}
}

func (e *ErrGrowingTreeGenerator) Error() string {
	return fmt.Sprintf("GrowingTreeGenerator error: %s", e.msg)
}
