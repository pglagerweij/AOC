package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var inputFile string = "input.txt"

func main() {
	totalSize := 99
	threedim := [99][]int{}
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	s := strings.ReplaceAll(fileContent, "\n", ";")
	s3 := strings.Split(s, ";")
	// currentIndex := 0
	for indexA, element := range s3 {
		// fmt.Printf("the element we are processing is %v with index %v\n", indexA, element)
		tempSlice := []int{}
		for _, character := range element {
			// fmt.Printf("the current charecter: %c\n", character)
			intChar, _ := strconv.Atoi(string(character))
			// threedim[indexA][indexB] = int(character)
			tempSlice = append(tempSlice, intChar)
		}
		threedim[indexA] = tempSlice

	}

	totalScore := 0
	for rowNumber := 1; rowNumber < totalSize-1; rowNumber++ {
		for columnNumber := 1; columnNumber < totalSize-1; columnNumber++ {
			sliceArrayLeft, sliceArrayRight, sliceArrayTop, sliceArrayBottom := getSlices(threedim, rowNumber, columnNumber)
			// fmt.Printf("%v is visible %v \n", sliceArrayRight, determineView(sliceArrayRight))
			sliceArrayLeftReverse := Reverse(sliceArrayLeft)
			// fmt.Printf("%v is visible %v \n", sliceArrayLeft, determineView(sliceArrayLeftReverse))
			// fmt.Printf("%v is visible %v \n", sliceArrayBottom, determineView(sliceArrayBottom))
			sliceArrayTopReverse := Reverse(sliceArrayTop)
			// fmt.Printf("%v is visible %v \n", sliceArrayTop, determineView(sliceArrayTopReverse))
			viewLeft := determineView(sliceArrayLeftReverse)
			viewRight := determineView(sliceArrayRight)
			viewDown := determineView(sliceArrayTopReverse)
			viewUp := determineView(sliceArrayBottom)
			score := viewLeft * viewRight * viewDown * viewUp
			if score > totalScore {
				totalScore = score
			}
		}

	}

	fmt.Printf("The total maximum score is %v", totalScore)
}
func determineView(heightSlice []int) int {
	var firstNumber int = heightSlice[0]
	viewPoints := 0
	for _, char := range heightSlice[1:] {
		if char < firstNumber {
			viewPoints += 1
		} else if char == firstNumber {
			viewPoints += 1
			return viewPoints
		} else {
			viewPoints += 1
			return viewPoints
		}
	}
	return viewPoints
}

func determineVisibility(heightSlice []int) bool {
	// fmt.Printf("%v", heightSlice)
	var lastnumber int = heightSlice[len(heightSlice)-1]
	// fmt.Printf("%v", lastnumber)
	for _, char := range heightSlice[:len(heightSlice)-1] {
		if char >= lastnumber {
			return false
		}
	}
	return true
}

func getSlices(board [99][]int, rowIndex int, columnIndex int) ([]int, []int, []int, []int) {
	sliceArrayLeft := board[rowIndex][:columnIndex+1]
	sliceArrayRight := board[rowIndex][columnIndex:]
	sliceArrayColumn := boardColumn(board, columnIndex)
	sliceArrayTop := sliceArrayColumn[:rowIndex+1]
	sliceArrayBottom := sliceArrayColumn[rowIndex:]
	return sliceArrayLeft, sliceArrayRight, sliceArrayTop, sliceArrayBottom
}

func boardColumn(board [99][]int, columnIndex int) (column []int) {
	column = make([]int, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return
}

func reverse[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func Reverse[T any](original []T) (reversed []T) {
	reversed = make([]T, len(original))
	copy(reversed, original)

	for i := len(reversed)/2 - 1; i >= 0; i-- {
		tmp := len(reversed) - 1 - i
		reversed[i], reversed[tmp] = reversed[tmp], reversed[i]
	}

	return
}
