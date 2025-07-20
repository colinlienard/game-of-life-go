package main

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestNeighborsCount(t *testing.T) {
	cells := Cells{
		{X: 5, Y: 7}: true,
		{X: 4, Y: 5}: true,
		{X: 2, Y: 5}: true,
		{X: 6, Y: 4}: true,
		{X: 5, Y: 6}: true,
		{X: 6, Y: 3}: true,
	}
	expectedResult := Cells{
		{X: 4, Y: 5}: true,
		{X: 6, Y: 4}: true,
		{X: 5, Y: 6}: true,
	}
	result := getNeighborsCount(cells, rl.Vector2{X: 5, Y: 5})

	if result != len(expectedResult) {
		t.Errorf("Got length %d, expected %d", result, len(expectedResult))
	}
}

func TestFillNeighbors(t *testing.T) {
	cells := Cells{
		{X: 2, Y: 2}: true,
		{X: 3, Y: 3}: true,
		{X: 5, Y: 6}: true,
	}
	result := fillNeighbors(cells)
	expectedResult := Cells{
		{X: 1, Y: 1}: true,
		{X: 1, Y: 2}: true,
		{X: 1, Y: 3}: true,
		{X: 2, Y: 1}: true,
		{X: 2, Y: 3}: true,
		{X: 3, Y: 1}: true,
		{X: 3, Y: 2}: true,

		{X: 2, Y: 4}: true,
		{X: 3, Y: 4}: true,
		{X: 4, Y: 2}: true,
		{X: 4, Y: 3}: true,
		{X: 4, Y: 4}: true,

		{X: 4, Y: 5}: true,
		{X: 4, Y: 6}: true,
		{X: 4, Y: 7}: true,
		{X: 5, Y: 5}: true,
		{X: 5, Y: 7}: true,
		{X: 6, Y: 5}: true,
		{X: 6, Y: 6}: true,
		{X: 6, Y: 7}: true,
	}

	if len(result) != len(expectedResult) {
		t.Errorf("Got length %d, expected %d", len(result), len(expectedResult))
	}
	for item := range expectedResult {
		if !result[item] {
			t.Errorf("Missing %+v", item)
		}
	}
}

func TestGetLifeStatus(t *testing.T) {
	tests := []struct {
		name     string
		cells    Cells
		cell     rl.Vector2
		alive    bool
		expected bool
	}{
		// Live cell tests
		{
			name: "Live cell with 0 neighbors dies (underpopulation)",
			cells: Cells{
				{X: 5, Y: 5}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: false,
		},
		{
			name: "Live cell with 1 neighbor dies (underpopulation)",
			cells: Cells{
				{X: 5, Y: 5}: true,
				{X: 5, Y: 6}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: false,
		},
		{
			name: "Live cell with 2 neighbors survives",
			cells: Cells{
				{X: 5, Y: 5}: true,
				{X: 5, Y: 6}: true,
				{X: 6, Y: 5}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: true,
		},
		{
			name: "Live cell with 3 neighbors survives",
			cells: Cells{
				{X: 5, Y: 5}: true,
				{X: 5, Y: 6}: true,
				{X: 6, Y: 5}: true,
				{X: 6, Y: 6}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: true,
		},
		{
			name: "Live cell with 4 neighbors dies (overpopulation)",
			cells: Cells{
				{X: 5, Y: 5}: true,
				{X: 5, Y: 6}: true,
				{X: 6, Y: 5}: true,
				{X: 6, Y: 6}: true,
				{X: 4, Y: 5}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: false,
		},
		{
			name: "Live cell with 5 neighbors dies (overpopulation)",
			cells: Cells{
				{X: 5, Y: 5}: true,
				{X: 5, Y: 6}: true,
				{X: 6, Y: 5}: true,
				{X: 6, Y: 6}: true,
				{X: 4, Y: 5}: true,
				{X: 4, Y: 6}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: false,
		},
		// Dead cell tests
		{
			name:     "Dead cell with 0 neighbors stays dead",
			cells:    Cells{},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		{
			name: "Dead cell with 1 neighbor stays dead",
			cells: Cells{
				{X: 5, Y: 6}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		{
			name: "Dead cell with 2 neighbors stays dead",
			cells: Cells{
				{X: 5, Y: 6}: true,
				{X: 6, Y: 5}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		{
			name: "Dead cell with 3 neighbors becomes alive (reproduction)",
			cells: Cells{
				{X: 5, Y: 6}: true,
				{X: 6, Y: 5}: true,
				{X: 6, Y: 6}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: true,
		},
		{
			name: "Dead cell with 4 neighbors stays dead",
			cells: Cells{
				{X: 5, Y: 6}: true,
				{X: 6, Y: 5}: true,
				{X: 6, Y: 6}: true,
				{X: 4, Y: 5}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		{
			name: "Dead cell with 5 neighbors stays dead",
			cells: Cells{
				{X: 5, Y: 6}: true,
				{X: 6, Y: 5}: true,
				{X: 6, Y: 6}: true,
				{X: 4, Y: 5}: true,
				{X: 4, Y: 6}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: false,
		},
		// Edge cases with diagonal neighbors
		{
			name: "Live cell with 3 diagonal neighbors survives",
			cells: Cells{
				{X: 5, Y: 5}: true,
				{X: 4, Y: 4}: true,
				{X: 4, Y: 6}: true,
				{X: 6, Y: 4}: true,
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    true,
			expected: true,
		},
		{
			name: "Dead cell with 3 mixed neighbors becomes alive",
			cells: Cells{
				{X: 5, Y: 4}: true, // above
				{X: 6, Y: 6}: true, // diagonal
				{X: 4, Y: 5}: true, // left
			},
			cell:     rl.Vector2{X: 5, Y: 5},
			alive:    false,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getLifeStatus(tt.cells, tt.cell, tt.alive)
			if result != tt.expected {
				t.Errorf("getLifeStatus() = %v, want %v", result, tt.expected)
			}
		})
	}
}
