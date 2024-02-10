package main

import (
	"fmt"
	"strings"
	"time"
)

type Grid [][]bool

// initializeGrid creates a grid with Glider pattern
func initializeGrid(size int) Grid {
	grid := make(Grid, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}
	grid[size/2][size/2] = true
	grid[size/2][size/2+1] = true
	grid[size/2][size/2+2] = true
	grid[size/2-1][size/2+2] = true
	grid[size/2-2][size/2+1] = true
	return grid
}

// printGrid prints the grid :D
func printGrid(grid Grid) {
	var builder strings.Builder

	for _, row := range grid {
		for _, cell := range row {
			if cell {
				builder.WriteByte('#')
			} else {
				builder.WriteByte('_')
			}
		}
		builder.WriteByte('\n')
	}

	fmt.Print(builder.String())
}

// NextGeneration calculates the next generation of the grid based on Conway's Game of Life rules.
func NextGeneration(currentGrid Grid) Grid {
	size := len(currentGrid)
	newGrid := make(Grid, size)
	for i := range newGrid {
		newGrid[i] = make([]bool, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			liveNeighbors := countLiveNeighbors(currentGrid, i, j)

			// Apply rules of the Game of Life
			if currentGrid[i][j] {
				newGrid[i][j] = liveNeighbors == 2 || liveNeighbors == 3
			} else {
				newGrid[i][j] = liveNeighbors == 3
			}
		}
	}

	return newGrid
}

// countLiveNeighbors counts the number of live neighbors for a given cell.
func countLiveNeighbors(grid Grid, row, col int) int {
	size := len(grid)
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			nx := (row + x + size) % size
			ny := (col + y + size) % size
			if grid[nx][ny] {
				count++
			}
		}
	}
	return count
}

func main() {
	var gridSize int
	fmt.Println("Enter grid size (maximum size is", 250, "):")
	fmt.Scanln(&gridSize)

	if gridSize <= 0 || gridSize > 250 {
		fmt.Println("Invalid grid size. Exiting.")
		return
	}

	grid := initializeGrid(gridSize)

	fmt.Println("Initial Grid:")
	printGrid(grid)
	fmt.Println()

	gen := 0
	for {
		grid = NextGeneration(grid)

		fmt.Printf("Generation %d: \n", gen+1)
		printGrid(grid)
		fmt.Println()
		gen++

		time.Sleep(500 * time.Millisecond)
	}
}
