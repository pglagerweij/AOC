package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var inputFile string = "trial.txt"

var inputFile string = "input.txt"

var indexToCheck []int = []int{1000, 2000, 3000}

type totalSlice struct {
	location, value int
}

func main() {
	input, _ := os.ReadFile(inputFile)
	inputSlice := map[int]totalSlice{}
	iterationSlice := []totalSlice{}
	// convert the file binary into a string using string
	for ind, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		sint, _ := strconv.Atoi(strings.TrimSpace(s))
		inputSlice[ind] = totalSlice{ind, sint}
		iterationSlice = append(iterationSlice, totalSlice{ind, sint})
		// fmt.Printf("the value is sint %v\n", sint)
	}
	// for _, char := range iterationSlice {
	// 	fmt.Printf("%v ", char.value)
	// }
	// fmt.Printf("\n")
	totalLength := len(iterationSlice)
	for ind := 0; ind < totalLength; ind++ {
		iterationSlice = updatSlice(iterationSlice, inputSlice[ind], totalLength)
	}

	indexNull := 0
	for ind, char := range iterationSlice {
		// fmt.Printf("%v ", char.value)
		if char.value == 0 {
			indexNull = ind
		}
	}
	fmt.Printf("\n")
	totalSum := 0
	for _, index := range indexToCheck {
		fmt.Printf("The index null is at %v\n", indexNull)
		relevantIndex := (index + indexNull) % totalLength
		number := iterationSlice[relevantIndex].value
		totalSum += number
		fmt.Printf("%v gives index %v with number %v\n", index, (index+indexNull)%totalLength, number)
	}
	fmt.Printf("The total sum is %v\n", totalSum)

}

func updatSlice(iterationSlice []totalSlice, inputelem totalSlice, totalLength int) []totalSlice {
	for index, adjustedElem := range iterationSlice {
		if inputelem == adjustedElem {
			stepsToMove := inputelem.value % totalLength
			newIndex := modLikePython(index+stepsToMove, totalLength)
			// fmt.Printf("The element %v is now at index %v and we are moving it %v and is moved to index %v\n", adjustedElem.value, index, stepsToMove, newIndex)
			if newIndex == 0 {
				iterationSlice = insert(RemoveIndex(iterationSlice, index), totalLength-1, adjustedElem)
			} else if index < newIndex {
				iterationSlice = insert(RemoveIndex(iterationSlice, index), newIndex, adjustedElem)
			} else if index > newIndex {
				iterationSlice = RemoveIndex(insert(iterationSlice, newIndex, adjustedElem), index+1)
			}

			// for _, char := range iterationSlice {
			// 	fmt.Printf("%v ", char.value)
			// }
			// fmt.Printf("\n")
			return iterationSlice
		}

	}

	panic("cannot find inputelem")
}

// 0 <= index <= len(a)
func insert(a []totalSlice, index int, value totalSlice) []totalSlice {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func RemoveIndex(s []totalSlice, index int) []totalSlice {
	ret := make([]totalSlice, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// func modLikePython(d, m int) int {
// 	var res int = d % m
// 	if res < 0 && m > 0 {
// 		return res + m
// 	}
// 	return res
// }

func modLikePython(d, m int) int {
	if d < m && d > 0 {
		return d
	} else if d%m == 0 && d < 2*m {
		return 0
	} else if (d%m < 0) && (m > 0) && (-d <= m) {
		return d%m + m - 1
	} else if d%m > 0 && d < 2*m {
		return d%m + 1
	} else {
		panic("help")
	}

}
