package Maze

type Maze struct {
	Height, Width int
	Start, End    Point
	Grid          [][]Cell
}

func NewMaze(height, width int, start, end Point) *Maze {
	Maze := new(Maze)
	Maze.Height = height
	Maze.Width = width
	Maze.Start = start
	Maze.End = end
	Maze.Grid = make([][]Cell, Maze.Height)
	for i := 0; i < Maze.Height; i++ {
		row := make([]Cell, Maze.Width)
		for j := 0; j < Maze.Width; j++ {
			row[j] = Wall
		}
		Maze.Grid[i] = row
	}

	return Maze
}
