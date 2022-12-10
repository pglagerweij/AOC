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
	currentSpriteLoc := 1
	drawing := []string{}
	for _, element := range s3 {
		fmt.Printf("Start cycle %v Begin executing %v\n", currentIndex+1, element)
		if strings.HasPrefix(element, "noop") {
			fmt.Printf("During cycle %v Begin executing %v\n", currentIndex+1, element)
			currentIndex += 1
			charstoadd := chartoDraw(drawing, currentIndex, currentSpriteLoc)
			// fmt.Printf("Char we drawing %v\n", charstoadd)
			drawing = append(drawing, charstoadd)
			valueOverTime = append(valueOverTime, valueOverTime[currentIndex-1])
			// fmt.Println(currentPath)
		} else if strings.HasPrefix(element, "addx ") {
			splitElement := strings.Split(element, " ")
			value, _ := strconv.Atoi(strings.TrimSpace(splitElement[1]))
			fmt.Printf("During cycle %v Begin executing %v\n", currentIndex+1, element)
			currentIndex += 1
			charstoadd := chartoDraw(drawing, currentIndex, currentSpriteLoc)
			drawing = append(drawing, charstoadd)

			valueOverTime = append(valueOverTime, valueOverTime[currentIndex-1])

			fmt.Printf("During cycle %v Begin executing %v\n", currentIndex+1, element)
			currentIndex += 1
			charstoadd2 := chartoDraw(drawing, currentIndex, currentSpriteLoc)
			fmt.Printf("Char we drawing %v\n", charstoadd)
			drawing = append(drawing, charstoadd2)

			valueOverTime = append(valueOverTime, valueOverTime[currentIndex-1]+value)
			currentSpriteLoc = currentSpriteLoc + value
			// fmt.Printf("next ls command: %v\n", element)
			// allFiles = fillDirectory(allFiles, element, currentPath)
		}

	}

	for index, char := range drawing {
		fmt.Printf("%v", char)
		if (index+1)%40 == 0 && index != 0 {
			fmt.Printf("\n")
		}
	}

	// fmt.Printf("The total visible trees are %v", valueOverTime)
}

func chartoDraw(input []string, cycle int, spritIndex int) string {
	fmt.Printf("Current cycle is %v the mainspritindex is %v\n", cycle, spritIndex)
	cyc := cycle % 40
	if cyc >= spritIndex && cyc <= spritIndex+2 {
		fmt.Printf("Current cycle is %v and there for %v the mainspritindex is %v this results in a #\n", cycle, cyc, spritIndex)
		return "#"
	} else {
		fmt.Printf("Current cycle is %v and there for %v the mainspritindex is %v this results in a .\n", cycle, cyc, spritIndex)
		return "."
	}

}
