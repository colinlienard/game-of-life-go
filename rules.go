package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func getNeighborsCount(cells Cells, cell rl.Vector2) int {
	result := 0

	iterateNeighbors(func(x int, y int) {
		newCell := rl.Vector2{X: cell.X + float32(x), Y: cell.Y + float32(y)}
		if cells[newCell] {
			result++
		}
	})

	return result
}

func fillNeighbors(cells Cells) Cells {
	result := Cells{}

	for cell := range cells {
		iterateNeighbors(func(x int, y int) {
			newCell := rl.Vector2{X: cell.X + float32(x), Y: cell.Y + float32(y)}
			if !cells[newCell] && !result[newCell] {
				result[newCell] = true
			}
		})
	}

	return result
}

func iterateNeighbors(fn func(x int, y int)) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			fn(x, y)
		}
	}
}

func getLifeStatus(cells Cells, cell rl.Vector2, alive bool) bool {
	neighborsCount := getNeighborsCount(cells, cell)

	switch {
	// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	case alive && neighborsCount < 2:
		return false
	// Any live cell with two or three live neighbours lives on to the next generation.
	case alive && (neighborsCount == 2 || neighborsCount == 3):
		return true
	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	case alive && neighborsCount > 3:
		return false
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	case !alive && neighborsCount == 3:
		return true
	default:
		return false
	}
}
