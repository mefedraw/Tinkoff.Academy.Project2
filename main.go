package main

func main() {

	NewApplication().Run()

	//// Установка seed для генератора случайных чисел
	//rand.Seed(time.Now().UnixNano())
	//
	//// Задаем размеры лабиринта
	//height := 20
	//width := 30
	//
	//// Случайная генерация начальной точки
	//startX := rand.Intn(width-2) + 1  // Диапазон от 1 до width-2
	//startY := rand.Intn(height-2) + 1 // Диапазон от 1 до height-2
	//start := Maze.Point{X: startX, Y: startY}
	//
	//// Случайная генерация конечной точки
	//endX := rand.Intn(width-2) + 1  // Диапазон от 1 до width-2
	//endY := rand.Intn(height-2) + 1 // Диапазон от 1 до height-2
	//end := Maze.Point{X: endX, Y: endY}
	//
	//// Инициализация генератора и генерация лабиринта
	//generator := Generators.GrowingTreeGenerator{}
	//var maze, _ = generator.Generate(height, width, start, end)
	//
	//// Инициализация рендера и отображение лабиринта
	//var render = Renders.UnicodeRender{}
	//render.Init()
	//render.Render(&maze)
}
