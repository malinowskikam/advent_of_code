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

	acc := 1
	cycle := 0
	strAcc := 0
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		cmd := split[0]

		if cmd == "addx" {
			val, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}

			cycle++
			if (cycle-20)%40 == 0 {
				strAcc += cycle * acc
				fmt.Println(cycle, 0, acc, cycle*acc, strAcc)
			}

			cycle++

			if (cycle-20)%40 == 0 {
				strAcc += cycle * acc
				fmt.Println(cycle, val, acc, cycle*acc, strAcc)
			}
			acc += val
		} else {
			cycle++
			if (cycle-20)%40 == 0 {
				strAcc += cycle * acc
				fmt.Println(cycle, acc, cycle*acc, strAcc)
			}
		}

	}

	fmt.Println(strAcc)
}

func part2() {
}
