package Solvers

import "TinkofMaze/Maze"

type Solver interface {
	Solve(maze *Maze.Maze) ([]Maze.Point, error)
}
