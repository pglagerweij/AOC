package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile string = "trial.txt"
var inputLength string = "500"
var inputHeight int = 0

// var inputFile string = "input.txt"
// var inputLength string = "500.000"
// var inputHeight int = 000

func main() {

	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	s := strings.ReplaceAll(fileContent, "\n", ";")
	s3 := strings.Split(s, ";")

	maxElementX := 500
	minElementX := 500
	maxElementY := inputHeight
	minElementY := inputHeight
	for _, element := range s3 {
		elements := strings.Split(element, " -> ")
		for _, element := range elements {
			number := strings.Split(element, ",")
			elementX, _ := strconv.Atoi(number[0])
			elementY, _ := strconv.Atoi(number[1])
			if elementX > maxElementX {
				maxElementX = elementX
			}
			if elementX < minElementX {
				minElementX = elementX
			}
			if elementY > maxElementY {
				maxElementY = elementY
			}
			if elementY < minElementY {
				minElementY = elementY
			}
		}

	}
	// fmt.Printf("Value to draw is from %v to %v. And in lower direction from %v to %v.\n", minElementX, maxElementX, minElementY, maxElementY)
	anullindex := minElementX
	bnullindex := 0
	a := make([][]uint8, maxElementX-minElementX+1)
	for i := range a {
		a[i] = make([]uint8, maxElementY+1)
	}
	var lastElementX int
	var lastElementY int
	for _, element := range s3 {
		elements := strings.Split(element, " -> ")
		fmt.Printf("processing %v\n", element)
		for index, element2 := range elements {
			// fmt.Printf("processing %v with index %v\n", element2, index)
			if index == 0 {
				elem := strings.Split(element2, ",")
				lastElementX, _ = strconv.Atoi(elem[0])
				lastElementY, _ = strconv.Atoi(elem[1])
				fmt.Printf("updated last elements %v,%v\n", lastElementX, lastElementY)
			} else if index != 0 {
				elem := strings.Split(element2, ",")
				elemX, _ := strconv.Atoi(strings.TrimSpace(elem[0]))
				elemY, _ := strconv.Atoi(strings.TrimSpace(elem[1]))
				a = drawrockinMap(a, lastElementX, elemX, lastElementY, elemY, anullindex, bnullindex)
				lastElementX = elemX
				lastElementY = elemY
				fmt.Printf("updated last elements %v,%v\n", lastElementX, lastElementY)
			} else {
				fmt.Printf("why am I here")
			}

		}

	}
	for _, row := range a {
		fmt.Printf("%v\n", row)
	}

}

func drawrockinMap(a [][]uint8, lastx int, nextx int, lasty int, nexty int, nullindexx int, nullindexy int) [][]uint8 {
	if lastx == nextx && lasty == nexty {
		return a
	} else if lastx == nextx {
		xindex := lastx - nullindexx
		fmt.Printf("drawing in y direction from %v to %v on index %v \n", lasty, nexty, xindex)
		if lasty > nexty {
			for yindex := nexty - nullindexy; yindex <= lasty-nullindexy; yindex++ {
				a[yindex][xindex] = 1
			}
			return a
		} else if lasty < nexty {
			for yindex := lasty - nullindexy; yindex <= nexty-nullindexy; yindex++ {
				a[yindex][xindex] = 1
			}
			return a
		} else {
			panic("error")
		}
	} else if lasty == nexty {
		yindex := lasty - nullindexy
		fmt.Printf("drawing in y direction from %v to %v on index %v \n", lastx, nextx, yindex)
		if lastx > nextx {
			for xindex := nextx - nullindexx; xindex <= lastx-nullindexx; xindex++ {
				a[yindex][xindex] = 1
			}
			return a
		} else if lastx < nextx {
			for xindex := lastx - nullindexx; xindex <= nextx-nullindexx; xindex++ {
				a[yindex][xindex] = 1
			}
			return a
		} else {
			panic("error")
		}
	} else {
		fmt.Printf("drawing in x direction from %v to %v\n", lastx, nextx)
		fmt.Printf("drawing in y direction from %v to %v\n", lasty, nexty)
		panic("Cannot find options for drawing")
	}

}
