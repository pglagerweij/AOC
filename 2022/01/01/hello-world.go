package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
	s := strings.ReplaceAll(fileContent, "\n\n", ";")
	s2 := strings.ReplaceAll(s, "\n", ",")
	s3 := strings.Split(s2, ";")

	var totals []uint64
	for _, element := range s3 {
		array := strings.Split(element, ",")
		var result uint64 = 0
		for _, i := range array {
			j, err := strconv.ParseUint(i, 10, 64)
			if err != nil {
				panic(err)
			}
			result += j
		}
		totals = append(totals, result)
	}
	fmt.Printf("%v\n", totals)

	// Get maximum of totals Slice
	var m uint64
	for i, e := range totals {
		if i == 0 || e > m {
			m = e
		}
	}

	fmt.Printf("maximum value is: %v", m)
}
