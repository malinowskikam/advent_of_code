package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

type Pos struct {
	x int
	y int
}

func (p *Pos) step(dir string) {
	switch dir {
	case "L":
		p.x--
	case "U":
		p.y++
	case "R":
		p.x++
	case "D":
		p.y--
	default:
		panic("unknown direction")
	}
}

func (p *Pos) follow(head *Pos) {
	if abs(head.x-p.x) <= 1 && abs(head.y-p.y) <= 1 {
		return
	} else {
		if head.x != p.x {
			p.x += diffSign(head.x, p.x)
		}
		if head.y != p.y {
			p.y += diffSign(head.y, p.y)
		}
	}

}

func part1() {
	fd, err := os.Open("input/input09.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	visited := make(map[Pos]bool)
	posTail := Pos{x: 0, y: 0}
	posHead := Pos{x: 0, y: 0}
	visited[posTail] = true

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		dir := split[0]
		steps, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		for range steps {
			posHead.step(dir)
			posTail.follow(&posHead)
			visited[posTail] = true
		}
	}

	fmt.Println(len(visited))
}

func part2() {
	fd, err := os.Open("input/input09.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	visited := make(map[Pos]bool)
	posKnots := make([]Pos, 10)
	for i := 0; i < len(posKnots); i++ {
		posKnots[i] = Pos{x: 0, y: 0}
	}
	visited[posKnots[len(posKnots)-1]] = true

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		dir := split[0]
		steps, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		for range steps {
			posKnots[0].step(dir)
			for i := 1; i < len(posKnots); i++ {
				posKnots[i].follow(&posKnots[i-1])
			}
			visited[posKnots[len(posKnots)-1]] = true
		}
	}

	fmt.Println(len(visited))
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func diffSign(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
