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
	fd, err := os.Open("input/input07.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	scanner.Scan()
	line := scanner.Text()

	if line != "$ cd /" {
		panic("File should start with \"$ cd /\"")
	}

	dirs := make(map[string]int64)
	parseDirectory(scanner, "/", dirs)

	sum := int64(0)
	for _, val := range dirs {
		if val <= 100000 {
			sum += val
		}
	}
	fmt.Println(sum)
}

func part2() {
	fd, err := os.Open("input/input07.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	scanner.Scan()
	line := scanner.Text()

	if line != "$ cd /" {
		panic("File should start with \"$ cd /\"")
	}

	dirs := make(map[string]int64)
	parseDirectory(scanner, "/", dirs)

    total := int64(70000000)
    required := int64(30000000)
    current := dirs["/"]
    toDelete := current

    for _, val := range dirs {
        if total - (current - val) >= required && val < toDelete {
            toDelete = val
        }
    }

    fmt.Println(toDelete)
}

func parseDirectory(scanner *bufio.Scanner, name string, dirs map[string]int64) int64 {
	size := int64(0)
	loop := true
	for loop && scanner.Scan() {
		line := strings.Fields(scanner.Text())

		switch {
		case slices.Equal(line, []string{"$", "ls"}):
			// Irrelevant
		case line[0] == "dir":
			// Irrelevant
		case slices.Equal(line[0:2], []string{"$", "cd"}):
			if line[2] == ".." {
				loop = false
			} else {
				size += parseDirectory(scanner, name+line[2]+"/", dirs)
			}
		default:
			val, err := strconv.ParseInt(line[0], 10, 64)
			if err != nil {
				panic(fmt.Sprintf("unknown line: %s\n", line))
			} else {
				size += val
			}
		}
	}

	dirs[name] = size
	return size
}

