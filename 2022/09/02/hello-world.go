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
	s2 := strings.ReplaceAll(s, " ", ",")
	s3 := strings.Split(s2, ";")
	locationHead := [2]int{}
	locationHead[0] = 0
	locationHead[1] = 0
	locationKnots := [8][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	locationTail := [2]int{0, 0}
	locationsofTail := [][2]int{{0, 0}}
	// currentIndex := 0
	for _, element := range s3 {
		// fmt.Printf("the element we are processing is %v\n", element)
		// fmt.Printf("the current location is %v\n", locationHead)
		elementArray := strings.Split(element, ",")
		direction := elementArray[0]
		steps, _ := strconv.Atoi(strings.TrimSpace(elementArray[1]))
		// fmt.Printf("%v", error)
		for step := 0; step < steps; step++ {
			locationHead = moveHead(locationHead, string(direction), 1)
			// fmt.Printf("The head moved in the direction %v with %v steps and is now at %v\n", direction, steps, locationHead)
			for index, _ := range locationKnots {
				// fmt.Printf("%v", index)
				if index == 0 {
					locationKnots[index] = moveTail(locationKnots[index], locationHead)
					// fmt.Printf("The knot %v moved and is now at %v\n", index, locationKnots[index])
				} else {
					locationKnots[index] = moveTail(locationKnots[index], locationKnots[index-1])
					// fmt.Printf("The knot %v moved and is now at %v\n", index, locationKnots[index])
				}

			}

			locationTail = moveTail(locationTail, locationKnots[7])
			// fmt.Printf("The tail moved and is now at %v\n", locationTail)
			locationsofTail = append(locationsofTail, locationTail)
		}

	}
	// fmt.Printf("%v\n", locationsofTail)
	var uniqueLocationsofTail [][2]int
	for _, v := range locationsofTail {
		skip := false
		for _, u := range uniqueLocationsofTail {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			uniqueLocationsofTail = append(uniqueLocationsofTail, v)
		}
	}
	// fmt.Printf("the total number of unique spots of the tail is: %v", uniqueLocationsofTail)
	fmt.Printf("the total number of unique spots of the tail is: %v", len(uniqueLocationsofTail))

}

func unqique() {

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func moveTail(tailLocation [2]int, headLocation [2]int) [2]int {
	updatedTailLocation := tailLocation
	// fmt.Printf("The tail moved and is now at %v\n", updatedTailLocation)
	rowDistance := abs(headLocation[0] - tailLocation[0])
	columnDistance := abs(headLocation[1] - tailLocation[1])
	// fmt.Printf("The distance is %v and %v\n", rowDistance, columnDistance)
	// Tail doesn't have to move
	if (rowDistance == 1 && columnDistance == 1) || (rowDistance == 0 && columnDistance == 0) {
		return updatedTailLocation
	} else if rowDistance == 1 && columnDistance == 0 {
		return updatedTailLocation
	} else if rowDistance == 0 && columnDistance == 1 {
		return updatedTailLocation
	} else if rowDistance == 2 && columnDistance == 0 {
		updatedTailLocation[0] = (headLocation[0] + tailLocation[0]) / 2
		return updatedTailLocation
	} else if rowDistance == 0 && columnDistance == 2 {
		updatedTailLocation[1] = (headLocation[1] + tailLocation[1]) / 2
		return updatedTailLocation
	} else if rowDistance == 2 && columnDistance == 1 {
		updatedTailLocation[0] = (headLocation[0] + tailLocation[0]) / 2
		updatedTailLocation[1] = headLocation[1]
		return updatedTailLocation
	} else if rowDistance == 1 && columnDistance == 2 {
		updatedTailLocation[0] = headLocation[0]
		updatedTailLocation[1] = (headLocation[1] + tailLocation[1]) / 2
		return updatedTailLocation
	} else if rowDistance == 2 && columnDistance == 2 {
		updatedTailLocation[0] = (headLocation[0] + tailLocation[0]) / 2
		updatedTailLocation[1] = (headLocation[1] + tailLocation[1]) / 2
		return updatedTailLocation
	} else {
		panic("hello")
	}
}

func moveHead(initialLocation [2]int, direction string, numberOfSteps int) [2]int {
	updatedLocation := initialLocation
	// fmt.Printf("%v\n", updatedLocation)
	// fmt.Printf("%v\n", direction)
	if direction == "R" {
		// fmt.Printf("%v with %v steps \n", direction, numberOfSteps)
		updatedLocation[0] = updatedLocation[0] + numberOfSteps
		return updatedLocation
	} else if direction == "L" {
		// fmt.Printf("%v with %v steps \n", direction, numberOfSteps)
		updatedLocation[0] = updatedLocation[0] - numberOfSteps
		return updatedLocation
	} else if direction == "U" {
		// fmt.Printf("%v with %v steps \n", direction, numberOfSteps)
		updatedLocation[1] = updatedLocation[1] + numberOfSteps
		return updatedLocation
	} else if direction == "D" {
		// fmt.Printf("%v with %v steps \n", direction, numberOfSteps)
		updatedLocation[1] = updatedLocation[1] - numberOfSteps
		return updatedLocation
	} else {
		panic("hello")
	}
}
