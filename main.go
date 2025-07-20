package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const CELL_SIZE float32 = 20
const CAMERA_ZOOM_SPEED float32 = 0.05
const SIMULATION_SPEED_THRESHOLD = 5

func main() {
	rl.InitWindow(1280, 720, "Game of Life")
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
	if game.mode == EditGameMode {
		rl.DrawText("Start with 's'", 8, 8, 16, rl.Gray)
		rl.DrawText("Edit", 300, 8, 16, rl.Blue)
	} else {
		rl.DrawText("Stop with 's'", 8, 8, 16, rl.Gray)
		rl.DrawText("Reset with 'r'", 8, 8+20, 16, rl.Gray)
		rl.DrawText("Increment speed with 'i'", 8, 8+20*2, 16, rl.Gray)
		rl.DrawText("Decrement speed with 'd'", 8, 8+20*3, 16, rl.Gray)
		rl.DrawText(fmt.Sprintf("Simulation (%dfps)", rl.GetFPS()/int32(game.simulationSpeed)), 300, 8, 16, rl.Blue)
	}
}
