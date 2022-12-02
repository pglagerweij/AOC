package main

import (
	"fmt"
	"io/ioutil"
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
	s2 := strings.ReplaceAll(s, " ", ",")
	s3 := strings.Split(s2, ";")
	// fmt.Println(s3)
	var total_score int = 0
	for _, element := range s3 {
		array := strings.Split(element, ",")
		score := calculateScore(array)
		// fmt.Printf("%v\n", score)
		total_score += score
	}
	fmt.Printf("The total score is: %v", total_score)
}

func calculateScore(array []string) int {
	scoring_table := map[[2]string]int{
		{"A", "X"}: 3, // Rock vs Sciccors - lose (0+3)
		{"A", "Y"}: 4, // Rock vs Rock  - draw (3+1)
		{"A", "Z"}: 8, // Rock vs paper - win (6+2)
		{"B", "X"}: 1, // Paper vs Rock  - lose (0+1)
		{"B", "Y"}: 5, // Paper vs Paper - draw (3+2)
		{"B", "Z"}: 9, // Paper vs Sciccors -  win (6+3)
		{"C", "X"}: 2, // Sciccors vs paper  - lose (0+2)
		{"C", "Y"}: 6, // Sciccors vs Sciccors -  draw (3+3)
		{"C", "Z"}: 7, // Sciccors vs Rock  - win (6+1)
	}
	var score int = 0
	var yourself string = array[1]
	var opponent string = array[0]

	var total_thing [2]string
	total_thing[0] = opponent
	total_thing[1] = yourself

	value, found := scoring_table[total_thing]
	if found == false {
		panic("error")
	}
	score += value

	return score
}
