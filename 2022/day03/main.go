package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	fd, err := os.Open("input/input03.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := len(line) / 2
		lineLeft := line[:lineSplit]
		lineRight := line[lineSplit:]

		set := map[byte]bool{}
		commonValue := byte(0)
		for _, b := range []byte(lineLeft) {
			set[b] = true
		}

		for _, b := range []byte(lineRight) {
			_, exists := set[b]
			if exists {
				commonValue = b
				break
			}
		}

		if commonValue >= 'A' && commonValue <= 'Z' {
			sum += int(commonValue - 'A' + 27)
		} else if commonValue >= 'a' && commonValue <= 'z' {
			sum += int(commonValue - 'a' + 1)
		}
	}
	fmt.Println(sum)
}

func part2() {
	fd, err := os.Open("input/input03.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	sum := 0

	var lines = make([]string, 3)
	var index = -1

	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		index++
		lines[index] = scanner.Text()

		if index == 2 {
			index = -1

			setFirstLine := map[byte]bool{}
			for _, b := range []byte(lines[0]) {
				setFirstLine[b] = true
			}

			setFirstTwoLines := map[byte]bool{}
			for _, b := range []byte(lines[1]) {
				_, exists := setFirstLine[b]
				if exists {
					setFirstTwoLines[b] = true
				}
			}

			commonValue := byte(0)
			for _, b := range []byte(lines[2]) {
				_, exists := setFirstTwoLines[b]
				if exists {
					commonValue = b
					break
				}
			}

			fmt.Println(commonValue, string(commonValue))

			if commonValue >= 'A' && commonValue <= 'Z' {
				sum += int(commonValue - 'A' + 27)
			} else if commonValue >= 'a' && commonValue <= 'z' {
				sum += int(commonValue - 'a' + 1)
			}
		}

	}
	fmt.Println(sum)
}
