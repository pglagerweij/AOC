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
		[2]string{"A", "X"}: 3, // Rock vs Rock - draw
		[2]string{"A", "Y"}: 6, // Rock vs Paper - win
		[2]string{"A", "Z"}: 0, // Rock vs Sciccors - lose
		[2]string{"B", "X"}: 0, // Paper vs Rock  - lose
		[2]string{"B", "Y"}: 3, // Paper vs Paper - draw
		[2]string{"B", "Z"}: 6, // Paper vs Sciccors -  win
		[2]string{"C", "X"}: 6, // Sciccors vs Rock  - win
		[2]string{"C", "Y"}: 0, // Sciccors vs Paper -  lose
		[2]string{"C", "Z"}: 3, // Sciccors vs Sciccors - draw
	}
	var score int = 0
	var yourself string = array[1]
	var opponent string = array[0]

	if yourself == "X" {
		score += 1
	} else if yourself == "Y" {
		score += 2
	} else if yourself == "Z" {
		score += 3
	}

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
