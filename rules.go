package main

import (
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func getCellsAround(grid []rl.Vector2, cell rl.Vector2) []rl.Vector2 {
	result := []rl.Vector2{}

	for _, c := range grid {
		if c.X >= cell.X-1 && c.X <= cell.X+1 && c.Y >= cell.Y-1 && c.Y <= cell.Y+1 && !(c.X == cell.X && c.Y == cell.Y) {
			result = append(result, c)
		}
	}

	return result
}

func fillAroundCells(grid []rl.Vector2) []rl.Vector2 {
	result := []rl.Vector2{}

	for _, cell := range grid {
		for x := -1; x < 2; x++ {
			for y := -1; y < 2; y++ {
				if !(x == 0 && y == 0) {
					_x := cell.X + float32(x)
					_y := cell.Y + float32(y)
					contains := slices.ContainsFunc(grid, func(e rl.Vector2) bool {
						return e.X == _x && e.Y == _y
					})
					if contains {
						continue
					}
					contains = slices.ContainsFunc(result, func(e rl.Vector2) bool {
						return e.X == _x && e.Y == _y
					})
					if contains {
						continue
					}
					result = append(result, rl.Vector2{X: _x, Y: _y})
				}
			}
		}
	}

	return result
}

func getLifeStatus(grid []rl.Vector2, cell rl.Vector2, alive bool) bool {
	cellsAround := len(getCellsAround(grid, cell))

	switch {
	// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	case alive && cellsAround < 2:
		return false
	// Any live cell with two or three live neighbours lives on to the next generation.
	case alive && (cellsAround == 2 || cellsAround == 3):
		return true
	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	case alive && cellsAround > 3:
		return false
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	case !alive && cellsAround == 3:
		return true
	default:
		return false
	}
}
