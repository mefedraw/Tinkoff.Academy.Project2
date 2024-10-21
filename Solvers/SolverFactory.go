package Solvers

import (
	"fmt"
)

type SolverFactory struct{}

func NewSolverFactoryFactory() *SolverFactory {
	return &SolverFactory{}
}

func (f *SolverFactory) GetSolver(solverType SolverType) (Solver, error) {
	switch solverType {
	case BFS:
		return &BfsSolver{}, nil
	case Astar:
		return &AstarSolver{}, nil
	default:
		return nil, fmt.Errorf("unknown solver type: %d", solverType)
	}
}
