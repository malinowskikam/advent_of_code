package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Line struct {
	a, b *Point
}

func main() {
	part1()
	part2()
}

func gatherLines(file string) ([]Line, int, int, int) {
	fd, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	lines := make([]Line, 0)
	minX := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt

	//Gather lines
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " -> ")
		var lastP *Point = nil
		for _, ps := range split {
			coords := strings.Split(ps, ",")
			x, err := strconv.Atoi(coords[0])
			if err != nil {
				panic(err)
			}

			if x-1 < minX {
				minX = x - 1
			}
			if x+1 > maxX {
				maxX = x + 1
			}

			y, err := strconv.Atoi(coords[1])
			if err != nil {
				panic(err)
			}

			if y+1 > maxY {
				maxY = y + 1
			}

			p := &Point{x: x, y: y}
			if lastP != nil {
				lines = append(lines, Line{a: lastP, b: p})
			}

			lastP = p
		}
	}
	return lines, minX, maxX, maxY
}

func createGrid(minX int, maxX int, maxY int) [][]rune {
	grid := make([][]rune, maxY+1)
	for i := range len(grid) {
		grid[i] = make([]rune, maxX-minX+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

func renderLines(grid [][]rune, lines []Line, minX int) {
	for _, line := range lines {

		if line.a.x == line.b.x {
			// horizontal line
			y1, y2 := line.a.y, line.b.y
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				grid[i][line.a.x-minX] = '#'
			}
		} else if line.a.y == line.b.y {
			// vertical line
			x1, x2 := line.a.x, line.b.x
			if x1 > x2 {
				x1, x2 = x2, x1
			}

			for i := x1; i <= x2; i++ {
				grid[line.a.y][i-minX] = '#'
			}
		} else {
			panic("line has to be axis aligned")
		}
	}
}

func simulate(grid [][]rune, minX int, maxY int, render bool, part2 bool) int {
	inProgress := true
	sum := 0

	curSandX := -1
	curSandY := -1

	for inProgress {
		if grid[0][500-minX] != '.' {
			break
		}

		if curSandX == -1 {
			curSandX = 500
			curSandY = 0
			grid[curSandY][curSandX-minX] = '+'
			sum++
		}

		//Physics
		if curSandY < maxY && grid[curSandY+1][curSandX-minX] == '.' {
			// Move down
			grid[curSandY][curSandX-minX] = '.'
			grid[curSandY+1][curSandX-minX] = '+'
			curSandY++
		} else if curSandY < maxY && grid[curSandY+1][curSandX-minX-1] == '.' {
			// Move down-left
			grid[curSandY][curSandX-minX] = '.'
			grid[curSandY+1][curSandX-minX-1] = '+'
			curSandY++
			curSandX--
		} else if curSandY < maxY && grid[curSandY+1][curSandX-minX+1] == '.' {
			// Move down-left
			grid[curSandY][curSandX-minX] = '.'
			grid[curSandY+1][curSandX-minX+1] = '+'
			curSandY++
			curSandX++
		} else {
			grid[curSandY][curSandX-minX] = 'O'
			curSandX = -1
			curSandY = -1
		}

		if curSandY == maxY && !part2 {
			inProgress = false
			sum--
		}

		// Animation only works fine for the example
		if render {
			printGrid(grid)
			time.Sleep(30 * time.Millisecond)
		}
	}

	return sum
}

func printGrid(grid [][]rune) {
	fmt.Print("\033[H\033[2J")
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
}

func part1() {
	lines, minX, maxX, maxY := gatherLines("input/input14.txt")
	grid := createGrid(minX, maxX, maxY)
	renderLines(grid, lines, minX)
	sum := simulate(grid, minX, maxY, false, false)

	fmt.Println(sum)
}

func part2() {
	lines, minX, maxX, maxY := gatherLines("input/input14.txt")
	minX = 500 - maxY - 1
	maxX = 500 + maxY + 1
	grid := createGrid(minX, maxX, maxY)
	renderLines(grid, lines, minX)
	sum := simulate(grid, minX, maxY, false, true)

	fmt.Println(sum)
}
