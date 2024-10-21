package Generators

import (
	"TinkofMaze/Maze"
)

type Generator interface {
	Generate(height, width int, start, end Maze.Point) (Maze.Maze, error)
}
