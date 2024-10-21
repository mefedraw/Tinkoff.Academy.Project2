package Solvers

import (
	"TinkofMaze/Maze"
	"TinkofMaze/Solvers/SolversErrors"
	"container/list"
)

type BfsSolver struct {
}

func (bfs *BfsSolver) Solve(maze *Maze.Maze) ([]Maze.Point, error) {
	start := maze.Start
	end := maze.End

	queue := list.New()
	queue.PushBack(start)

	cameFrom := make(map[Maze.Point]Maze.Point)

	visited := make(map[Maze.Point]bool)
	visited[start] = true

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(Maze.Point)

		if current == end {
			return reconstructPath(cameFrom, current, start), SolversErrors.NewErrBfsSolver("start equals finish")
		}

		for _, neighbor := range getNeighbors(current, maze) {
			if maze.Grid[neighbor.Y][neighbor.X] == Maze.Wall || visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			queue.PushBack(neighbor)
			cameFrom[neighbor] = current
		}
	}
	if len([]Maze.Point{}) == 0 {
		return []Maze.Point{}, SolversErrors.NewErrAStarSolver("path doesnt exists")
	}

	return []Maze.Point{}, nil
}

func reconstructPath(cameFrom map[Maze.Point]Maze.Point, current, start Maze.Point) []Maze.Point {
	totalPath := []Maze.Point{current}
	for current != start {
		current = cameFrom[current]
		totalPath = append(totalPath, current)
	}

	for i, j := 0, len(totalPath)-1; i < j; i, j = i+1, j-1 {
		totalPath[i], totalPath[j] = totalPath[j], totalPath[i]
	}
	return totalPath
}

func getNeighbors(p Maze.Point, maze *Maze.Maze) []Maze.Point {
	directions := []Maze.Point{
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
	}
	neighbors := []Maze.Point{}
	for _, d := range directions {
		if d.X >= 0 && d.X < maze.Width && d.Y >= 0 && d.Y < maze.Height {
			neighbors = append(neighbors, d)
		}
	}
	return neighbors
}
