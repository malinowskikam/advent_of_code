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
		line_split := len(line) / 2
		line_left := line[:line_split]
		line_right := line[line_split:]

		set := map[byte]bool{}
		common_value := byte(0)
		for _, b := range []byte(line_left) {
			set[b] = true
		}

		for _, b := range []byte(line_right) {
			_, exists := set[b]
			if exists {
				common_value = b
				break
			}
		}

		if common_value >= 'A' && common_value <= 'Z' {
			sum += int(common_value - 'A' + 27)
		} else if common_value >= 'a' && common_value <= 'z' {
			sum += int(common_value - 'a' + 1)
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

			set_first_line := map[byte]bool{}
			for _, b := range []byte(lines[0]) {
				set_first_line[b] = true
			}

			set_first_two_lines := map[byte]bool{}
			for _, b := range []byte(lines[1]) {
				_, exists := set_first_line[b]
				if exists {
					set_first_two_lines[b] = true
				}
			}

			common_value := byte(0)
			for _, b := range []byte(lines[2]) {
				_, exists := set_first_two_lines[b]
				if exists {
					common_value = b
					break
				}
			}

			fmt.Println(common_value, string(common_value))

			if common_value >= 'A' && common_value <= 'Z' {
				sum += int(common_value - 'A' + 27)
			} else if common_value >= 'a' && common_value <= 'z' {
				sum += int(common_value - 'a' + 1)
			}
		}

	}
	fmt.Println(sum)
}
