package main

import (
	"fmt"
	"io/ioutil"
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
	for index := 0; index < len(fileContent)-3; index++ {
		stringToCheck := fileContent[index : index+4]
		if unique(stringToCheck) {
			fmt.Printf("unique string is %v on index %v\n", stringToCheck, index+4)
			break
		}

	}
}

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}
