package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// var inputFile string = "trial.txt"
// var indextograp int = 20

var inputFile string = "input.txt"
var indextograp int = 4000000

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
	// for xindex := 0; xindex <= indextograp; xindex++ {
	// 	var result bool = false
	// 	for yindex := 0; yindex <= indextograp; yindex++ {
	// 		fmt.Printf("x value %v, y value %v\n", xindex, yindex)
	// 		for _, element := range s3 {
	// 			sensorx, sensory, beaconx, beacony := retrieveElements(element)
	// 			distance := calculateDistance(sensorx, sensory, beaconx, beacony)
	// 			distanceCoor := calculateDistance(sensorx, sensory, xindex, yindex)
	// 			if distanceCoor < distance {
	// 				result = true
	// 				break
	// 			}
	// 		}
	// 		if result == false {
	// 			fmt.Printf("the number without support is %v, %v", xindex, yindex)
	// 			break
	// 		}
	// 	}
	// }
	// xindex := 14
	// yindex := 11
	// for xindex := 0; xindex <= indextograp; xindex++ {
	// 	for yindex := 0; yindex <= indextograp; yindex++ {
	// 		totalresult := true
	// 		fmt.Printf("the values are %v and %v\n", xindex, yindex)
	// 		for _, element := range s3 {
	// 			sensorx, sensory, beaconx, beacony := retrieveElements(element)
	// 			distance := calculateDistance(sensorx, sensory, beaconx, beacony)
	// 			distanceCoor := calculateDistance(sensorx, sensory, xindex, yindex)
	// 			fmt.Printf("the values are %v and %v for sensor %v\n", distance, distanceCoor, element)
	// 			if distance >= distanceCoor {
	// 				totalresult = false
	// 				break
	// 			}

	// 		}
	// 		if totalresult == true {
	// 			fmt.Printf("The result is %v for %v,%v\n", totalresult, xindex, yindex)
	// 			os.Exit(3)
	// 		}
	// 	}
	// }
	// xindex := 0
	for xindex := 0; xindex <= indextograp; xindex++ {
		fmt.Printf("the x-index are %v\n", xindex)
		for yindex := 0; yindex <= indextograp; {
			totalresult := true
			// fmt.Printf("the index are %v and %v\n", xindex, yindex)
			for _, element := range s3 {
				sensorx, sensory, beaconx, beacony := retrieveElements(element)
				distance := calculateDistance(sensorx, sensory, beaconx, beacony)
				distanceCoor := calculateDistance(sensorx, sensory, xindex, yindex)
				// fmt.Printf("the values are %v and %v for sensor %v\n", distance, distanceCoor, element)
				if distance >= distanceCoor {
					// fmt.Printf("the distance is %v and %v for sensor %v\n", distance, distanceCoor, element)
					leftover := distance - distanceCoor
					totalresult = false
					yindex += leftover + 1
					break
				}

			}

			if totalresult == true {
				fmt.Printf("The result is %v for %v,%v\n", totalresult, xindex, yindex)
				fmt.Printf("The result is %v\n", xindex*4000000+yindex)
				os.Exit(3)
			}
		}
	}

}

func drawInMapYaxis(totalSituation []uint8, sensorx int, sensory int, xnullindex int, ynullindex int, indextograp int, objecttype string) []uint8 {

	xIndex := sensorx - xnullindex
	if objecttype == "sensor" {
		if sensory == indextograp {
			totalSituation[xIndex] = 5
		}
		return totalSituation
	} else if objecttype == "beacon" {
		if sensory == indextograp {
			totalSituation[xIndex] = 2
		}
		return totalSituation
	} else {
		panic("cannot find type")
	}
}

