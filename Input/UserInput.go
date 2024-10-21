package Input

import (
	"TinkofMaze/Generators"
	"TinkofMaze/Input/UserInputErrors"
	"TinkofMaze/Solvers"
	"fmt"
)

type UserInput struct {
	Height, Width, EntranceX, EntranceY, ExitX, ExitY, Saturation int
	GeneratorType                                                 Generators.GeneratorType
	SolverType                                                    Solvers.SolverType
}

func NewUserInput() (*UserInput, error) {
	var ui = UserInput{}
	err := ui.getAllData()
	if err != nil {
		return nil, err
	}
	return &ui, nil
}

func (a *UserInput) getHeight() error {
	fmt.Println("Введите высоту лабиринта:")
	_, err := fmt.Scan(&a.Height)
	if err != nil {
		return UserInputErrors.NewErrUserInput("invalid input for height")
	}
	if a.Height <= 0 {
		return UserInputErrors.NewErrUserInput("height isn't positive")
	}
	fmt.Print("\x1b[!p")
	return nil
}

func (a *UserInput) getWidth() error {
	fmt.Println("Введите ширину лабиринта:")
	_, err := fmt.Scan(&a.Width)
	if err != nil {
		return UserInputErrors.NewErrUserInput("invalid input for width")
	}
	if a.Width <= 0 {
		return UserInputErrors.NewErrUserInput("width isn't positive")
	}
	fmt.Print("\x1b[!p")
	return nil
}

func (a *UserInput) getEntrance() error {
	fmt.Println("Введите координаты входа (x, y):")
	_, err := fmt.Scan(&a.EntranceX, &a.EntranceY)
	if err != nil {
		return UserInputErrors.NewErrUserInput("invalid input for entrance coordinates")
	}
	if a.EntranceX < 0 || a.EntranceY < 0 || a.EntranceX >= a.Width || a.EntranceY >= a.Height {
		return UserInputErrors.NewErrUserInput("entrance coordinates out of bounds")
	}
	fmt.Print("\x1b[!p")
	return nil
}

func (a *UserInput) getExit() error {
	fmt.Println("Введите координаты выхода (x, y):")
	_, err := fmt.Scan(&a.ExitX, &a.ExitY)
	if err != nil {
		return UserInputErrors.NewErrUserInput("invalid input for exit coordinates")
	}
	if a.ExitX < 0 || a.ExitY < 0 || a.ExitX >= a.Width || a.ExitY >= a.Height {
		return UserInputErrors.NewErrUserInput("exit coordinates out of bounds")
	}
	fmt.Print("\x1b[!p")
	return nil
}

func (a *UserInput) getSaturation() error {
	fmt.Println("Введите насыщенность (например, 70 для 70%):")
	_, err := fmt.Scan(&a.Saturation)
	if err != nil {
		return UserInputErrors.NewErrUserInput("invalid input for saturation")
	}
	if a.Saturation < 0 || a.Saturation > 100 {
		return UserInputErrors.NewErrUserInput("saturation must be between 0 and 100")
	}
	fmt.Print("\x1b[!p")
	return nil
}

func (a *UserInput) getGeneratorType() error {
	fmt.Println("Введите тип генератора ('t для tree' или 'p для prim'):")
	var genType = ""
	_, err := fmt.Scan(&genType)
	if err != nil {
		return UserInputErrors.NewErrUserInput("invalid input for generator type")
	}
	if genType != "t" && genType != "p" {
		return UserInputErrors.NewErrUserInput("generator type must be 't' or 'p'")
	}
	if genType == "t" {
		a.GeneratorType = Generators.GrowingTree
	} else {
		a.GeneratorType = Generators.PRIM
	}

	fmt.Print("\x1b[!p")
	return nil
}

func (a *UserInput) getSolverType() error {
	fmt.Println("Введите тип решателя (например, 'b для bfs' или 'a для astar'):")
	var solverType = ""
	_, err := fmt.Scan(&solverType)
	if err != nil {
		return UserInputErrors.NewErrUserInput("invalid input for solver type")
	}
	if solverType != "b" && solverType != "a" {
		return UserInputErrors.NewErrUserInput("solver type must be 'b' or 'a'")
	}

	if solverType == "b" {
		a.SolverType = Solvers.BFS
	} else {
		a.SolverType = Solvers.Astar
	}
	fmt.Print("\x1b[!p")
	return nil
}

func (a *UserInput) getAllData() error {
	if err := a.getHeight(); err != nil {
		return err
	}
	if err := a.getWidth(); err != nil {
		return err
	}
	if err := a.getEntrance(); err != nil {
		return err
	}
	if err := a.getExit(); err != nil {
		return err
	}
	if err := a.getSaturation(); err != nil {
		return err
	}
	if err := a.getGeneratorType(); err != nil {
		return err
	}
	if err := a.getSolverType(); err != nil {
		return err
	}
	return nil
}
