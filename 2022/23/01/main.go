package main

import (
	"fmt"
	"os"
	"strings"
)

// var inputFile string = "trial.txt"

var inputFile string = "input.txt"

var directions map[int]string = map[int]string{
	0: "North",
	1: "South",
	2: "West",
	3: "East",
}

var Movement map[int]Position = map[int]Position{
	0: {X: -1, Y: 0},
	1: {X: 1, Y: 0},
	2: {X: 0, Y: -1},
	3: {X: 0, Y: 1},
}

var Options map[int][]Position = map[int][]Position{
	0: {{X: -1, Y: 0}, {X: -1, Y: 1}, {X: -1, Y: -1}},
	1: {{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: -1}},
	2: {{X: 0, Y: -1}, {X: 1, Y: -1}, {X: -1, Y: -1}},
	3: {{X: 0, Y: 1}, {X: 1, Y: 1}, {X: -1, Y: 1}},
}

var Allround []Position = []Position{
	{X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: -1}, {X: 0, Y: 1}, {X: 1, Y: -1}, {X: 1, Y: 0}, {X: 1, Y: 1},
}

type Position struct {
	X, Y int
}

func (p Position) Add(q Position) Position {
	return Position{p.X + q.X, p.Y + q.Y}
}

func (p Position) Min(q Position) Position {
	return Position{p.X - q.X, p.Y - q.Y}
}

var positions = map[Position]bool{}
var currentdirection int = 0

func main() {
	input, _ := os.ReadFile(inputFile)
	// Construct Map element
	for index, row := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for charnumber, elem := range row {
			if elem == '#' {
				positions[Position{X: index, Y: charnumber}] = true
			}
		}
	}

	for time := 0; time < 10; time++ {
		currentdirection := time % 4
		newpositionsmap := map[Position]Position{}
		newpositions := map[Position]bool{}
		// Propose new positions
		for pos := range positions {
			// fmt.Printf("The position is %v \n", pos)
			newpos := pos
			nomovement := true
			for _, direct := range Allround {
				if positions[pos.Add(direct)] == true {
					nomovement = false
					break
				}
			}
			if nomovement {
				newpositions[pos] = true
				newpositionsmap[pos] = pos
				continue
			}
			for direction := currentdirection; direction < currentdirection+4; direction++ {
				// fmt.Printf("We are checking direction %v\n", directions[direction%4])
				valid := true
				for _, move := range Options[direction%4] {
					if positions[pos.Add(move)] {
						valid = false

						break
					}
				}
				if valid {
					newpos = pos.Add(Movement[direction%4])
					break
				}

			}
			newpositions[newpos] = true
			newpositionsmap[pos] = newpos
		}
		// fmt.Printf("The mapped positions are %v\n", newpositionsmap)

		stepPositions := map[Position]bool{}
		// Loop over proposed and determine unique
		if len(newpositions) == len(positions) { // Positions unique so continue at once
			stepPositions = newpositions
		} else {
			for oldpos, newpos := range newpositionsmap {
				count := 0
				for oldposCheck, newposCheck := range newpositionsmap {
					if oldposCheck != oldpos && newposCheck == newpos {
						count = 2
						stepPositions[oldpos] = true
						stepPositions[oldposCheck] = true
						delete(newpositionsmap, oldpos)
						delete(newpositionsmap, oldposCheck)
						break

					}
				}
				if count == 0 {
					delete(newpositionsmap, oldpos)
					stepPositions[newpos] = true
				}
			}

		}
		// fmt.Printf("The new positions are %v\n", stepPositions)
		positions = stepPositions
	}

	minX := 99999999
	maxX := -9999999999
	minY := 9999999999999
	maxY := -9999999999999999
	for pos := range positions {
		if pos.X < minX {
			minX = pos.X
		} else if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y < minY {
			minY = pos.Y
		} else if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	total := (maxX - minX + 1) * (maxY - minY + 1)
	fmt.Printf("%v\n", total-len(positions))

}

func modLikePython(d, m int) int {
	var res int = d % m
	if res < 0 && m > 0 {
		return res + m
	}
	return res
}
