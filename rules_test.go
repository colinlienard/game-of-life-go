package main

import (
	"slices"
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestCellsAround(t *testing.T) {
	grid := []rl.Vector2{
		{X: 5, Y: 7},
		{X: 4, Y: 5},
		{X: 2, Y: 5},
		{X: 6, Y: 4},
		{X: 5, Y: 6},
		{X: 6, Y: 3},
	}
	expectedResult := []rl.Vector2{
		{X: 4, Y: 5},
		{X: 6, Y: 4},
		{X: 5, Y: 6},
	}
	result := getCellsAround(grid, rl.Vector2{X: 5, Y: 5})

	for _, item := range expectedResult {
		contains := slices.ContainsFunc(result, func(e rl.Vector2) bool {
			return e.X == item.X && e.Y == item.Y
		})
		if !contains {
			t.Errorf("Missing %+v", item)
		}
	}
}

func TestFillAround(t *testing.T) {
	grid := []rl.Vector2{
		{X: 2, Y: 2},
		{X: 3, Y: 3},
		{X: 5, Y: 6},
	}
	result := fillAroundCells(grid)
	expectedResult := []rl.Vector2{
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 3},
		{X: 2, Y: 1},
		{X: 2, Y: 3},
		{X: 3, Y: 1},
		{X: 3, Y: 2},

		{X: 2, Y: 4},
		{X: 3, Y: 4},
		{X: 4, Y: 2},
		{X: 4, Y: 3},
		{X: 4, Y: 4},

		{X: 4, Y: 5},
		{X: 4, Y: 6},
		{X: 4, Y: 7},
		{X: 5, Y: 5},
		{X: 5, Y: 7},
		{X: 6, Y: 5},
		{X: 6, Y: 6},
		{X: 6, Y: 7},
	}

	if len(result) != len(expectedResult) {
		t.Errorf("Got length %d, expected %d", len(result), len(expectedResult))
	}
	for _, item := range expectedResult {
		contains := slices.ContainsFunc(result, func(e rl.Vector2) bool {
			return e.X == item.X && e.Y == item.Y
		})
		if !contains {
			t.Errorf("Missing %+v", item)
		}
	}
}

func TestGetLifeStatus(t *testing.T) {
	tests := []struct {
		name     string
		grid     []rl.Vector2
		cell     rl.Vector2
		alive    bool
		expected bool
	}{
		// Live cell tests
		{
			name: "Live cell with 0 neighbors dies (underpopulation)",
			grid: []rl.Vector2{
				{X: 5, Y: 5},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: false,
		},
		{
			name: "Live cell with 1 neighbor dies (underpopulation)",
			grid: []rl.Vector2{
				{X: 5, Y: 5},
				{X: 5, Y: 6},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: false,
		},
		{
			name: "Live cell with 2 neighbors survives",
			grid: []rl.Vector2{
				{X: 5, Y: 5},
				{X: 5, Y: 6},
				{X: 6, Y: 5},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: true,
		},
		{
			name: "Live cell with 3 neighbors survives",
			grid: []rl.Vector2{
				{X: 5, Y: 5},
				{X: 5, Y: 6},
				{X: 6, Y: 5},
				{X: 6, Y: 6},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: true,
		},
		{
			name: "Live cell with 4 neighbors dies (overpopulation)",
			grid: []rl.Vector2{
				{X: 5, Y: 5},
				{X: 5, Y: 6},
				{X: 6, Y: 5},
				{X: 6, Y: 6},
				{X: 4, Y: 5},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: false,
		},
		{
			name: "Live cell with 5 neighbors dies (overpopulation)",
			grid: []rl.Vector2{
				{X: 5, Y: 5},
				{X: 5, Y: 6},
				{X: 6, Y: 5},
				{X: 6, Y: 6},
				{X: 4, Y: 5},
				{X: 4, Y: 6},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: false,
		},
		// Dead cell tests
		{
			name:     "Dead cell with 0 neighbors stays dead",
			grid:     []rl.Vector2{},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		{
			name: "Dead cell with 1 neighbor stays dead",
			grid: []rl.Vector2{
				{X: 5, Y: 6},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		{
			name: "Dead cell with 2 neighbors stays dead",
			grid: []rl.Vector2{
				{X: 5, Y: 6},
				{X: 6, Y: 5},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		{
			name: "Dead cell with 3 neighbors becomes alive (reproduction)",
			grid: []rl.Vector2{
				{X: 5, Y: 6},
				{X: 6, Y: 5},
				{X: 6, Y: 6},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: true,
		},
		{
			name: "Dead cell with 4 neighbors stays dead",
			grid: []rl.Vector2{
				{X: 5, Y: 6},
				{X: 6, Y: 5},
				{X: 6, Y: 6},
				{X: 4, Y: 5},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		{
			name: "Dead cell with 5 neighbors stays dead",
			grid: []rl.Vector2{
				{X: 5, Y: 6},
				{X: 6, Y: 5},
				{X: 6, Y: 6},
				{X: 4, Y: 5},
				{X: 4, Y: 6},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		// Edge cases with diagonal neighbors
		{
			name: "Live cell with 3 diagonal neighbors survives",
			grid: []rl.Vector2{
				{X: 5, Y: 5},
				{X: 4, Y: 4},
				{X: 4, Y: 6},
				{X: 6, Y: 4},
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: true,
		},
		{
			name: "Dead cell with 3 mixed neighbors becomes alive",
			grid: []rl.Vector2{
				{X: 5, Y: 4}, // above
				{X: 6, Y: 6}, // diagonal
				{X: 4, Y: 5}, // left
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getLifeStatus(tt.grid, tt.cell, tt.alive)
			if result != tt.expected {
				t.Errorf("getLifeStatus() = %v, want %v", result, tt.expected)
			}
		})
	}
}
