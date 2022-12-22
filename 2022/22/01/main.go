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
			fmt.Printf("The current position is %v we are facing in direction %v we are walking %v steps\n", currentposition, directions[facing], steps)
			currentposition = walkoverGrid(currentposition, facing, steps)
		} else if s == "L" {
			facing = modLikePython(facing-1, 4)
			fmt.Printf("we are turning left new direction %v\n", directions[facing])
		} else if s == "R" {
			facing = modLikePython(facing+1, 4)
			fmt.Printf("we are turning right new direction %v\n", directions[facing])
		}

	}

	fmt.Printf("The final position is %v with direction %v\n", currentposition, directions[facing])
	fmt.Printf("Total score: %v\n", 1000*currentposition.X+4*currentposition.Y+facing)

}

func walkoverGrid(position Position, direction int, steps int) Position {
	// fmt.Printf("calling walkover grid function\n")
	var movement Position
	if direction == 0 {
		movement = Position{X: 0, Y: 1}
	} else if direction == 1 {
		movement = Position{X: 1, Y: 0}
	} else if direction == 2 {
		movement = Position{X: 0, Y: -1}
	} else if direction == 3 {
		movement = Position{X: -1, Y: 0}
	} else {
		panic("I dont know this direction")
	}
	fmt.Printf("we are moving from position %v with this movement %v for steps %v\n", position, movement, steps)
	return settingsteps(position, movement, steps)
}

func settingsteps(pos Position, movement Position, steps int) Position {
	startpos := pos
	for step := 0; step < steps; step++ {
		newposition := pos.Add(movement)
		value, ok := grid[newposition]
		if ok {
			if value == 1 { // BLocked by #
				return pos
			} else if value == 0 {
				pos = newposition
				continue
			}
		} else {
			backwardpos := startpos.Min(movement)
			_, ok2 := grid[backwardpos]
			for ok2 {
				backwardpos = backwardpos.Min(movement)
				_, ok2 = grid[backwardpos]
			}
			firsteleminrow := backwardpos.Add(movement)
			valuenewelem := grid[firsteleminrow]
			if valuenewelem == 1 { // BLocked by #
				return pos
			} else if value == 0 {
				pos = firsteleminrow
				continue
			}
		}
	}

	return pos
}

func modLikePython(d, m int) int {
	var res int = d % m
	if res < 0 && m > 0 {
		return res + m
	}
	return res
}
