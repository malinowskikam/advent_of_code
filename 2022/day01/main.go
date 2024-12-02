package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	fd, err := os.Open("input/input01.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	maxValue := 0
	currentValue := 0

	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		line := scanner.Text()

		if len(line) > 0 {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currentValue += num
		} else {
			if currentValue > maxValue {
				maxValue = currentValue
			}
			currentValue = 0
		}
	}
	fmt.Println(maxValue)
}

func part2() {
	fd, err := os.Open("input/input01.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	maxValue := []int{0, 0, 0}
	currentValue := 0

	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		line := scanner.Text()

		if len(line) > 0 {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currentValue += num
		} else {
			// Finish a sequence
			for i := 0; i < len(maxValue); i++ {
				if currentValue > maxValue[i] {
					// Move lower values left
					if i > 0 {
						maxValue[i-1] = maxValue[i]
					}
					maxValue[i] = currentValue
				}
			}
			currentValue = 0
		}
	}

	sum := 0
	for _, num := range maxValue {
		sum += num
	}
	fmt.Println(sum)
}
