package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// var inputFile string = "triaal2.txt"

// var inputFile string = "trial.txt"

var inputFile string = "input.txt"

func main() {

	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileContent := string(file)
	s := strings.Split(fileContent, "\n")
	xnumbers := []int{}
	ynumbers := []int{}
	znumbers := []int{}
	for _, elem := range s {
		elemsplitted := strings.Split(elem, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[0]))
		xnumbers = append(xnumbers, x)
		y, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[1]))
		ynumbers = append(ynumbers, y)
		z, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[2]))
		znumbers = append(znumbers, z)

	}

	// ytocheck := 2
	// xtocheck := 1
	totalzside := 0
	for _, ytocheck := range unique(ynumbers) {
		for _, xtocheck := range unique(xnumbers) {
			zaxis := []int{}
			for _, elem := range s {
				elemsplitted := strings.Split(elem, ",")
				x, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[0]))
				y, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[1]))
				z, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[2]))
				if y == ytocheck && x == xtocheck {
					zaxis = append(zaxis, z)
				}
			}
			// fmt.Printf("For side with x: %v, y: %v we have zaxis: %v.\n", xtocheck, ytocheck, zaxis)
			outsides := checkOutsides(zaxis)
			totalzside += outsides
		}
	}

	fmt.Printf("Total over zside is %v\n", totalzside)

	totalyside := 0
	for _, ztocheck := range unique(znumbers) {
		for _, xtocheck := range unique(xnumbers) {
			yaxis := []int{}
			for _, elem := range s {
				elemsplitted := strings.Split(elem, ",")
				x, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[0]))
				y, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[1]))
				z, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[2]))
				if z == ztocheck && x == xtocheck {
					yaxis = append(yaxis, y)
				}
			}
			// fmt.Printf("For side with x: %v, z: %v we have zaxis: %v.\n", xtocheck, ztocheck, yaxis)
			outsides := checkOutsides(yaxis)
			totalyside += outsides
		}
	}
	fmt.Printf("Total over yside is %v\n", totalyside)

	totalxside := 0
	for _, ztocheck := range unique(znumbers) {
		for _, ytocheck := range unique(xnumbers) {
			xaxis := []int{}
			for _, elem := range s {
				elemsplitted := strings.Split(elem, ",")
				x, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[0]))
				y, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[1]))
				z, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[2]))
				if z == ztocheck && y == ytocheck {
					xaxis = append(xaxis, x)
				}
			}
			// fmt.Printf("For side with x: %v, z: %v we have zaxis: %v.\n", ytocheck, ztocheck, xaxis)
			outsides := checkOutsides(xaxis)
			totalxside += outsides
		}
	}
	fmt.Printf("Total over xside is %v\n", totalxside)

	totalsides := totalxside + totalyside + totalzside

	fmt.Printf("Total over all sides is %v\n", totalsides)
	// loop over x angle
	// totalXsum := 0
	// for x := 0; x <= 20; x++ {
	// 	for y := 0; y <= 20; y++ {
	// 		array := []int{}
	// 		for z := 0; z <= 20; z++ {
	// 			if sArray[x][y][z] {
	// 				array[z] = 1
	// 			}
	// 		}

	// 	}
	// }

	// loop over y angle

	// loop over z angle
}

func checkOutsides(axis []int) int {
	outsides := 0
	var lastindex int = -30
	for loop := 0; loop <= 21; loop++ {
		if elementExist(axis, loop) {
			// fmt.Printf("Found match %v\n", loop)
			if loop == lastindex+1 {
				lastindex = loop
				// fmt.Printf("We dont add an outside%v\n", loop)
			} else {
				lastindex = loop
				if outsides == 0 {
					outsides += 1
				} else {
					outsides += 2
				}
				// fmt.Printf("Add an outside %v\n", outsides)
			}
		}
	}
	if outsides > 0 {
		outsides += 1
	}

	return outsides
}

func elementExist(input []int, tocheck int) bool {
	var result bool = false
	for _, x := range input {
		if x == tocheck {
			result = true
			break
		}
	}

	return result
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
