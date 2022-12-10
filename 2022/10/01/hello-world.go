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
	s := strings.ReplaceAll(fileContent, "\n", ";")
	s3 := strings.Split(s, ";")
	currentIndex := 0
	valueOverTime := []int{1}
	for _, element := range s3 {
		if strings.HasPrefix(element, "noop") {
			// fmt.Printf("%v\n", element)
			currentIndex += 1
			valueOverTime = append(valueOverTime, valueOverTime[currentIndex-1])
			// fmt.Println(currentPath)
		} else if strings.HasPrefix(element, "addx ") {
			splitElement := strings.Split(element, " ")
			value, _ := strconv.Atoi(strings.TrimSpace(splitElement[1]))
			currentIndex += 1
			valueOverTime = append(valueOverTime, valueOverTime[currentIndex-1])
			currentIndex += 1
			valueOverTime = append(valueOverTime, valueOverTime[currentIndex-1]+value)
			// fmt.Printf("next ls command: %v\n", element)
			// allFiles = fillDirectory(allFiles, element, currentPath)
		}

	}
	totalOfSum := 0
	for i := 19; i < 221; i += 40 {
		totalOfSum += valueOverTime[i] * (i + 1)
		fmt.Printf("The index %v gives value %v\n", i+1, valueOverTime[i]*(i+1))
	}

	// fmt.Printf("The total visible trees are %v", valueOverTime)
	fmt.Printf("The total visible trees are %v", totalOfSum)
}
