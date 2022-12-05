package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var inputFile string = "input.txt"

// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

func main() {
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	// print file content
	// fmt.Println(fileContent)
	totalInput := strings.Split(strings.ReplaceAll(fileContent, "\n\n", ";"), ";")

	// puzzleState := totalInput[0]
	// var startState [3][]rune
	// startState[0] = []rune{'Z', 'N'}
	// startState[1] = []rune{'M', 'C', 'D'}
	// startState[2] = []rune{'P'}
	var startState [9][]rune
	startState[0] = []rune{'S', 'M', 'R', 'N', 'W', 'J', 'V', 'T'}
	startState[1] = []rune{'B', 'W', 'D', 'J', 'Q', 'P', 'C', 'V'}
	startState[2] = []rune{'B', 'J', 'F', 'H', 'D', 'R', 'P'}
	startState[3] = []rune{'F', 'R', 'P', 'B', 'M', 'N', 'D'}
	startState[4] = []rune{'H', 'V', 'R', 'P', 'T', 'B'}
	startState[5] = []rune{'C', 'B', 'P', 'T'}
	startState[6] = []rune{'B', 'J', 'R', 'P', 'L'}
	startState[7] = []rune{'N', 'C', 'S', 'L', 'T', 'Z', 'B', 'W'}
	startState[8] = []rune{'L', 'S', 'G'}
	fmt.Printf("%v\n", startState)

	// Get instructions
	puzzleInstructions := totalInput[1]
	s := strings.ReplaceAll(puzzleInstructions, "\n", ";")
	s3 := strings.Split(s, ";")
	fmt.Println(s3)
	currentState := startState
	for _, element := range s3 {
		fmt.Printf("%v\n", element)
		wordBreakDown := strings.Fields(element)
		numbertoMove, _ := strconv.Atoi(wordBreakDown[1])
		fromStack, _ := strconv.Atoi(wordBreakDown[3])
		toStack, _ := strconv.Atoi(wordBreakDown[5])
		// fmt.Printf("%v %v %v\n", numbertoMove, fromStack, toStack)
		currentState = updateState(currentState, numbertoMove, fromStack, toStack)
		fmt.Printf("%v\n", currentState)
	}
	for _, rows := range currentState {
		fmt.Printf("%v", string(rows[len(rows)-1]))
	}

	// currentState[0][]

}

func updateState(currentState [9][]rune, numbertoMove int, fromStack int, toStack int) [9][]rune {
	fromStackSlice := currentState[fromStack-1]
	stackToMove := fromStackSlice[len(fromStackSlice)-numbertoMove:]
	// fmt.Printf("%v\n", stackToMove)
	// ReverseSlice(stackToMove)
	// fmt.Printf("%v\n", stackToMove)
	currentState[toStack-1] = append(currentState[toStack-1], stackToMove...)
	currentState[fromStack-1] = fromStackSlice[:len(fromStackSlice)-numbertoMove]
	return currentState
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}
