package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile string = "trial.txt"

func main() {

	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	s := strings.ReplaceAll(fileContent, "\n\n", ";")
	s3 := strings.Split(s, ";")
	largerIndex := []int{}
	for index, element := range s3 {
		left := strings.Split(element, "\n")[0]
		right := strings.Split(element, "\n")[1]
		// fmt.Printf("comparing %v with %v\n", left, right)
		resultTotal := compareElementss(left, right)
		if resultTotal == "right" {
			largerIndex = append(largerIndex, index+1)
		}
		fmt.Printf("comparing %v with %v result is %v\n", left, right, resultTotal)
	}
	// leftElement, restLeft := getElement(left)
	// rightElement, restRight := getElement(right)
	// fmt.Printf("comparing %v with %v\n", leftElement, rightElement)
	// if leftElement == rightElement {
	// 	fmt.Printf("comparing %v with %v\n", restLeft, restRight)
	// 	leftElement2, restLeft2 := getElement(restLeft)
	// 	rightElement2, restRight2 := getElement(restRight)
	// 	fmt.Printf("comparing %v with %v\n", leftElement2, rightElement2)
	// 	fmt.Printf("leftovers %v with %v\n", restLeft2, restRight2)
	// 	// result := compareElement(restLeft, restRight)
	// 	// return result
	// } else {
	// 	fmt.Printf("really comparing now")
	// }

	// largerIndex := []int{}
	// for index, element := range s3 {
	// 	elementLeft := strings.Split(element, "\n")[0]
	// 	elementRight := strings.Split(element, "\n")[1]
	// 	result := compareElementss(elementLeft, elementLeft)
	// 	if result == "right" {
	// 		largerIndex = append(largerIndex, index+1)
	// 	}
	// 	fmt.Printf("On index %v comparing %v with %v\n", index+1, elementLeft, elementRight)

	// }

	// Sum items for final answer
	fmt.Printf("The final indexes are %v\n", largerIndex)

	fmt.Printf("The final sum of indexes is %v\n", sumSlice(largerIndex))
}

func isLeftEarly(leftInput string, rightInput string) string {
	if (leftInput[0] != '[') && (rightInput[0] != '[') {
		leftint, _ := strconv.Atoi(leftInput)
		rightint, _ := strconv.Atoi(rightInput)
		if leftint < rightint {
			return "right"
		} else if leftint > rightint {
			return "left"
		} else {
			panic("left and right didnt match" + leftInput + " and " + rightInput)
		}
	} else {
		return "stillhavetobuildcomparison"
	}
}

func compareElementss(left string, right string) string {
	fmt.Printf("comparing %v with %v\n", left, right)
	leftElement, restLeft := getElement(left)
	rightElement, restRight := getElement(right)
	fmt.Printf("comparing main argument %v with %v\n", leftElement, rightElement)
	if leftElement == rightElement {
		result := compareElementss(restLeft, restRight)
		return result
	} else if leftElement == "" {
		return "right"
	} else if rightElement == "" {
		return "left"
	} else if (leftElement[0] != '[') && rightElement[0] == '[' {
		fmt.Printf("converting left element to list %v\n", leftElement)
		result := compareElementss("["+leftElement+"]", rightElement)
		return result
	} else if (leftElement[0] == '[') && rightElement[0] != '[' {
		fmt.Printf("converting right element to list: %v\n", rightElement)
		result := compareElementss(leftElement, "["+rightElement+"]")
		return result
	} else {
		fmt.Printf("really comparing now element %v and %v\n", leftElement, rightElement)
		result := isLeftEarly(leftElement, rightElement)
		return result
	}
}

func getElement(input string) (string, string) {
	bracketCounter := 1
	var leftBracket byte = '['
	var rightBracket byte = ']'
	for index := 1; index <= len(input)-1; index++ {
		if input[index] == leftBracket {
			bracketCounter += 1
			// fmt.Printf("bracket counter went up %v\n", bracketCounter)
		} else if input[index] == rightBracket {
			bracketCounter -= 1
			// fmt.Printf("bracket counter went down %v\n", bracketCounter)
		}

		if bracketCounter == 1 && input[index] == ',' {
			// fmt.Printf("returning string  %v\n", input[1:index])
			return input[1:index], "[" + input[index+1:]
		} else if bracketCounter == 0 {
			// fmt.Printf("returning string  %v\n", input[1:index])
			return input[1:index], ""
		}

	}
	return "[]", ""
	// fmt.Printf("hi input is %v", input)
	// panic("didnt find match on input" + input)
}

func compareElements(left string, right string) bool {
	return true
}

func sumSlice(numarray []int) int {
	arrSum := 0

	for i := 0; i < len(numarray); i++ {
		arrSum = arrSum + numarray[i]
	}
	return arrSum
}
