package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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
	sArray := make([][3]int, 0)
	for _, elem := range s {
		elemsplitted := strings.Split(elem, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[0]))
		xnumbers = append(xnumbers, x)
		y, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[1]))
		ynumbers = append(ynumbers, y)
		z, _ := strconv.Atoi(strings.TrimSpace(elemsplitted[2]))
		znumbers = append(znumbers, z)
		sArray = append(sArray, [3]int{x, y, z})
	}

	xmin := minofslice(unique(xnumbers))
	xmax := maxofslice(unique(xnumbers))

	ymin := minofslice(unique(ynumbers))
	ymax := maxofslice(unique(ynumbers))

	zmin := minofslice(unique(znumbers))
	zmax := maxofslice(unique(znumbers))
	// fmt.Printf("Total grid, x from %v to %v, y from %v to %v, z from %v to %v\n", xmin, xmax, ymin, ymax, zmin, zmax)
	airlist := make([][3]int, 0)
	for x := xmin + 1; x <= xmax-1; x++ {
		for y := ymin + 1; y <= ymax-1; y++ {
			for z := zmin + 1; z <= zmax-1; z++ {
				result := false
				for _, allDrops := range sArray {
					if [3]int{x, y, z} == allDrops {
						result = true
						break
					}
				}
				if result == false {
					airlist = append(airlist, [3]int{x, y, z})
				}
			}
		}
	}

	// fmt.Printf("the total air list is %v\n", airlist)

	groupedin := getClosed(airlist, sArray, xmin, xmax, ymin, ymax, zmin, zmax)

	fillTotalset := append(sArray, groupedin...)
	fmt.Printf("the old length was %v new length %v\n", len(sArray), len(fillTotalset))
	xnumbers = []int{}
	ynumbers = []int{}
	znumbers = []int{}
	for _, elem := range fillTotalset {
		xnumbers = append(xnumbers, elem[0])
		ynumbers = append(ynumbers, elem[1])
		znumbers = append(znumbers, elem[2])
	}

	totalzside := 0
	for _, ytocheck := range unique(ynumbers) {
		for _, xtocheck := range unique(xnumbers) {
			zaxis := []int{}
			for _, elem := range fillTotalset {
				x := elem[0]
				y := elem[1]
				z := elem[2]
				if y == ytocheck && x == xtocheck {
					zaxis = append(zaxis, z)
				}
			}

			outsides := checkOutsides(zaxis)
			// fmt.Printf("For side with x: %v, y: %v we have zaxis: %v with outsides %v.\n", xtocheck, ytocheck, zaxis, outsides)
			totalzside += outsides
		}
	}

	fmt.Printf("Total over zside is %v\n", totalzside)

	totalyside := 0
	for _, ztocheck := range unique(znumbers) {
		for _, xtocheck := range unique(xnumbers) {
			yaxis := []int{}
			for _, elem := range fillTotalset {
				x := elem[0]
				y := elem[1]
				z := elem[2]
				if z == ztocheck && x == xtocheck {
					yaxis = append(yaxis, y)
				}
			}

			outsides := checkOutsides(yaxis)
			// fmt.Printf("For side with x: %v, z: %v we have zaxis: %v with outsides %v.\n", xtocheck, ztocheck, yaxis, outsides)
			totalyside += outsides
		}
	}
	fmt.Printf("Total over yside is %v\n", totalyside)

	totalxside := 0
	for _, ztocheck := range unique(znumbers) {
		for _, ytocheck := range unique(xnumbers) {
			xaxis := []int{}
			for _, elem := range fillTotalset {
				x := elem[0]
				y := elem[1]
				z := elem[2]
				if z == ztocheck && y == ytocheck {
					xaxis = append(xaxis, x)
				}
			}

			outsides := checkOutsides(xaxis)
			// fmt.Printf("For side with x: %v, z: %v we have zaxis: %v with outsides %v.\n", ytocheck, ztocheck, xaxis, outsides)
			totalxside += outsides
		}
	}
	fmt.Printf("Total over xside is %v\n", totalxside)

	totalsides := totalxside + totalyside + totalzside

	fmt.Printf("Total exterior over all sides is %v\n", totalsides)

}

func getClosed(air [][3]int, drops [][3]int, xmin int, xmax int, ymin int, ymax int, zmin int, zmax int) [][3]int {
	output := make([][3]int, 0)
	for _, airDrop := range air {
		// fmt.Printf("%v\n ", airDrop)
		xair := airDrop[0]
		yair := airDrop[1]
		zair := airDrop[2]
		xminres := false
		for x := xmin; x < xair; x++ {
			for _, drops := range drops {
				// fmt.Printf("for x %v with %v\n ", x, drops)
				if [3]int{x, yair, zair} == drops {
					xminres = true
					break
				}
			}
		}
		if xminres == true {
			xmaxres := false
			for x := xmax; x > xair; x-- {
				for _, drops := range drops {
					// fmt.Printf("for x %v with %v\n ", x, drops)
					if [3]int{x, yair, zair} == drops {
						xmaxres = true
						break
					}
				}
			}

			if xmaxres == true {
				yminres := false
				for y := ymin; y < yair; y++ {
					for _, drops := range drops {
						// fmt.Printf("for y %v with %v\n ", y, drops)
						if [3]int{xair, y, zair} == drops {
							yminres = true
							break
						}
					}
				}
				if yminres == true {
					ymaxres := false
					for y := ymax; y > yair; y-- {
						for _, drops := range drops {
							// fmt.Printf("for y %v with %v\n ", y, drops)
							if [3]int{xair, y, zair} == drops {
								ymaxres = true
								break
							}
						}
					}

					if ymaxres == true {
						zminres := false
						for z := zmin; z < zair; z++ {
							for _, drops := range drops {
								// fmt.Printf("for z %v with %v\n ", z, drops)
								if [3]int{xair, yair, z} == drops {
									zminres = true
									break
								}
							}
						}
						if zminres == true {
							zmaxres := false
							for z := zmax; z > zair; z-- {
								for _, drops := range drops {
									// fmt.Printf("for z %v with %v\n ", z, drops)
									if [3]int{xair, yair, z} == drops {
										zmaxres = true
										break
									}
								}
							}

							if zmaxres == true {
								output = append(output, airDrop)
							}

						}
					}
				}
			}
		}

	}

	return output
}
func checkArray(start [][3]int, check [][3]int) [][3]int {
	group := start
	elemensweDelete := make([][3]int, 0)
	for _, startElem := range start {
		for _, checkElem := range check {
			if abs(startElem[0]-checkElem[0])+abs(startElem[1]-checkElem[1])+abs(startElem[2]-checkElem[2]) == 1 {
				group = append(group, checkElem)
				elemensweDelete = append(elemensweDelete, checkElem)
			}
		}
	}

	if len(elemensweDelete) == 0 {
		return group
	} else {
		checkleft := make([][3]int, 0)
		for _, elem := range check {
			result := false
			for _, elemDel := range elemensweDelete {
				if elem == elemDel {
					result = true
					break
				}
			}
			if result == false {
				checkleft = append(checkleft, elem)
			}
		}
		return checkArray(group, checkleft)
	}

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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxofslice(x []int) int {
	max := x[0]
	for _, xelem := range x {
		if xelem > max {
			max = xelem
		}
	}
	return max
}

func minofslice(x []int) int {
	min := x[0]
	for _, xelem := range x {
		if xelem < min {
			min = xelem
		}
	}
	return min
}
