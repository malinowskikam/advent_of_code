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

func part1() {
	fd, err := os.Open("input/input10.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	cycle := 1
	regX := 1
	sSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		cmd := split[0]

		if cmd == "addx" {
			val, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}

			//first cycle
			if (cycle-20)%40 == 0 {
				sSum += cycle * regX
			}
			cycle++

			//second cycle
			if (cycle-20)%40 == 0 {
				sSum += cycle * regX
			}
			// regX changed AFTER the second cycke
			regX += val
			cycle++

		} else {
			//first cycle
			if (cycle-20)%40 == 0 {
				sSum += cycle * regX
			}
			cycle++
		}

	}

	fmt.Println(sSum)
}

func drawPixel(pixels []rune, regX int, cycle int) {
	if cycle > 240 {
		panic("Cycle > 240")
	}

	beamPos := (cycle - 1) % 40
	if regX-1 <= beamPos && regX+1 >= beamPos {
		pixels[cycle-1] = '#'
	} else {
		pixels[cycle-1] = '.'
	}
}

func part2() {
	fd, err := os.Open("input/input10.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	cycle := 1
	regX := 1
	pixels := make([]rune, 240)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		cmd := split[0]

		if cmd == "addx" {
			val, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}

			//first cycle

			drawPixel(pixels, regX, cycle)
			cycle++

			//second cycle
			drawPixel(pixels, regX, cycle)
			cycle++
			regX += val
		} else {
			//first cycle
			drawPixel(pixels, regX, cycle)
			cycle++
		}

	}

	for i := range 6 {
		for j := range 40 {
			fmt.Print(string(pixels[i*40+j]))
		}
		fmt.Println()
	}
}
