package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH               = 1280
	WINDOW_HEIGHT              = 720
	CELL_SIZE                  = 20
	CAMERA_ZOOM_SPEED          = 0.05
	SIMULATION_SPEED_THRESHOLD = 4
)

func main() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Game of Life")
	rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))
	defer rl.CloseWindow()

	game := NewGame()

	for !rl.WindowShouldClose() {
		game.updateCameraPan()
		game.updateCameraZoom()
		game.updateGameMode()
		game.updateReset()
		game.updateSimulationSpeed()
		game.updateCells()
		game.simulateCells()

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(game.camera)

		for cell := range game.cells {
			rl.DrawRectangle(
				int32(cell.X*CELL_SIZE),
				int32(cell.Y*CELL_SIZE),
				int32(CELL_SIZE),
				int32(CELL_SIZE),
				rl.White,
			)
		}

		rl.EndMode2D()

		drawUI(game)

		rl.EndDrawing()
	}
}

func drawUI(game *Game) {
	rl.DrawText("Start/stop with 's'", 8, 8, 16, rl.Gray)
	rl.DrawText("Reset with 'r'", 8, 8+20, 16, rl.Gray)

	if game.mode == EditGameMode {
		rl.DrawText("Edit", (WINDOW_WIDTH-rl.MeasureText("Edit", 24))/2, 8, 24, rl.Blue)
	} else {
		rl.DrawText("Increment speed with 'i'", 8, 8+20*2, 16, rl.Gray)
		rl.DrawText("Decrement speed with 'd'", 8, 8+20*3, 16, rl.Gray)

		rl.DrawText("Simulation", (WINDOW_WIDTH-rl.MeasureText("Simulation", 24))/2, 8, 24, rl.Green)

		genText := fmt.Sprintf("Generation: %d", game.generationCount)
		rl.DrawText(genText, WINDOW_WIDTH-8-rl.MeasureText(genText, 16), 8, 16, rl.Gray)
		popText := fmt.Sprintf("Population: %d", len(game.cells))
		rl.DrawText(popText, WINDOW_WIDTH-8-rl.MeasureText(popText, 16), 8+20, 16, rl.Gray)
	}
}
