package Generators

import (
	"TinkofMaze/Generators/GeneratorsErrors"
	"TinkofMaze/Maze"
	"math/rand"
)

type PrimsGenerator struct {
	Saturation int
}

func NewPrimsGenerator(saturation int) *PrimsGenerator {
	return &PrimsGenerator{
		Saturation: saturation,
	}
}

func (pg *PrimsGenerator) Generate(height, width int, start, end Maze.Point) (Maze.Maze, error) {

	maze := Maze.NewMaze(height, width, start, end)

	maze.Grid[start.Y][start.X] = Maze.Passage
	maze.Grid[end.Y][end.X] = Maze.Passage

	pg.GeneratePrim(maze, start, end)

	err := pg.ApplySaturation(maze, pg.Saturation)
	if err != nil {
		return *maze, err
	}

	return *maze, nil
}

type CellWall struct {
	Point Maze.Point
	From  Maze.Point
}

func (pg *PrimsGenerator) GeneratePrim(m *Maze.Maze, start, end Maze.Point) {
	var walls []CellWall

	addWalls := func(p Maze.Point) {
		directions := []Maze.Point{
			{X: p.X, Y: p.Y - 2},
			{X: p.X, Y: p.Y + 2},
			{X: p.X - 2, Y: p.Y},
			{X: p.X + 2, Y: p.Y},
		}
		for _, d := range directions {
			if d.Y >= 0 && d.Y < m.Height && d.X >= 0 && d.X < m.Width && m.Grid[d.Y][d.X] == Maze.Wall {
				walls = append(walls, CellWall{Point: d, From: p})
			}
		}
	}

	addWalls(start)

	for len(walls) > 0 {
		idx := rand.Intn(len(walls))
		wall := walls[idx]

		opposite := Maze.Point{
			X: wall.Point.X + (wall.Point.X - wall.From.X),
			Y: wall.Point.Y + (wall.Point.Y - wall.From.Y),
		}

		if opposite.Y >= 0 && opposite.Y < m.Height && opposite.X >= 0 && opposite.X < m.Width {
			if m.Grid[opposite.Y][opposite.X] == Maze.Wall {
				mazeWall := Maze.Point{
					X: (wall.Point.X + wall.From.X) / 2,
					Y: (wall.Point.Y + wall.From.Y) / 2,
				}
				m.Grid[mazeWall.Y][mazeWall.X] = Maze.Passage
				m.Grid[opposite.Y][opposite.X] = Maze.Passage

				addWalls(opposite)
			}
		}

		walls = append(walls[:idx], walls[idx+1:]...)
	}
}

func (pg *PrimsGenerator) ApplySaturation(m *Maze.Maze, saturation int) error {
	if saturation <= 0 {
		return GeneratorsErrors.NewErrPrimsGenerator("saturation below zero")
	}

	for y := 1; y < m.Height-1; y++ {
		for x := 1; x < m.Width-1; x++ {
			if m.Grid[y][x] == Maze.Wall {
				if rand.Intn(100) < pg.Saturation {
					m.Grid[y][x] = Maze.Passage
				}
			}
		}
	}
	return nil
}
