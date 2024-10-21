package Generators

import (
	"TinkofMaze/DataStructures"
	"TinkofMaze/Generators/GeneratorsErrors"
	"TinkofMaze/Maze"
	"math/rand"
	"time"
)

type GrowingTreeGenerator struct{}

func NewGrowingTreeGenerator() *GrowingTreeGenerator {
	return &GrowingTreeGenerator{}
}

func (tg *GrowingTreeGenerator) Generate(height, width int, start, end Maze.Point) (Maze.Maze, error) {
	// Проверка входных параметров
	if height <= 0 || width <= 0 {
		return Maze.Maze{}, GeneratorsErrors.NewErrGrowingTreeGenerator("Height and width must be positive integers")
	}

	if !tg.isValidPoint(start, height, width) {
		return Maze.Maze{}, GeneratorsErrors.NewErrGrowingTreeGenerator("Start point is out of maze boundaries")
	}

	if !tg.isValidPoint(end, height, width) {
		return Maze.Maze{}, GeneratorsErrors.NewErrGrowingTreeGenerator("End point is out of maze boundaries")
	}

	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создание новой сетки лабиринта
	maze := Maze.NewMaze(height, width, start, end)

	// Установка начальной и конечной точек
	maze.Grid[start.Y][start.X] = Maze.Passage
	maze.Grid[end.Y][end.X] = Maze.Passage

	// Инициализация структуры UnionFind
	uf := DataStructures.NewUnionFind()

	// Инициализация первого ряда
	for cell := 0; cell < width; cell++ {
		if err := uf.Add(cell); err != nil {
			return Maze.Maze{}, err
		}
	}

	// Списки активных точек
	pointsX := []int{start.X}
	pointsY := []int{start.Y}

	// Основной цикл генерации лабиринта
	for len(pointsX) > 0 {
		// Выбор случайной точки из списка активных
		idx := rand.Intn(len(pointsX))
		x := pointsX[idx]
		y := pointsY[idx]

		// Определение возможных направлений
		directions := rand.Perm(4) // Перемешивание направлений для случайности

		carved := false

		for _, direction := range directions {
			newX, newY := x, y

			switch direction {
			case 0: // Вверх
				newY = y - 2
			case 1: // Вправо
				newX = x + 2
			case 2: // Вниз
				newY = y + 2
			case 3: // Влево
				newX = x - 2
			}

			// Проверка возможности прорезать путь
			if tg.testXY(maze, newY, newX) {
				// Прорезание пути
				maze.Grid[(y+newY)/2][(x+newX)/2] = Maze.Passage
				maze.Grid[newY][newX] = Maze.Passage

				// Добавление новой точки в активный список
				pointsX = append(pointsX, newX)
				pointsY = append(pointsY, newY)

				// Объединение множеств в UnionFind
				currentCell := y*width + x
				newCell := newY*width + newX
				if err := uf.Add(newCell); err != nil {
					return Maze.Maze{}, err
				}
				if err := uf.Union(currentCell, newCell); err != nil {
					return Maze.Maze{}, err
				}

				carved = true
				break
			}
		}

		if !carved {
			// Если не удалось прорезать путь, удаляем точку из активного списка
			pointsX = append(pointsX[:idx], pointsX[idx+1:]...)
			pointsY = append(pointsY[:idx], pointsY[idx+1:]...)
		}
	}

	// Проверка и создание пути до конечной точки, если необходимо
	if maze.Grid[end.Y][end.X] != Maze.Passage {
		if err := tg.createPathToEnd(maze, start, end); err != nil {
			return Maze.Maze{}, err
		}
	}

	return *maze, nil
}

// Проверка, находится ли точка внутри границ лабиринта
func (tg *GrowingTreeGenerator) isValidPoint(p Maze.Point, height, width int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

// Проверка возможности прорезать путь в заданной точке
func (tg *GrowingTreeGenerator) testXY(m *Maze.Maze, y, x int) bool {
	if x < 1 || y < 1 || x >= m.Width-1 || y >= m.Height-1 {
		return false
	}
	if m.Grid[y][x] == Maze.Passage {
		return false
	}
	return true
}

// Метод для создания прямого пути до конечной точки
func (tg *GrowingTreeGenerator) createPathToEnd(m *Maze.Maze, start, end Maze.Point) error {
	current := start

	// Горизонтальное перемещение
	for current.X != end.X {
		if current.X < end.X {
			current.X++
		} else {
			current.X--
		}
		if current.X < 0 || current.X >= m.Width || current.Y < 0 || current.Y >= m.Height {
			return GeneratorsErrors.NewErrGrowingTreeGenerator("Path creation went out of maze boundaries")
		}
		m.Grid[current.Y][current.X] = Maze.Passage
	}

	// Вертикальное перемещение
	for current.Y != end.Y {
		if current.Y < end.Y {
			current.Y++
		} else {
			current.Y--
		}
		if current.X < 0 || current.X >= m.Width || current.Y < 0 || current.Y >= m.Height {
			return GeneratorsErrors.NewErrGrowingTreeGenerator("Path creation went out of maze boundaries")
		}
		m.Grid[current.Y][current.X] = Maze.Passage
	}

	if m.Grid[end.Y][end.X] != Maze.Passage {
		return GeneratorsErrors.NewErrGrowingTreeGenerator("Failed to create path to end")
	}

	return nil
}
