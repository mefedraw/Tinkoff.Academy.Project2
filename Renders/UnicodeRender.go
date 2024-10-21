package Renders

import (
	"TinkofMaze/Maze"
	"TinkofMaze/Renders/RendersErrors"
	"TinkofMaze/Solvers"
	"fmt"
	"time"
)

type UnicodeRender struct {
	passageSymbol  string
	wallSymbol     string
	borderSymbol   string
	entranceSymbol string
	exitSymbol     string
	pathSymbol     string
}

func (cr *UnicodeRender) Init() {
	cr.passageSymbol = "⬛"
	cr.wallSymbol = "⬜"
	cr.borderSymbol = "🟫"
	cr.entranceSymbol = "🟨"
	cr.exitSymbol = "🟪"
	cr.pathSymbol = "🟥"
}

func (cr *UnicodeRender) Render(m *Maze.Maze, solver Solvers.Solver) {
	cr.Init()
	err := cr.RenderShowUtil(m)
	if err != nil {
		println(err.Error())
	}
	time.Sleep(2500 * time.Millisecond)
	var path, errSolv = solver.Solve(m)
	if errSolv != nil {
		println(errSolv.Error())
	}
	for i := 1; i < len(path)-1; i++ {
		curPath := path[i]
		m.Grid[curPath.Y][curPath.X] = Maze.Path
		fmt.Print("\x1b[!p")
		err := cr.RenderShowUtil(m)

		if err != nil {
			println(err.Error())
		}
		time.Sleep(150 * time.Millisecond)
	}
}

func (cr *UnicodeRender) RenderShowUtil(m *Maze.Maze) error {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			var curMazeElem = m.Grid[y][x]
			switch {
			case curMazeElem == Maze.Wall:
				fmt.Print(cr.wallSymbol)
			case curMazeElem == Maze.Path:
				fmt.Print(cr.pathSymbol)
			case y == m.End.Y && x == m.End.X:
				fmt.Print(cr.exitSymbol)
			case y == m.Start.Y && x == m.Start.X:
				fmt.Print(cr.entranceSymbol)
			case curMazeElem == Maze.Passage:
				fmt.Print(cr.passageSymbol)
			default:
				return RendersErrors.NewErrRender("unknown maze Cell entity")
			}
		}
		fmt.Println()
	}
	return nil
}
