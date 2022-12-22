package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputFile string = "trial.txt"

// var inputFile string = "input.txt"

var directions map[int]string = map[int]string{
	0: "Right",
	1: "Down",
	2: "Left",
	3: "Up",
}

func main() {
	input, _ := os.ReadFile(inputFile)
	total := strings.Split(string(input), "\n\n")
	drawing := total[0]
	instructions := total[1]
	currentPosition := 0
	// Construct Map element
	for _, row := range strings.Split(drawing, "\n") {
		fmt.Printf("%v\n", row)
	}

	// Loop over instructions
	for _, s := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(instructions, "R", ";R;"), "L", ";L;"), ";") {
		if s == "L" {
			currentPosition = modLikePython(currentPosition-1, 4)
			continue
		} else if s == "R" {
			currentPosition = modLikePython(currentPosition+1, 4)
			continue
		}

		sint, _ := strconv.Atoi(s)
		fmt.Printf("The value is %v\n", sint)

	}

}

func modLikePython(d, m int) int {
	var res int = d % m
	if res < 0 && m > 0 {
		return res + m
	}
	return res
}
