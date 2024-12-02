package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	part1()
	part2()
}

type Cell struct {
	val     int
	visible bool
}

func part1() {
	fd, err := os.Open("input/input08.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	grid := make([][]Cell, 0)
	var colHigh []int

	// first pass
	// Check the visibility from top and left, generate grid
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		row := make([]Cell, lineLen)
		grid = append(grid, row)

		leftHigh := -1
		if colHigh == nil {
			// generate topHigh slice
			colHigh = make([]int, len(line))
			fill(colHigh, -1)
		}

		for i := 0; i < lineLen; i++ {

			visible := false
			val := int(line[i])

			if val > leftHigh {
				leftHigh = val
				visible = true
			}

			if val > colHigh[i] {
				colHigh[i] = val
				visible = true
			}

			row[i].visible = visible
			row[i].val = val

		}
	}

	// second pass
	// Check the visibility from right and bottom, iterating in reverse order

	fill(colHigh, -1)
	count := 0
	for i := len(grid) - 1; i >= 0; i-- {
		rightHigh := -1

		for j := len(grid[i]) - 1; j >= 0; j-- {

			if grid[i][j].val > rightHigh {
				rightHigh = grid[i][j].val
				grid[i][j].visible = true
			}

			if grid[i][j].val > colHigh[j] {
				colHigh[j] = grid[i][j].val
				grid[i][j].visible = true
			}

			if grid[i][j].visible {
				count++
			}
		}
	}

	fmt.Println(count)
}

type Cell2 struct {
	val int

	//We only need to cache the up and left
	visibilityLeft int
	visibilityUp   int
}

func part2() {
	fd, err := os.Open("input/input08.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	grid := make([][]Cell2, 0)

	var lastSeenCol [][]int

	// first pass
	// Check the visibility from top and left, generate grid

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		row := make([]Cell2, lineLen)
		grid = append(grid, row)

		lastSeenLeft := make([]int, 10)
		fill(lastSeenLeft, -1)

		if lastSeenCol == nil {
			lastSeenCol = make([][]int, len(line))
			for i := 0; i < len(line); i++ {
				lastSeenCol[i] = make([]int, 10)
				fill(lastSeenCol[i], -1)
			}
		}

		for j := 0; j < lineLen; j++ {
			val := int(line[j] - '0')
			row[j].val = val

			blockingIndexLeft := slices.Max(lastSeenLeft[val:])

			if blockingIndexLeft == -1 {
				row[j].visibilityLeft = j
			} else {
				row[j].visibilityLeft = j - blockingIndexLeft
			}

			blockingIndexUp := slices.Max(lastSeenCol[j][val:])

			if blockingIndexUp == -1 {
				row[j].visibilityUp = i
			} else {
				row[j].visibilityUp = i - blockingIndexUp
			}

			lastSeenLeft[val] = j   // Store last column (j) of val
			lastSeenCol[j][val] = i // Store last row (i) of val in column j
		}

		i++
	}

	// second pass
	// Check the visibility from right and bottom, iterating in reverse order

	fill2(lastSeenCol, math.MaxInt)
	maxScore := 0
	for i := len(grid) - 1; i >= 0; i-- {
		lastSeenRight := make([]int, 10)
		fill(lastSeenRight, math.MaxInt)

		for j := len(grid[i]) - 1; j >= 0; j-- {
			cell := &grid[i][j]
			blockingIndexRight := slices.Min(lastSeenRight[cell.val:])

			var visibilityRight int
			if blockingIndexRight == math.MaxInt {
				visibilityRight = len(grid[i]) - j - 1
			} else {
				visibilityRight = blockingIndexRight - j
			}

			blockingIndexDown := slices.Min(lastSeenCol[j][cell.val:])

			var visibilityDown int
			if blockingIndexDown == math.MaxInt {
				visibilityDown = len(grid) - i - 1
			} else {
				visibilityDown = blockingIndexDown - i
			}

			score := cell.visibilityLeft * cell.visibilityUp * visibilityRight * visibilityDown
			if score > maxScore {
				maxScore = score
			}

			lastSeenRight[cell.val] = j  // Store last column (j) of val
			lastSeenCol[j][cell.val] = i // Store last row (i) of val in column j
		}
	}

	fmt.Println(maxScore)
}

func fill(arr []int, val int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = val
	}
}

func fill2(arr [][]int, val int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = val
		}
	}
}
