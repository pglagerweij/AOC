package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile string = "input.txt"

func main() {

	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	s := strings.ReplaceAll(fileContent, "\n", ";")
	s3 := strings.Split(s, ";")

	wordBreakDown := make([][]rune, len(s3))
	var rowLength int
	var startIndex []int
	var endIndex []int
	for indexA, element := range s3 {
		rowLength = len(element)
		for index := range element {
			wordBreakDown[indexA] = append(wordBreakDown[indexA], rune(element[index]))
			if element[index] == 83 {
				wordBreakDown[indexA][index] = 97
				startIndex = []int{indexA, index}
			} else if element[index] == 69 {
				wordBreakDown[indexA][index] = 122
				endIndex = []int{indexA, index}
			}
		}
	}
	fmt.Printf("%v\n", rowLength)
	fmt.Printf("%v\n", startIndex)
	fmt.Printf("%v\n", endIndex)

	// Initilaze array
	stepsBreakDown := make([][]int, len(s3))
	for i := 0; i < len(stepsBreakDown); i++ {
		stepsBreakDown[i] = make([]int, rowLength)
	}

	for steps := 0; stepsBreakDown[endIndex[0]][endIndex[1]] == 0; steps++ {
		// for steps := 1; steps < 30; steps++ {
		locationList := [][]int{}
		if steps == 1 {
			locationList = append(locationList, startIndex)
		} else {
			locationList = getLocationList(stepsBreakDown, steps-1)
		}
		// fmt.Printf("%v\n", locationList)
		for _, location := range locationList {
			stepsBreakDown = updateStepBreakdown(stepsBreakDown, wordBreakDown, location, steps, rowLength)
		}

		// fmt.Printf("after %v the %v\n", steps, stepsBreakDown)
	}

	fmt.Printf("The final end stage step break down is %v\n", stepsBreakDown[endIndex[0]][endIndex[1]])
}

func updateStepBreakdown(stepsBreakDown [][]int, wordBreakDown [][]rune, location []int, steps int, rowLength int) [][]int {
	startRowIndex := location[0]
	startColumnIndex := location[1]
	startWord := wordBreakDown[startRowIndex][startColumnIndex]

	// Go down
	if startRowIndex != len(stepsBreakDown)-1 {
		if stepsBreakDown[startRowIndex+1][startColumnIndex] == 0 {
			compareWordDown := wordBreakDown[startRowIndex+1][startColumnIndex]
			// fmt.Printf("start word %c and down word %c\n", startWord, compareWordDown)
			if compareWordDown-startWord <= 1 {
				stepsBreakDown[startRowIndex+1][startColumnIndex] = steps
			}
		}
	}
	// // Go up
	if startRowIndex != 0 {
		if stepsBreakDown[startRowIndex-1][startColumnIndex] == 0 {
			compareWordUp := wordBreakDown[startRowIndex-1][startColumnIndex]
			// fmt.Printf("start word %c and down word %c\n", startWord, compareWordUp)
			if compareWordUp-startWord <= 1 {
				stepsBreakDown[startRowIndex-1][startColumnIndex] = steps
			}
		}
	}

	// // Go Left
	if startColumnIndex != 0 {
		if stepsBreakDown[startRowIndex][startColumnIndex-1] == 0 {
			compareWordLeft := wordBreakDown[startRowIndex][startColumnIndex-1]
			// fmt.Printf("start word %c and down word %c\n", startWord, compareWordLeft)
			if compareWordLeft-startWord <= 1 {
				stepsBreakDown[startRowIndex][startColumnIndex-1] = steps
			}
		}
	}

	// Go Right
	if startColumnIndex != rowLength-1 {
		if stepsBreakDown[startRowIndex][startColumnIndex+1] == 0 {
			compareWordRight := wordBreakDown[startRowIndex][startColumnIndex+1]
			// fmt.Printf("start word %c and down word %c\n", startWord, compareWordRight)
			if compareWordRight-startWord <= 1 {
				stepsBreakDown[startRowIndex][startColumnIndex+1] = steps
			}
		}
	}
	return stepsBreakDown
}

func getLocationList(steps [][]int, numberOfSteps int) [][]int {
	locList := [][]int{}
	for indexA, element := range steps {
		for index := range element {
			// fmt.Printf("%v\n", steps[indexA][index])
			if steps[indexA][index] == numberOfSteps {
				locList = append(locList, []int{indexA, index})
			}
		}
	}

	return locList
}