func drawNonBeaconsYaxis(relevantYaxis []uint8, sensorx int, sensory int, xnullindex int, ynullindex int, indextograp int, distance int) []uint8 {

	if sensory-distance <= indextograp && sensory+distance >= indextograp {
		ydistance := abs(sensory - indextograp)
		xdistance := distance - ydistance
		for xloop := -1 * xdistance; xloop <= xdistance; xloop++ {
			// fmt.Printf("At value x: %v, y: %v\n", index, yloop)
			xIndex := sensorx - xnullindex - xloop
			if relevantYaxis[xIndex] == 0 {
				relevantYaxis[xIndex] = 1
			}
		}
		return relevantYaxis
	} else {
		return relevantYaxis
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateDistance(sensorx int, sensory int, beaconx int, beacony int) int {
	xDistance := abs(sensorx - beaconx)
	yDistance := abs(sensory - beacony)
	return xDistance + yDistance
}

func retrieveElements(input string) (int, int, int, int) {
	elements := strings.Split(input, ": ")
	sensorInfo := strings.Split(elements[0], ",")
	// fmt.Printf("%v\n", sensorInfo)
	sensorX := strings.TrimSpace(sensorInfo[0][12:])
	// fmt.Printf("%v\n", sensorX)
	sensorY := strings.TrimSpace(sensorInfo[1][3:])
	// fmt.Printf("%v\n", sensorY)
	sensorXint, _ := strconv.Atoi(sensorX)
	sensorYint, _ := strconv.Atoi(sensorY)

	// Obtain beacon info
	beaconInfo := strings.Split(elements[1], ",")
	// fmt.Printf("%v\n", beaconInfo)
	beaconX := strings.TrimSpace(beaconInfo[0][23:])
	// fmt.Printf("%v\n", beaconX)
	beaconY := strings.TrimSpace(beaconInfo[1][3:])
	// fmt.Printf("%v\n", beaconY)
	beaconXint, _ := strconv.Atoi(beaconX)
	beaconYint, _ := strconv.Atoi(beaconY)

	return sensorXint, sensorYint, beaconXint, beaconYint
}

func letItSnow(a [][]uint8, snowPositionx int, maxPositionY int, maxPositionX int) ([][]uint8, bool) {
	snowPositionY := 1
	for {
		nextSnowPositionY, nextSnowPositionX, result := calculateNextPos(a, snowPositionx, snowPositionY, maxPositionY, maxPositionX)

		if result == true {
			return a, true
		} else if nextSnowPositionY == snowPositionY && nextSnowPositionX == snowPositionx {
			// fmt.Printf("snow didnt move\n")
			a[snowPositionY][snowPositionx] = 2
			return a, false
		}
		// fmt.Printf("value new position %v,%v\n", nextSnowPositionY, nextSnowPositionX)
		snowPositionx = nextSnowPositionX
		snowPositionY = nextSnowPositionY
	}

	// fmt.Printf("value new position %v,%v\n", nextSnowPositionY, nextSnowPositionX)
	// return a
}

func calculateNextPos(a [][]uint8, snowPositionx int, snowPositionY int, maxPositionY int, maxPositionX int) (int, int, bool) {
	if snowPositionY+1 > maxPositionY {
		return snowPositionY, snowPositionx, true
	} else if a[snowPositionY+1][snowPositionx] == 0 {
		return snowPositionY + 1, snowPositionx, false
	} else if a[snowPositionY+1][snowPositionx] != 0 {
		if snowPositionx-1 < 0 {
			return snowPositionY, snowPositionx, true
		} else if a[snowPositionY+1][snowPositionx-1] == 0 {
			return snowPositionY + 1, snowPositionx - 1, false
		} else if snowPositionx+1 > maxPositionX {
			return snowPositionY, snowPositionx, true
		} else if a[snowPositionY+1][snowPositionx+1] == 0 {
			return snowPositionY + 1, snowPositionx + 1, false
		} else {
			return snowPositionY, snowPositionx, false
		}
	} else {
		panic("help cannot find match")
	}
}

// func checkNextPosition(a [][]uint8, snowPositionx int, snowPositionY int) string {

// 	return "free"
// }

func drawrockinMap(a [][]uint8, lastx int, nextx int, lasty int, nexty int, nullindexx int, nullindexy int) [][]uint8 {
	if lastx == nextx && lasty == nexty {
		return a
	} else if lastx == nextx {
		xindex := lastx - nullindexx
		// fmt.Printf("drawing in y direction from %v to %v on index %v \n", lasty, nexty, xindex)
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
		// fmt.Printf("drawing in y direction from %v to %v on index %v \n", lastx, nextx, yindex)
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
