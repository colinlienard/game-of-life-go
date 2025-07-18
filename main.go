package main

import (
	"fmt"

	r "github.com/gen2brain/raylib-go/raylib"
)

const CELL_SIZE float32 = 20
const CAMERA_ZOOM_SPEED float32 = 0.1

type Cell struct {
	X float32
	Y float32
}

var camera = r.Camera2D{Zoom: 1}
var panStart = r.Vector2{}
var cameraPanStart = r.Vector2{}

func main() {
	r.InitWindow(800, 450, "Game of Life")
	defer r.CloseWindow()

	r.SetTargetFPS(120)

	cells := []Cell{}

	for !r.WindowShouldClose() {
		r.BeginDrawing()

		r.BeginMode2D(camera)

		r.ClearBackground(r.Black)

		mousePos := r.GetMousePosition()

		if r.IsMouseButtonPressed(r.MouseButtonLeft) {
			fmt.Println(mousePos)
			x := float32(int((mousePos.X-camera.Offset.X)/camera.Zoom/CELL_SIZE)) * CELL_SIZE
			y := float32(int((mousePos.Y-camera.Offset.Y)/camera.Zoom/CELL_SIZE)) * CELL_SIZE
			cells = append(cells, Cell{X: x, Y: y})
		}

		fmt.Println(camera.Zoom)
		scroll := r.GetMouseWheelMoveV().Y
		if scroll < 0 && camera.Zoom-CAMERA_ZOOM_SPEED > 0.1 {
			camera.Zoom -= CAMERA_ZOOM_SPEED
			// camera.Offset.X = float32(r.GetScreenWidth()) * CAMERA_ZOOM_SPEED
			// camera.Offset.Y = float32(r.GetScreenHeight()) * CAMERA_ZOOM_SPEED
		} else if scroll > 0 {
			camera.Zoom += CAMERA_ZOOM_SPEED
		}

		if r.IsMouseButtonPressed(r.MouseButtonMiddle) {
			panStart.X = mousePos.X
			panStart.Y = mousePos.Y
			cameraPanStart.X = camera.Offset.X
			cameraPanStart.Y = camera.Offset.Y
		} else if r.IsMouseButtonDown(r.MouseButtonMiddle) {
			camera.Offset.X = cameraPanStart.X - (panStart.X - mousePos.X)
			camera.Offset.Y = cameraPanStart.Y - (panStart.Y - mousePos.Y)
		}

		for _, cell := range cells {
			r.DrawRectanglePro(r.Rectangle{X: cell.X, Y: cell.Y, Width: CELL_SIZE, Height: CELL_SIZE}, r.Vector2{}, 0, r.White)
		}

		r.EndMode2D()

		r.EndDrawing()
	}
}
