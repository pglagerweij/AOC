package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	s := strings.ReplaceAll(fileContent, "\n\n", ";")
	s3 := strings.Split(s, ";")
	totalAssessment := make([]int, len(s3))
	for index := range totalAssessment {
		totalAssessment[index] = 0
	}
	// Initializ starting values
	currentValues := make([][]int, len(s3))
	for _, element := range s3 {
		monkeyIndex, startingNumbers, _, _, _, _ := extractInstructions(element)
		currentValues[monkeyIndex] = startingNumbers
	}

	// Go through the rounds
	for roundNumber := 1; roundNumber <= 1000; roundNumber++ {
		for _, element := range s3 {
			monkeyIndex, _, instructions, divNumber, trueMonkey, falseMonkey := extractInstructions(element)
			// fmt.Printf("For Round %v Monkey %v and current values %v with instructions %v, test is to divide by %v. If true throw to %v if false throw to %v.\n", roundNumber, monkeyIndex, currentValues[monkeyIndex], instructions, divNumber, trueMonkey, falseMonkey)

			// Perform operation for all values
			for _, currentVal := range currentValues[monkeyIndex] {
				totalAssessment[monkeyIndex] += 1
				newNumber := performOperation(instructions, currentVal, 3)
				result, numberToWrite := determineResult(newNumber, divNumber)
				// fmt.Printf("Dividing %v with %v result is %v.\n", newNumber, divNumber, result)
				if result == true {
					currentValues[trueMonkey] = append(currentValues[trueMonkey], numberToWrite)
				} else if result == false {
					currentValues[falseMonkey] = append(currentValues[falseMonkey], numberToWrite)
				}
				currentValues[monkeyIndex] = nil
			}

		}
		fmt.Printf("Current values after round %v are %v.\n", roundNumber, currentValues)
		fmt.Printf("total assesments after round %v are %v.\n", roundNumber, totalAssessment)

	}
	// get Largest 2 numbers
	var large1 int = 0
	var large2 int = 0
	large1 = totalAssessment[0]
	for i := 1; i <= len(totalAssessment)-1; i++ {
		if large1 < totalAssessment[i] {
			large2 = large1
			large1 = totalAssessment[i]
		} else if large2 < totalAssessment[i] {
			large2 = totalAssessment[i]
		}
	}
	fmt.Printf("total score after rounds is %v.\n", large2*large1)
}
func determineResult(number int, divnumber int) (bool, int) {
	if number%divnumber == 0 {
		return true, number % 96577
	} else {
		return false, number % 96577
	}
}

func performOperation(operation string, inputNumber int, divisionNumber int) int {
	if operation == "new = old * 3" {
		return int(math.Floor(float64((inputNumber * 3)) / float64(divisionNumber)))
	} else if operation == "new = old * 7" {
		return int(math.Floor(float64((inputNumber * 7)) / float64(divisionNumber)))
	} else if operation == "new = old + 8" {
		return int(math.Floor(float64((inputNumber + 8)) / float64(divisionNumber)))
	} else if operation == "new = old + 5" {
		return int(math.Floor(float64((inputNumber + 5)) / float64(divisionNumber)))
	} else if operation == "new = old * old" {
		return int(math.Floor(float64((inputNumber * inputNumber)) / float64(divisionNumber)))
	} else if operation == "new = old + 4" {
		return int(math.Floor(float64((inputNumber + 4)) / float64(divisionNumber)))
	} else if operation == "new = old + 3" {
		return int(math.Floor(float64((inputNumber + 3)) / float64(divisionNumber)))
	} else if operation == "new = old + 2" {
		return int(math.Floor(float64((inputNumber + 2)) / float64(divisionNumber)))
	} else if operation == "new = old + 1" {
		return int(math.Floor(float64((inputNumber + 1)) / float64(divisionNumber)))
	} else {
		panic("operation not found " + operation)
	}
}

func extractInstructions(instructions string) (int, []int, string, int, int, int) {
	instructionsArray := strings.Split(instructions, "\n")

	// Get monkey number
	monkeyLine := instructionsArray[0]
	monkeyNumber, _ := strconv.Atoi(string(monkeyLine[7]))
	// Get starting numbers
	startingInstructions := strings.TrimSpace(instructionsArray[1])
	numbersString := strings.ReplaceAll(startingInstructions, "Starting items: ", "")
	numbersArray := strings.Split(numbersString, ", ")
	numbersArrayInt, _ := sliceAtoi(numbersArray)

	// Get instructions
	operationInstructions := strings.TrimSpace(instructionsArray[2])
	operationsString := strings.ReplaceAll(operationInstructions, "Operation: ", "")

	// Get divisibleInstructions
	divisblestring := strings.TrimSpace(instructionsArray[3])
	divString := strings.ReplaceAll(divisblestring, "Test: divisible by ", "")
	divisibleNumber, _ := strconv.Atoi(divString)

	// Get trueNUmbber
	truestring := strings.TrimSpace(instructionsArray[4])
	truetrimmedString := strings.ReplaceAll(truestring, "If true: throw to monkey ", "")
	trueNumber, _ := strconv.Atoi(truetrimmedString)

	// Get trueNUmbber
	falsestring := strings.TrimSpace(instructionsArray[5])
	falsetrimmedString := strings.ReplaceAll(falsestring, "If false: throw to monkey ", "")
	falseNumber, _ := strconv.Atoi(falsetrimmedString)

	return monkeyNumber, numbersArrayInt, operationsString, divisibleNumber, trueNumber, falseNumber
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
