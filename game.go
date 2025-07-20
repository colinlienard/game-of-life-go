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

type Cells map[rl.Vector2]bool

type Game struct {
	cells                Cells
	camera               rl.Camera2D
	mode                 GameMode
	simulationSpeed      int
	simulationFrameDelay int
	panStart             rl.Vector2
	cameraPanStart       rl.Vector2
}

func NewGame() *Game {
	return &Game{
		cells:           Cells{},
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

	centerX := (float32(rl.GetScreenWidth())/2 - g.camera.Offset.X) / g.camera.Zoom
	centerY := (float32(rl.GetScreenHeight())/2 - g.camera.Offset.Y) / g.camera.Zoom

	if scroll < 0 {
		g.camera.Zoom -= CAMERA_ZOOM_SPEED
	} else {
		g.camera.Zoom += CAMERA_ZOOM_SPEED
	}
	g.camera.Zoom = max(g.camera.Zoom, 0.1)

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

func (g *Game) updateReset() {
	if rl.IsKeyPressed(rl.KeyR) {
		g.cells = Cells{}
		g.camera = rl.Camera2D{Zoom: 1}
	}
}

func (g *Game) updateSimulationSpeed() {
	if g.mode == EditGameMode {
		return
	}
	if rl.IsKeyPressed(rl.KeyD) {
		g.simulationSpeed += SIMULATION_SPEED_THRESHOLD
	} else if rl.IsKeyPressed(rl.KeyI) {
		g.simulationSpeed -= SIMULATION_SPEED_THRESHOLD
	}
	g.simulationSpeed = max(g.simulationSpeed, 1)
}

func (g *Game) updateCells() {
	if g.mode != EditGameMode || !rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		return
	}

	mousePos := rl.GetMousePosition()
	worldPos := rl.GetScreenToWorld2D(mousePos, g.camera)
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
	if g.mode != SimulationGameMode {
		return
	}

	if g.simulationFrameDelay < g.simulationSpeed {
		g.simulationFrameDelay++
		return
	}
	g.simulationFrameDelay = 0

	result := Cells{}

	for cell := range g.cells {
		alive := getLifeStatus(g.cells, cell, true)
		if alive {
			result[cell] = true
		}
	}

	deadCells := fillNeighbors(g.cells)
	for deadCell := range deadCells {
		alive := getLifeStatus(g.cells, deadCell, false)
		if alive {
			result[deadCell] = true
		}
	}

	g.cells = result
}
