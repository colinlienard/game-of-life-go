package main

import (
	"fmt"
	"math"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ = fmt.Print

const CELL_SIZE float32 = 20
const CAMERA_ZOOM_SPEED float32 = 0.05

var camera = rl.Camera2D{Zoom: 1}
var panStart = rl.Vector2{}
var cameraPanStart = rl.Vector2{}
var editMode = true
var simulationSpeed = 20
var simulationSpeedTreshold = 0

func main() {
	rl.InitWindow(800, 450, "Game of Life")
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	cells := []rl.Vector2{}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.BeginMode2D(camera)

		rl.ClearBackground(rl.Black)

		handleZoom()

		handlePan()

		handleKeys()

		if editMode {
			cells = handleClick(cells)
		} else if simulationSpeedTreshold < simulationSpeed {
			simulationSpeedTreshold++
		} else {
			simulationSpeedTreshold = 0
			cells = simulateCells(cells)
		}

		for _, cell := range cells {
			rl.DrawRectanglePro(rl.Rectangle{
				X:      cell.X * CELL_SIZE,
				Y:      cell.Y * CELL_SIZE,
				Width:  CELL_SIZE,
				Height: CELL_SIZE,
			}, rl.Vector2{}, 0, rl.White)
		}

		rl.EndMode2D()

		renderUI()

		rl.EndDrawing()
	}
}

func renderUI() {
	if editMode {
		rl.DrawText("Start with 's'", 8, 8, 16, rl.Gray)
		rl.DrawText("Edit", 300, 8, 16, rl.Blue)
	} else {
		rl.DrawText("Stop with 's'", 8, 8, 16, rl.Gray)
		rl.DrawText("Pause with 'p'", 8, 8+20, 16, rl.Gray)
		rl.DrawText("Reset with 'r'", 8, 8+20*2, 16, rl.Gray)
		rl.DrawText("Increment speed with 'i'", 8, 8+20*3, 16, rl.Gray)
		rl.DrawText("Decrement speed with 'd'", 8, 8+20*4, 16, rl.Gray)
		rl.DrawText(fmt.Sprintf("Simulation speed: %d", simulationSpeed), 300, 8, 16, rl.Blue)
	}
}

func handleZoom() {
	scroll := rl.GetMouseWheelMoveV().Y
	if scroll == 0 {
		return
	}

	centerX := (float32(rl.GetScreenWidth())/2 - camera.Offset.X) / camera.Zoom
	centerY := (float32(rl.GetScreenHeight())/2 - camera.Offset.Y) / camera.Zoom

	if scroll < 0 && camera.Zoom-CAMERA_ZOOM_SPEED > 0.1 {
		camera.Zoom -= CAMERA_ZOOM_SPEED
	} else {
		camera.Zoom += CAMERA_ZOOM_SPEED
	}

	camera.Offset.X = float32(rl.GetScreenWidth())/2 - centerX*camera.Zoom
	camera.Offset.Y = float32(rl.GetScreenHeight())/2 - centerY*camera.Zoom
}

func handlePan() {
	mousePos := rl.GetMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseButtonMiddle) || rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		panStart.X = mousePos.X
		panStart.Y = mousePos.Y
		cameraPanStart.X = camera.Offset.X
		cameraPanStart.Y = camera.Offset.Y
	} else if rl.IsMouseButtonDown(rl.MouseButtonMiddle) || rl.IsMouseButtonDown(rl.MouseButtonRight) {
		camera.Offset.X = cameraPanStart.X - (panStart.X - mousePos.X)
		camera.Offset.Y = cameraPanStart.Y - (panStart.Y - mousePos.Y)
	}
}

func handleClick(cells []rl.Vector2) []rl.Vector2 {
	mousePos := rl.GetMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		worldPos := rl.GetScreenToWorld2D(mousePos, camera)
		x := float32(math.Floor(float64(worldPos.X / CELL_SIZE)))
		y := float32(math.Floor(float64(worldPos.Y / CELL_SIZE)))

		i := slices.IndexFunc(cells, func(c rl.Vector2) bool {
			return c.X == x && c.Y == y
		})
		if i == -1 {
			return append(cells, rl.Vector2{X: x, Y: y})
		}
		return slices.Delete(cells, i, i+1)
	}

	return cells
}

func simulateCells(cells []rl.Vector2) []rl.Vector2 {
	result := []rl.Vector2{}
	deadCells := fillAroundCells(cells)

	for _, deadCell := range deadCells {
		alive := getLifeStatus(cells, deadCell, false)
		if alive {
			result = append(result, deadCell)
		}
	}

	for _, cell := range cells {
		alive := getLifeStatus(cells, cell, true)
		if alive {
			result = append(result, cell)
		}
	}

	return result
}

func handleKeys() {
	if rl.IsKeyPressed(rl.KeyS) {
		editMode = !editMode
	}

	if editMode {
		return
	}
	if rl.IsKeyPressed(rl.KeyI) {
		simulationSpeed++
	} else if rl.IsKeyPressed(rl.KeyD) {
		simulationSpeed--
	}
}
