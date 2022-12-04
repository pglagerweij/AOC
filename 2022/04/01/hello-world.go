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
	// fmt.Println(s3)
	total_count := 0
	for _, element := range s3 {
		splitted := strings.Split(element, ",")
		element_1 := splitted[0]
		element_2 := splitted[1]
		lowest_1, upper_1 := getElements(element_1)
		lowest_2, upper_2 := getElements(element_2)
		result := checkContains(lowest_1, upper_1, lowest_2, upper_2)
		if result {
			total_count += 1
		}
	}

	fmt.Println(total_count)

}

func checkContains(lowest_1 int, upper_1 int, lowest_2 int, upper_2 int) bool {

	if lowest_2 >= lowest_1 && upper_2 <= upper_1 {
		return true
	} else if lowest_1 >= lowest_2 && upper_1 <= upper_2 {
		return true
	}
	return false
}

func getElements(input string) (int, int) {
	inputArray := strings.Split(input, "-")
	lowest, _ := strconv.Atoi(inputArray[0])
	upper, _ := strconv.Atoi(inputArray[1])
	return lowest, upper
}
