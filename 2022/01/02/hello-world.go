package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	var totals []int
	for _, element := range s3 {
		array := strings.Split(element, ",")
		var result int = 0
		for _, i := range array {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			result += j
		}
		totals = append(totals, result)
	}
	fmt.Printf("%v\n", totals)

	sorted_slice := totals

	sort.Slice(sorted_slice, func(i, j int) bool {
		return sorted_slice[i] > sorted_slice[j]
	})
	total_first_3 := sorted_slice[0] + sorted_slice[1] + sorted_slice[2]
	fmt.Printf("Total value is: %v", total_first_3)
}
