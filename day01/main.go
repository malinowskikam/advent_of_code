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

	max_value := 0
	current_value := 0

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
			current_value += num
		} else {
			if current_value > max_value {
				max_value = current_value
			}
			current_value = 0
		}
	}
	fmt.Println(max_value)
}

func part2() {
	fd, err := os.Open("input/input01.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	max_values := []int{0, 0, 0}

	current_value := 0

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
			current_value += num
		} else {
			// Finish a sequence
			for i := 0; i < len(max_values); i++ {
				if current_value > max_values[i] {
					// Move lower values left
					if i > 0 {
						max_values[i-1] = max_values[i]
					}
					max_values[i] = current_value
				}
			}
			current_value = 0
		}
	}

	sum := 0
	for _, num := range max_values {
		sum += num
	}
	fmt.Println(sum)
}
