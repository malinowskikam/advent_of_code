package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	fd, err := os.Open("input/input05.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	var stackCount int
	var stacks [][]byte

	for scanner.Scan() {
		line := scanner.Text()

		if line[1] == '1' {
			break
		}

		if stacks == nil {
			stackCount = (len(line) + 1) / 4

			stacks = make([][]byte, stackCount)
			for i := range stacks {
				stacks[i] = make([]byte, 0)
			}
		}

		for i := 0; i < stackCount; i++ {
			if line[i*4+1] != ' ' {
				stacks[i] = append(stacks[i], line[i*4+1])
			}
		}
	}

	for i := range stacks {
		slices.Reverse(stacks[i])
	}

	scanner.Scan() // Skip empty line

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		n, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(fields[3])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(fields[5])
		if err != nil {
			panic(err)
		}

		for i := 0; i < n; i++ {
			v := stacks[from-1][len(stacks[from-1])-1]
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
			stacks[to-1] = append(stacks[to-1], v)
		}
	}

	for i := range stacks {
		for j := range stacks[i] {
			fmt.Printf("[%c] ", stacks[i][j])
		}
		fmt.Println()
	}
}

func part2() {
	fd, err := os.Open("input/input05.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	var stackCount int
	var stacks [][]byte

	for scanner.Scan() {
		line := scanner.Text()

		if line[1] == '1' {
			break
		}

		if stacks == nil {
			stackCount = (len(line) + 1) / 4

			stacks = make([][]byte, stackCount)
			for i := range stacks {
				stacks[i] = make([]byte, 0)
			}
		}

		for i := 0; i < stackCount; i++ {
			if line[i*4+1] != ' ' {
				stacks[i] = append(stacks[i], line[i*4+1])
			}
		}
	}

	for i := range stacks {
		slices.Reverse(stacks[i])
	}

	scanner.Scan() // Skip empty line

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		n, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(fields[3])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(fields[5])
		if err != nil {
			panic(err)
		}

		v := stacks[from-1][len(stacks[from-1])-n : len(stacks[from-1])]
		stacks[from-1] = stacks[from-1][:len(stacks[from-1])-n]
		stacks[to-1] = append(stacks[to-1], v...)
	}

	for i := range stacks {
		for j := range stacks[i] {
			fmt.Printf("[%c] ", stacks[i][j])
		}
		fmt.Println()
	}
}
