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

type LineIterator struct {
	buf string
	idx int
}

func (it *LineIterator) Next() byte {
	if it.idx < len(it.buf) {
		item := it.buf[it.idx]
		it.idx++
		return item
	} else {
		return 0
	}
}

func parse_array(it *LineIterator) []interface{} {
	arr := make([]interface{}, 0)

	num := -1

	b := it.Next()
	for b != 0 {

		switch b {
		case '[':
			arr = append(arr, parse_array(it))
		case ']':
			if num != -1 {
				arr = append(arr, num)
			}
			return arr
		case ',':
			if num != -1 {
				arr = append(arr, num)
				num = -1
			}
		default:
			if b < '0' || b > '9' {
				panic("b is not a digit")
			}
			if num == -1 {
				num = 0
			}
			num = num*10 + int(b-'0')
		}

		b = it.Next()
	}

	return arr
}

func compareArray(arr1 []interface{}, arr2 []interface{}) (bool, bool) {
	i := 0
	j := 0

	for i < len(arr1) && j < len(arr2) {
		int1, ok1 := arr1[i].(int)
		int2, ok2 := arr2[j].(int)
		if ok1 && ok2 {
			if int1 < int2 {
				return true, true
			} else if int1 > int2 {
				return false, true
			}
		} else if !ok1 && !ok2 {
			right, ans := compareArray(arr1[i].([]interface{}), arr2[j].([]interface{}))
			if ans {
				return right, true
			}
		} else {
			arg1 := arr1[i]
			if ok1 {
				arg1 = []interface{}{arr1[i].(int)}
			}

			arg2 := arr2[j]
			if ok2 {
				arg2 = []interface{}{arr2[j].(int)}
			}
			right, ans := compareArray(arg1.([]interface{}), arg2.([]interface{}))
			if ans {
				return right, true
			}
		}
		i++
		j++
	}

	if len(arr1) < len(arr2) {
		return true, true
	} else if len(arr1) > len(arr2) {
		return false, true
	}

	return false, false
}

func part1() {
	fd, err := os.Open("input/input13.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)
	inProgress := true

	i := 1
	sum := 0
	for inProgress {
		scanner.Scan()
		line1 := scanner.Text()
		arr1 := parse_array(&LineIterator{buf: line1, idx: 1})

		scanner.Scan()
		line2 := scanner.Text()
		arr2 := parse_array(&LineIterator{buf: line2, idx: 1})

		if right, ans := compareArray(arr1, arr2); right && ans {
			sum += i
		}

		inProgress = scanner.Scan()
		i++
	}
	fmt.Println(sum)
}

func part2() {
	fd, err := os.Open("input/input13.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	div1 := parse_array(&LineIterator{buf: "[[2]]", idx: 1})
	div2 := parse_array(&LineIterator{buf: "[[6]]", idx: 1})

	idx1 := 1
	idx2 := 2

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			arr := parse_array(&LineIterator{buf: line, idx: 1})

			if r, _ := compareArray(arr, div1); r {
				idx1++
			}

			if r, _ := compareArray(arr, div2); r {
				idx2++
			}
		}
	}
	fmt.Println(idx1 * idx2)
}
