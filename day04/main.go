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
	fd, err := os.Open("input/input04.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineSplit := strings.Split(line, ",")
		left, right := lineSplit[0], lineSplit[1]
		leftSplit := strings.Split(left, "-")
		leftStart, err := strconv.Atoi(leftSplit[0])
		if err != nil {
			panic(err)
		}
		leftEnd, err := strconv.Atoi(leftSplit[1])
		if err != nil {
			panic(err)
		}
		rightSplit := strings.Split(right, "-")
		rightStart, err := strconv.Atoi(rightSplit[0])
		if err != nil {
			panic(err)
		}
		rightEnd, err := strconv.Atoi(rightSplit[1])
		if err != nil {
			panic(err)
		}

		if leftStart >= rightStart && leftEnd <= rightEnd || rightStart >= leftStart && rightEnd <= leftEnd {
			count += 1
		}
	}
	fmt.Println(count)
}

func part2() {
	fd, err := os.Open("input/input04.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineSplit := strings.Split(line, ",")
		left, right := lineSplit[0], lineSplit[1]
		leftSplit := strings.Split(left, "-")
		leftStart, err := strconv.Atoi(leftSplit[0])
		if err != nil {
			panic(err)
		}
		leftEnd, err := strconv.Atoi(leftSplit[1])
		if err != nil {
			panic(err)
		}
		rightSplit := strings.Split(right, "-")
		rightStart, err := strconv.Atoi(rightSplit[0])
		if err != nil {
			panic(err)
		}
		rightEnd, err := strconv.Atoi(rightSplit[1])
		if err != nil {
			panic(err)
		}

		if leftStart <= rightEnd && leftEnd >= rightStart {
			count += 1
		}
	}
	fmt.Println(count)
}
