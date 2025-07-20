package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameMode int

const (
	EditGameMode GameMode = iota
	SimulationGameMode
)

type Game struct {
	cells                map[rl.Vector2]bool
	camera               rl.Camera2D
	mode                 GameMode
	simulationSpeed      int
	simulationFrameDelay int
	panStart             rl.Vector2
	cameraPanStart       rl.Vector2
}

func NewGame() *Game {
	return &Game{
		camera:          rl.Camera2D{Zoom: 1},
		mode:            EditGameMode,
		simulationSpeed: 21,
	}
}

func (g *Game) updateCameraZoom() {
	scroll := rl.GetMouseWheelMoveV().Y
	if scroll == 0 {
		return
	}

	centerX := (float32(rl.GetScreenWidth())/2 - camera.Offset.X) / camera.Zoom
	centerY := (float32(rl.GetScreenHeight())/2 - camera.Offset.Y) / camera.Zoom

	if scroll < 0 && g.camera.Zoom-CAMERA_ZOOM_SPEED > 0.1 {
		g.camera.Zoom -= CAMERA_ZOOM_SPEED
	} else {
		g.camera.Zoom += CAMERA_ZOOM_SPEED
	}

	g.camera.Offset.X = float32(rl.GetScreenWidth())/2 - centerX*g.camera.Zoom
	g.camera.Offset.Y = float32(rl.GetScreenHeight())/2 - centerY*g.camera.Zoom
}

func (g *Game) updateCameraPan() {
	mousePos := rl.GetMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseButtonMiddle) || rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		g.panStart.X = mousePos.X
		g.panStart.Y = mousePos.Y
		g.cameraPanStart.X = g.camera.Offset.X
		g.cameraPanStart.Y = g.camera.Offset.Y
	} else if rl.IsMouseButtonDown(rl.MouseButtonMiddle) || rl.IsMouseButtonDown(rl.MouseButtonRight) {
		g.camera.Offset.X = g.cameraPanStart.X - (g.panStart.X - mousePos.X)
		g.camera.Offset.Y = g.cameraPanStart.Y - (g.panStart.Y - mousePos.Y)
	}
}

func (g *Game) updateGameMode() {
	if rl.IsKeyPressed(rl.KeyS) {
		if g.mode == EditGameMode {
			g.mode = SimulationGameMode
		} else {
			g.mode = EditGameMode
		}
	}
}

func (g *Game) updateSimulationSpeed() {
	if g.mode == EditGameMode {
		return
	}
	if rl.IsKeyPressed(rl.KeyI) {
		g.simulationSpeed += SIMULATION_SPEED_THRESHOLD
	} else if rl.IsKeyPressed(rl.KeyD) && g.simulationSpeed-SIMULATION_SPEED_THRESHOLD >= 1 {
		// TODO: use Math.min or something
		g.simulationSpeed -= SIMULATION_SPEED_THRESHOLD
	}
}

func (g *Game) updateCells() {
	if g.mode != EditGameMode || !rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		return
	}

	mousePos := rl.GetMousePosition()
	worldPos := rl.GetScreenToWorld2D(mousePos, camera)
	x := float32(math.Floor(float64(worldPos.X / CELL_SIZE)))
	y := float32(math.Floor(float64(worldPos.Y / CELL_SIZE)))
	newCell := rl.Vector2{X: x, Y: y}

	if _, exists := g.cells[newCell]; exists {
		delete(g.cells, newCell)
	} else {
		g.cells[newCell] = true
	}
}

func (g *Game) simulateCells() {
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
}
