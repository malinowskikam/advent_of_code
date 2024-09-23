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
	fd, err := os.Open("input/input06.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	scanner.Scan()
	line := scanner.Text()

	i_beg := 0
	i_end := 4

	for i_end <= len(line) {
		chars := make(map[rune]bool)

		for _, rune := range line[i_beg:i_end] {
			chars[rune] = true
		}

		if len(chars) == 4 {
			break
		}

		i_beg++
		i_end++
	}

	fmt.Println(i_end)

}

func part2() {
	fd, err := os.Open("input/input06.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	scanner.Scan()
	line := scanner.Text()

	i_beg := 0
	i_end := 14

	for i_end <= len(line) {
		chars := make(map[rune]bool)

		for _, rune := range line[i_beg:i_end] {
			chars[rune] = true
		}

		if len(chars) == 14 {
			break
		}

		i_beg++
		i_end++
	}

	fmt.Println(i_end)
}
