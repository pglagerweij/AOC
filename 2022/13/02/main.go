package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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
	s := strings.ReplaceAll(fileContent, "\n\n", "\n")
	s3 := strings.Split(s, "\n")
	s3 = append(s3, "[[2]]")
	s3 = append(s3, "[[6]]")
	orderedElements := []string{}
	for _, elementToAdd := range s3 {
		fmt.Printf("checking element %v\n", elementToAdd)
		if len(orderedElements) == 0 {
			orderedElements = append(orderedElements, elementToAdd)
		} else {
			for index, compareElement := range orderedElements {
				fmt.Printf("Checking compared to element %v\n", compareElement)
				resultTotal := compareElementss2(compareElement, elementToAdd)
				fmt.Printf("The compared elements have result are %v\n", resultTotal)
				if resultTotal == "left" {
					orderedElements = insert(orderedElements, index, elementToAdd)
					break
				} else if resultTotal == "right" && (index == len(orderedElements)-1) {
					orderedElements = append(orderedElements, elementToAdd)
				}

			}
			// fmt.Printf("The compared elements have result are %v\n", orderedElements)
		}
	}

	// Sum items for final answer
	var index1 int
	var index2 int
	for index, element := range orderedElements {
		if element == "[[2]]" {
			index1 = index + 1
		} else if element == "[[6]]" {
			index2 = index + 1
		}
	}
	fmt.Printf("The final answer is %v\n", index1*index2)
	// fmt.Printf("The final sum of indexes is %v\n", sumSlice(largerIndex))
}

func insert(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func compareElementss2(left string, right string) string {
	// fmt.Printf("Compare %v vs %v\n", left, right)
	if (left == "") || (left == "[]") {
		fmt.Printf("Left side ran out of items, so inputs are in the right order\n")
		return "right"
	} else if (right == "") || (right == "[]") {
		fmt.Printf("Right side ran out of items, so inputs are not in the right order\n")
		return "left"
	}
	leftElement, restLeft := getElement(left)
	rightElement, restRight := getElement(right)
	// fmt.Printf("Compare %v vs %v\n", leftElement, rightElement)
	if leftElement == rightElement {
		result := compareElementss2(restLeft, restRight)
		return result
	} else if (leftElement[0] != '[') && rightElement[0] != '[' {
		leftint, _ := strconv.Atoi(leftElement)
		rightint, _ := strconv.Atoi(rightElement)
		if leftint < rightint {
			fmt.Printf("Left side is smaller, so inputs are in the right order\n")
			return "right"
		} else if leftint > rightint {
			fmt.Printf("Right side is smaller, so inputs are not in the right order\n")
			return "left"
		} else {
			panic("left and right didnt match" + leftElement + " and " + rightElement)
		}
	} else if (leftElement[0] != '[') && rightElement[0] == '[' {
		// fmt.Printf("Mixed types; convert left to %v and retry comparison\n", "["+leftElement+"]")
		result := compareElementss2("["+leftElement+"]", rightElement)
		return result
	} else if (leftElement[0] == '[') && rightElement[0] != '[' {
		// fmt.Printf("Mixed types; convert right to %v and retry comparison\n", "["+rightElement+"]")
		result := compareElementss2(leftElement, "["+rightElement+"]")
		return result
	} else if (rightElement == "[]") || (leftElement == "[]") || (leftElement[0] == '[' && rightElement[0] == '[') {
		result := compareElementss2(leftElement, rightElement)
		return result
	}
	panic("no comparision found for" + left + " and " + right)
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

func sumSlice(numarray []int) int {
	arrSum := 0

	for i := 0; i < len(numarray); i++ {
		arrSum = arrSum + numarray[i]
	}
	return arrSum
}
