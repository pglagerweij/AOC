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

	// Make larger grid with a lot of falses around the lava points
	for x := minNumbers.X - 1; x <= maxNumbers.X+1; x++ {
		for y := minNumbers.Y - 1; y <= maxNumbers.Y+1; y++ {
			for z := minNumbers.Z - 1; z <= maxNumbers.Z+1; z++ {
				lava[Point{x, y, z}] = lava[Point{x, y, z}]
			}
		}
	}

	rest := []Point{minNumbers}
	visit := map[Point]struct{}{minNumbers: {}}

	totalOutside2 := 0
	for len(rest) > 0 {
		current := rest[0]
		// fmt.Printf("current checking %v with length %v\n", current, len(rest))
		rest = rest[1:]

		for _, d := range delta {
			nextpos := current.Add(d)
			cube, valid := lava[nextpos]
			_, seen := visit[nextpos]
			if cube {
				totalOutside2++
			} else if valid && !seen {
				visit[nextpos] = struct{}{}
				rest = append(rest, nextpos)
				// fmt.Printf("the rest is %v with length %v\n", rest, len(rest))
			}
		}
	}
	fmt.Printf("the total lava is %v\n", totalOutside2)

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

// for x := min.X - 1; x <= max.X+1; x++ {
// 	for y := min.Y - 1; y <= max.Y+1; y++ {
// 		for z := min.Z - 1; z <= max.Z+1; z++ {
// 			lava[Point{x, y, z}] = lava[Point{x, y, z}]
// 		}
// 	}
// }

// queue := []Point{min}
// visited := map[Point]struct{}{min: {}}

// part2 := 0
// for len(queue) > 0 {
// 	cur := queue[0]
// 	queue = queue[1:]

// 	for _, d := range delta {
// 		next := cur.Add(d)

// 		if cube, valid := lava[next]; cube {
// 			part2++
// 		} else if _, seen := visited[next]; valid && !seen {
// 			visited[next] = struct{}{}
// 			queue = append(queue, next)
// 		}
// 	}
// }
// fmt.Println(part2)
