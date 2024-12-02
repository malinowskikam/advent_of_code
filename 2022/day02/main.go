package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	fd, err := os.Open("input/input02.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	resultValues := []int{0, 6, 3}
	score := 0

	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		line := scanner.Text()

		split := strings.Fields(line)
		opPick := int(split[0][0]-'A') + 1
		myPick := int(split[1][0]-'X') + 1
		result := (opPick - myPick + 2) % 3

		score += myPick + resultValues[result]
	}
	fmt.Println(score)
}

func part2() {
	fd, err := os.Open("input/input02.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	resultValues := []int{0, 3, 6}
	score := 0

	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		line := scanner.Text()

		split := strings.Fields(line)
		opPick := int(split[0][0]-'A') + 1
		result := int(split[1][0] - 'X')
		myPick := (opPick+result+1)%3 + 1

		score += myPick + resultValues[result]

	}
	fmt.Println(score)
}
