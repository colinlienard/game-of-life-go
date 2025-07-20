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
		rl.BeginDrawing()

		rl.BeginMode2D(game.camera)

		rl.ClearBackground(rl.Black)

		game.updateCameraPan()
		game.updateCameraZoom()
		game.updateGameMode()
		game.updateSimulationSpeed()
		game.updateCells()
		game.simulateCells()

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
		rl.DrawText("Pause with 'p'", 8, 8+20, 16, rl.Gray)
		rl.DrawText("Reset with 'r'", 8, 8+20*2, 16, rl.Gray)
		rl.DrawText("Increment speed with 'i'", 8, 8+20*3, 16, rl.Gray)
		rl.DrawText("Decrement speed with 'd'", 8, 8+20*4, 16, rl.Gray)
		rl.DrawText(fmt.Sprintf("Simulation speed: %d/s", 120/game.simulationSpeed), 300, 8, 16, rl.Blue)
	}
	rl.DrawFPS(600, 8)
}
