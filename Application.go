package main

import (
	"TinkofMaze/Generators"
	"TinkofMaze/Input"
	"TinkofMaze/Maze"
	"TinkofMaze/Renders"
	"TinkofMaze/Solvers"
	"fmt"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (a *Application) Run() {
	var userinput, userInputErr = Input.NewUserInput()
	handleError(userInputErr)
	var generator, getGenErr = Generators.NewGeneratorFactory().GetGenerator(userinput.GeneratorType)
	handleError(getGenErr)
	var maze, genErr = generator.Generate(userinput.Height, userinput.Width, Maze.Point{userinput.EntranceX, userinput.EntranceY}, Maze.Point{userinput.ExitX, userinput.ExitY})
	handleError(genErr)
	var solver, getSolverErr = Solvers.NewSolverFactoryFactory().GetSolver(userinput.SolverType)
	handleError(getSolverErr)
	var render = Renders.UnicodeRender{}
	render.Init()
	render.Render(&maze, solver)
}
