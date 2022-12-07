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
	// print file content
	// fmt.Println(fileContent)
	s := strings.ReplaceAll(fileContent, "\n", ";")
	s3 := strings.Split(s, ";")
	var lastElement int = 173
	var totalIncrease int = 0
	for _, element := range s3 {
		currentElement, _ := strconv.Atoi(element)
		if currentElement > lastElement {
			totalIncrease += 1
		}
		lastElement = currentElement
	}

	fmt.Printf("maximum value is: %v", totalIncrease)
}
