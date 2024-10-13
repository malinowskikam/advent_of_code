package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

type Monkey struct {
	items     []int
	changeOp  func(int) int
	testDiv   int
	destTrue  int
	destFalse int
	timesInsp int
}

func maxTwo[T any](arr []T, valFunc func(T) int) (int, int) {
	if len(arr) < 2 {
		panic("maxTwo requires array of len >= 2")
	}

	max := math.MinInt
	max2 := math.MinInt

	for _, m := range arr {
		val := valFunc(m)

		if val > max2 {
			max2 = val
		}

		if val > max {
			max2 = max
			max = val
		}
	}

	return max, max2
}

func parseIntPos(line string, pos int) int {
	split := strings.Split(line, " ")
	n, err := strconv.Atoi(split[pos])
	if err != nil {
		panic(err)
	}
	return n
}

func parseOp(line string) func(int) int {
	split := strings.Split(line, " ")
	op := split[6]
	arg := split[7]

	if op == "*" {
		if arg == "old" {
			return func(x int) int { return x * x }
		} else {
			val, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			return func(x int) int { return x * val }
		}
	} else if op == "+" {
		if arg == "old" {
			return func(x int) int { return x + x }
		} else {
			val, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			return func(x int) int { return x + val }
		}
	} else {
		panic("Unhandled op")
	}
}

func parseItems(line string) []int {
	split := strings.Split(line, ":")
	itemsStrs := strings.Split(split[1], ",")
	items := make([]int, 0, len(itemsStrs))
	for _, itemStr := range itemsStrs {
		item, err := strconv.Atoi(strings.TrimSpace(itemStr))
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	return items
}

func parseMonkeys() []Monkey {
	fd, err := os.Open("input/input11.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	monkeys := make([]Monkey, 0)
	var monkeyItems []int
	var monkeyChangeOp func(int) int
	var monkeyTestDiv int
	var monkeyDestTrue int
	var monkeyDestFalse int

	// Parsing monkeys
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			//Finalize monkey
			monkeys = append(monkeys, Monkey{
				items:     monkeyItems,
				changeOp:  monkeyChangeOp,
				testDiv:   monkeyTestDiv,
				destTrue:  monkeyDestTrue,
				destFalse: monkeyDestFalse,
			})

		} else if line[0] == 'M' {
			// Monkey header, assume sorted
			// nothing to do

		} else if line[2] == 'S' {
			split := strings.Split(line, ":")
			itemsStrs := strings.Split(split[1], ",")
			items := make([]int, 0, len(itemsStrs))
			for _, itemStr := range itemsStrs {
				item, err := strconv.Atoi(strings.TrimSpace(itemStr))
				if err != nil {
					panic(err)
				}
				items = append(items, item)
			}
			monkeyItems = items

		} else if line[2] == 'O' {
			monkeyChangeOp = parseOp(line)
		} else if line[2] == 'T' {
			monkeyTestDiv = parseIntPos(line, 5)
		} else if line[8] == 'r' {
			monkeyDestTrue = parseIntPos(line, 9)
		} else if line[8] == 'a' {
			monkeyDestFalse = parseIntPos(line, 9)
		} else {
			panic("Unhandled line")
		}
	}

	// Last monkey
	monkeys = append(monkeys, Monkey{
		items:     monkeyItems,
		changeOp:  monkeyChangeOp,
		testDiv:   monkeyTestDiv,
		destTrue:  monkeyDestTrue,
		destFalse: monkeyDestFalse,
	})

	return monkeys
}

func part1() {
	monkeys := parseMonkeys()

	//Rounds
	for range 20 {
		for i := 0; i < len(monkeys); i++ {
			for _, worry := range monkeys[i].items {
				newVal := monkeys[i].changeOp(worry) / 3
				var dest int
				if newVal%monkeys[i].testDiv == 0 {
					dest = monkeys[i].destTrue
				} else {
					dest = monkeys[i].destFalse
				}
				monkeys[dest].items = append(monkeys[dest].items, newVal)
				monkeys[i].timesInsp++
			}
			monkeys[i].items = make([]int, 0)
		}
	}

	max, max2 := maxTwo(monkeys, func(m Monkey) int { return m.timesInsp })
	fmt.Println(max * max2)
}

func part2() {
	monkeys := parseMonkeys()

	// Common modulo for worry value to keep the worry manageable (otherwise it will overflow long)
	commonMod := 1
	for _, m := range monkeys {
		commonMod *= m.testDiv
	}

	//Rounds
	for range 10000 {
		for i := 0; i < len(monkeys); i++ {
			for _, worry := range monkeys[i].items {
				newVal := monkeys[i].changeOp(worry) % commonMod
				var dest int
				if newVal%monkeys[i].testDiv == 0 {
					dest = monkeys[i].destTrue
				} else {
					dest = monkeys[i].destFalse
				}
				monkeys[dest].items = append(monkeys[dest].items, newVal)
				monkeys[i].timesInsp++
			}
			monkeys[i].items = make([]int, 0)
		}
	}

	max, max2 := maxTwo(monkeys, func(m Monkey) int { return m.timesInsp })
	fmt.Println(max * max2)
}
