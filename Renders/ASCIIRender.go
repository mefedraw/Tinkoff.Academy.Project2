package Renders

import (
	"TinkofMaze/Maze"
	"TinkofMaze/Renders/RendersErrors"
	"TinkofMaze/Solvers"
	"fmt"
	"time"
)

type ASCIIRender struct {
	passageSymbol  string
	wallSymbol     string
	borderSymbol   string
	entranceSymbol string
	exitSymbol     string
	pathSymbol     string
}

func (cr *ASCIIRender) Init() {
	cr.passageSymbol = " "
	cr.wallSymbol = "#"
	cr.borderSymbol = "█"
	cr.entranceSymbol = "\033[34mS\033[0m" // Синий
	cr.exitSymbol = "\033[32mF\033[0m"     // Зеленый
	cr.pathSymbol = "\033[31m*\033[0m"     // Красный
}

func ConsoleClear() {
	fmt.Print("\x1b[!p")
}

func (cr *ASCIIRender) Render(m *Maze.Maze) {
	cr.Init()
	cr.RenderShowUtil(m)
	time.Sleep(2500 * time.Millisecond)
	solver := Solvers.BfsSolver{}
	var path, errSolv = solver.Solve(m)
	if errSolv != nil {
		println(errSolv.Error())
	}
	for i := 1; i < len(path)-1; i++ {
		curPath := path[i]
		m.Grid[curPath.Y][curPath.X] = Maze.Path
		ConsoleClear()
		err := cr.RenderShowUtil(m)
		if err != nil {
			println(err.Error())
		}
		time.Sleep(150 * time.Millisecond)
	}
}

func (cr *ASCIIRender) RenderShowUtil(m *Maze.Maze) error {
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
				return RendersErrors.NewErrRender("unknown maze entity")
			}
		}
		fmt.Println()
	}
	return nil
}
