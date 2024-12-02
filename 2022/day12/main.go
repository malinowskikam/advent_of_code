package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	part1()
	part2()
}

type Point struct {
	row       int
	col       int
	elevation byte
	distance  int

	//PQ index
	index int
}

// PRIORITY QUEUE

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Point)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Point, distance int) {
	item.distance = distance
	heap.Fix(pq, item.index)
}

// END OF PRIORITY QUEUE

func (p *Point) isNeighbor(o *Point) bool {
	return p.elevation >= o.elevation-1
}

func part1() {
	fd, err := os.Open("input/input12.txt")
	if err != nil {
		panic(err)
	}

	grid := make([][]Point, 0)

	var end_row, end_col int

	scanner := bufio.NewScanner(fd)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)

		row := make([]Point, len(bytes))
		for j, b := range bytes {
			row[j].row = i
			row[j].col = j
			row[j].elevation = b
			row[j].distance = math.MaxInt
			if b == 'S' {
				row[j].elevation = 'a'
				row[j].distance = 0
			} else if b == 'E' {
				row[j].elevation = 'z'
				end_row = i
				end_col = j
			}
		}

		grid = append(grid, row)
		i++
	}

	// Dijkstra's
	pq := make(PriorityQueue, len(grid)*len(grid[0]))

	for i := range grid {
		for j := range grid[i] {
			pq[i*len(grid[0])+j] = &grid[i][j]
			grid[i][j].index = i*len(grid[0]) + j
		}
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Point)

		if p.row == end_row && p.col == end_col {
			fmt.Println(p.distance)
			break
		}

		d := p.distance + 1

		//above neighbor
		if p.row > 0 {
			o := &grid[p.row-1][p.col]
			if p.isNeighbor(o) && o.distance > d {
				pq.Update(o, d)
			}
		}

		//below neighbor
		if p.row < len(grid)-1 {
			o := &grid[p.row+1][p.col]
			if p.isNeighbor(o) && o.distance > d {
				pq.Update(o, d)
			}
		}

		//left neighbor
		if p.col > 0 {
			o := &grid[p.row][p.col-1]
			if p.isNeighbor(o) && o.distance > d {
				pq.Update(o, d)
			}
		}

		//right neighbor
		if p.col < len(grid[p.row])-1 {
			o := &grid[p.row][p.col+1]
			if p.isNeighbor(o) && o.distance > d {
				pq.Update(o, d)
			}
		}
	}

}

func (p *Point) isNeighbor2(o *Point) bool {
	return p.elevation-1 <= o.elevation
}

func part2() {
	fd, err := os.Open("input/input12.txt")
	if err != nil {
		panic(err)
	}

	grid := make([][]Point, 0)

	scanner := bufio.NewScanner(fd)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)

		row := make([]Point, len(bytes))
		for j, b := range bytes {
			row[j].row = i
			row[j].col = j
			row[j].elevation = b
			row[j].distance = math.MaxInt
			if b == 'S' {
				row[j].elevation = 'a'
			} else if b == 'E' {
				row[j].elevation = 'z'
				row[j].distance = 0
			}
		}

		grid = append(grid, row)
		i++
	}

	// Dijkstra's
	pq := make(PriorityQueue, len(grid)*len(grid[0]))

	for i := range grid {
		for j := range grid[i] {
			pq[i*len(grid[0])+j] = &grid[i][j]
			grid[i][j].index = i*len(grid[0]) + j
		}
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Point)

		if p.elevation == 'a' {
			fmt.Println(p.distance)
			break
		}

		d := p.distance + 1

		//above neighbor
		if p.row > 0 {
			o := &grid[p.row-1][p.col]
			if p.isNeighbor2(o) && o.distance > d {
				pq.Update(o, d)
			}
		}

		//below neighbor
		if p.row < len(grid)-1 {
			o := &grid[p.row+1][p.col]
			if p.isNeighbor2(o) && o.distance > d {
				pq.Update(o, d)
			}
		}

		//left neighbor
		if p.col > 0 {
			o := &grid[p.row][p.col-1]
			if p.isNeighbor2(o) && o.distance > d {
				pq.Update(o, d)
			}
		}

		//right neighbor
		if p.col < len(grid[p.row])-1 {
			o := &grid[p.row][p.col+1]
			if p.isNeighbor2(o) && o.distance > d {
				pq.Update(o, d)
			}
		}
	}
}
