package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var inputFile string = "trial.txt"

var inputFile string = "input.txt"

var directions map[int]string = map[int]string{
	0: "Right",
	1: "Down",
	2: "Left",
	3: "Up",
}

var Movement map[int]Position = map[int]Position{
	0: {X: 0, Y: 1},
	1: {X: 1, Y: 0},
	2: {X: 0, Y: -1},
	3: {X: -1, Y: 0},
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

var grid = map[Position]int{}

func main() {
	input, _ := os.ReadFile(inputFile)
	total := strings.Split(string(input), "\n\n")
	drawing := total[0]
	instructions := total[1]
	facing := 0
	// Construct Map element
	currentposition := Position{X: 1, Y: 999999999999999}
	for index, row := range strings.Split(drawing, "\n") {
		for charnumber, elem := range row {
			if elem == '.' {
				if charnumber+1 < currentposition.Y && index == 0 {
					currentposition.Y = charnumber + 1
				}
				grid[Position{X: index + 1, Y: charnumber + 1}] = 0
			} else if elem == '#' {
				grid[Position{X: index + 1, Y: charnumber + 1}] = 1
			}
		}
	}

	// Loop over instructions
	for _, s := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(instructions, "R", ";R;"), "L", ";L;"), ";") {
		if steps, err := strconv.Atoi(s); err == nil {
			// fmt.Printf("The current position is %v we are facing in direction %v we are walking %v steps\n", currentposition, directions[facing], steps)
			currentposition, facing = walkoverGrid(currentposition, facing, steps)
			// fmt.Printf("We ended up at position %v\n", currentposition)
		} else if s == "L" {
			facing = modLikePython(facing-1, 4)
			// fmt.Printf("we are turning left new direction %v\n", directions[facing])
		} else if s == "R" {
			facing = modLikePython(facing+1, 4)
			// fmt.Printf("we are turning right new direction %v\n", directions[facing])
		}
		fmt.Printf("for step %v\nEnding position is %v facing %v\n", s, currentposition, facing)

	}

	fmt.Printf("The final position is %v with direction %v\n", currentposition, directions[facing])
	fmt.Printf("Total score: %v\n", 1000*currentposition.X+4*currentposition.Y+facing)

}

func walkoverGrid(position Position, direction int, steps int) (Position, int) {
	// fmt.Printf("calling walkover grid function\n")
	movement := Movement[direction]
	// fmt.Printf("we are moving from position %v with this movement %v for steps %v\n", position, movement, steps)
	newposition, stepsleft := settingsteps(position, movement, steps)
	for stepsleft != 0 {
		var testposition Position
		var newdirection int
		if newposition.X >= 1 && newposition.X <= 50 && newposition.Y == 51 && direction == 2 { // 1 left towards 5 right
			testposition = Position{X: 151 - newposition.X, Y: 1}
			newdirection = 0
		} else if newposition.X >= 151 && newposition.X <= 200 && newposition.Y == 50 && direction == 0 { // 6 right towards 4 up
			testposition = Position{X: 150, Y: newposition.X - 100}
			newdirection = 3
		} else if newposition.X == 150 && newposition.Y >= 51 && newposition.Y <= 100 && direction == 1 { // 4 down towards 6 left
			testposition = Position{X: 100 + (newposition.Y), Y: 50}
			newdirection = 2
		} else if newposition.Y == 1 && newposition.X >= 151 && newposition.X <= 200 && direction == 2 { //  6 left towards 1 down
			testposition = Position{X: 1, Y: newposition.X - 100}
			newdirection = 1
		} else if newposition.X == 1 && newposition.Y >= 101 && newposition.Y <= 150 && direction == 3 { // 2 up towards 6 up
			testposition = Position{X: 200, Y: newposition.Y - 100}
			newdirection = 3
		} else if newposition.X == 200 && newposition.Y >= 1 && newposition.Y <= 50 && direction == 1 { // 6 down towards 2 down
			testposition = Position{X: 1, Y: newposition.Y + 100}
			newdirection = 1
		} else if newposition.X == 1 && newposition.Y >= 51 && newposition.Y <= 100 && direction == 3 { // 1 up towards 6 right
			testposition = Position{X: 100 + newposition.Y, Y: 1}
			newdirection = 0
		} else if newposition.Y == 150 && newposition.X >= 1 && newposition.X <= 50 && direction == 0 { // 2 right towards 4 left
			testposition = Position{Y: 100, X: 151 - newposition.X}
			newdirection = 2
		} else if newposition.Y == 100 && newposition.X >= 101 && newposition.X <= 150 && direction == 0 { // 4 right towards 2 left
			testposition = Position{Y: 150, X: 151 - newposition.X}
			newdirection = 2
		} else if newposition.Y == 100 && newposition.X >= 51 && newposition.X <= 100 && direction == 0 { // 3 right towards 2 up
			testposition = Position{X: 50, Y: newposition.X + 50}
			newdirection = 3
		} else if newposition.X == 50 && newposition.Y >= 101 && newposition.Y <= 150 && direction == 1 { // 2 down towards 3 left
			testposition = Position{Y: 100, X: newposition.Y - 50}
			newdirection = 2
		} else if newposition.Y == 1 && newposition.X >= 101 && newposition.X <= 150 && direction == 2 { // 5 left towards 1 right
			testposition = Position{Y: 51, X: 151 - newposition.X}
			newdirection = 0
		} else if newposition.Y == 51 && newposition.X >= 51 && newposition.X <= 100 && direction == 2 { // 3 left towards 5 down
			testposition = Position{X: 101, Y: newposition.X - 50}
			newdirection = 1
		} else if newposition.X == 101 && newposition.Y >= 1 && newposition.Y <= 50 && direction == 3 { // 5 up towards 3 right
			testposition = Position{Y: 51, X: newposition.Y + 50}
			newdirection = 0
		} else {

			fmt.Printf("we still have to do rotation on a cube we are at %v and move %v %v steps!\n", newposition, directions[direction], stepsleft)
			os.Exit(3)
		}

		if grid[testposition] == 1 {
			// fmt.Printf("We have some things with blocking cube")
			return newposition, direction
		}
		direction = newdirection
		newposition, stepsleft = settingsteps(testposition, Movement[newdirection], stepsleft-1)

	}
	return newposition, direction
}

func settingsteps(pos Position, movement Position, steps int) (Position, int) {
	for step := 0; step < steps; step++ {
		newposition := pos.Add(movement)
		value, ok := grid[newposition]
		if ok {
			if value == 1 { // BLocked by #
				return pos, 0
			} else if value == 0 {
				pos = newposition
				continue
			}
		} else {
			return pos, steps - step
		}
	}

	return pos, 0
}

func modLikePython(d, m int) int {
	var res int = d % m
	if res < 0 && m > 0 {
		return res + m
	}
	return res
}
