package Renders

import (
	"TinkofMaze/Maze"
	"TinkofMaze/Solvers"
)

type Renderer interface {
	Render(m *Maze.Maze, solver Solvers.Solver)
	Init()
}
