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
	var window = []int{}
	for index := 0; index < len(s3)-2; index++ {
		total := s3[index : index+3]
		// fmt.Printf("%v", sum(total))
		sumOfArray := sum(total)
		window = append(window, sumOfArray)
		// fmt.Printf("%v\n", total)
	}
	// fmt.Printf("%v\n", window)
	var lastElement int = 999999999
	var totalIncrease int = 0
	for _, element := range window {
		// currentElement, _ := strconv.Atoi(element)
		if element > lastElement {
			totalIncrease += 1
		}
		lastElement = element
	}

	fmt.Printf("maximum value is: %v", totalIncrease)
}

func sum(array []string) int {
	result := 0
	for _, v := range array {
		temp, _ := strconv.Atoi(v)
		result += temp
	}
	return result
}
