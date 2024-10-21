package Solvers

import (
	"TinkofMaze/DataStructures"
	"TinkofMaze/Maze"
	"TinkofMaze/Solvers/SolversErrors"
	"fmt"
)

type AstarSolver struct {
}

func (as *AstarSolver) Solve(maze *Maze.Maze) ([]Maze.Point, error) {
	start := maze.Start
	end := maze.End

	pq := &DataStructures.PriorityQueue{}
	pq.Push(&DataStructures.Item{
		Point:    start,
		Priority: heuristic(start, end),
	})

	cameFrom := make(map[Maze.Point]Maze.Point)
	gScore := make(map[Maze.Point]float64)
	gScore[start] = 0

	for pq.Len() > 0 {
		currentItem, err := pq.Pop()
		if err != nil {
			fmt.Println(err)
		}
		current := currentItem.Point

		if current == end {
			return reconstructPath(cameFrom, current, start), SolversErrors.NewErrAStarSolver("start equals finish")
		}

		for _, neighbor := range getNeighbors(current, maze) {
			if maze.Grid[neighbor.Y][neighbor.X] == Maze.Wall {
				continue
			}

			tentativeGScore := gScore[current] + distance(current, neighbor)

			if tentativeGScore < gScore[neighbor] || gScore[neighbor] == 0 {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore := tentativeGScore + heuristic(neighbor, end)

				found := false
				for _, item := range *pq {
					if item.Point == neighbor {
						if fScore < item.Priority {
							item.Priority = fScore
							pq.Down(item.Index)
						}
						found = true
						break
					}
				}

				if !found {
					pq.Push(&DataStructures.Item{
						Point:    neighbor,
						Priority: fScore,
					})
				}
			}
		}
	}

	if len([]Maze.Point{}) == 0 {
		return []Maze.Point{}, SolversErrors.NewErrAStarSolver("path doesnt exists")
	}

	return []Maze.Point{}, nil
}

func (as *AstarSolver) reconstructPath(cameFrom map[Maze.Point]Maze.Point, current, start Maze.Point) []Maze.Point {
	path := []Maze.Point{current}
	for current != start {
		current = cameFrom[current]
		path = append([]Maze.Point{current}, path...)
	}
	return path
}

func (as *AstarSolver) getNeighbors(p Maze.Point, maze *Maze.Maze) []Maze.Point {
	neighbors := []Maze.Point{}
	directions := []Maze.Point{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}

	for _, dir := range directions {
		neighbor := Maze.Point{
			X: p.X + dir.X,
			Y: p.Y + dir.Y,
		}

		if neighbor.X >= 0 && neighbor.X < maze.Width && neighbor.Y >= 0 && neighbor.Y < maze.Height {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func heuristic(a, b Maze.Point) float64 {
	return float64(abs(a.X-b.X) + abs(a.Y-b.Y))
}

func distance(a, b Maze.Point) float64 {
	return 1.0
}
