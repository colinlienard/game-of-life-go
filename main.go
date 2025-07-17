package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	X int32
	Y int32
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	cells := []Cell{}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			mousePos := rl.GetMousePosition()
			cells = append(cells, Cell{X: int32(mousePos.X), Y: int32(mousePos.Y)})
		}

		for _, cell := range cells {
			rl.DrawRectangle(cell.X, cell.Y, 20, 20, rl.White)
		}

		rl.EndDrawing()
	}
}
