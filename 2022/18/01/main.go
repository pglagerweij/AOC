package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// var inputFile string = "triaal2.txt"

// var inputFile string = "trial.txt"

var inputFile string = "input.txt"

type Point struct {
	X, Y, Z int
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y, p.Z + q.Z}
}

func main() {

	input, _ := os.ReadFile(inputFile)

	lava := map[Point]bool{}
	minNumbers := Point{math.MaxInt, math.MaxInt, math.MaxInt}
	maxNumbers := Point{math.MinInt, math.MinInt, math.MinInt}

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var p Point
		fmt.Sscanf(s, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		lava[p] = true

		minNumbers = Point{Min(p.X, minNumbers.X), Min(p.Y, minNumbers.Y), Min(p.Z, minNumbers.Z)}
		maxNumbers = Point{Max(p.X, maxNumbers.X), Max(p.Y, maxNumbers.Y), Max(p.Z, maxNumbers.Z)}
	}

	delta := []Point{
		{-1, 0, 0}, {0, -1, 0}, {0, 0, -1},
		{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
	}

	totalOutside := 0
	for point := range lava {
		for _, d := range delta {
			if !lava[point.Add(d)] {
				totalOutside += 1
			}
		}

	}
	fmt.Printf("The total outisde points is %v\n", totalOutside)

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
