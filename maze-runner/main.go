package main

import (
	"fmt"
)

func main() {
	maze := [6][8]bool{
		{false, false, false, false, false, false, false, false},
		{false, true, true, true, true, true, true, false},
		{false, true, false, false, false, true, true, false},
		{false, true, true, true, false, true, false, false},
		{false, true, false, true, true, true, true, false},
		{false, false, false, false, false, false, false, false},
	}
	if len(maze) == 0 {
		return
	}
	totalSolution := 0
	for y := 1; y < len(maze)-1; y++ {
		// Init ColumnStart
		columnStart := 4
		// Init RowStart
		rowStart := 1

		var isCanGoNorth bool
		north := 0
		for north < y {
			columnStart--

			isCanGoNorth = maze[columnStart][rowStart]
			if !isCanGoNorth {
				break
			}
			north++
		}

		var isCanGoEast bool
		east := 0
		for east < y {
			rowStart++
			isCanGoEast = maze[columnStart][rowStart]
			if !isCanGoEast {
				break
			}
			east++
		}

		var isCanGoSouth bool
		south := 0
		for south < y {
			columnStart++
			isCanGoSouth = maze[columnStart][rowStart]
			if !isCanGoSouth {
				break
			}
			south++
		}
		if isCanGoNorth && isCanGoEast && isCanGoSouth {
			totalSolution++
		}
	}
	fmt.Println("TotalSolution", totalSolution)
}
