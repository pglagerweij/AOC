package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// var inputFile string = "trial.txt"
// var indextograp int = 10

var inputFile string = "input.txt"
var indextograp int = 2000000

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

	maxElementX := 0
	minElementX := 900
	maxElementY := 0
	minElementY := 50000000000000000
	maxdistance := 0
	for _, element := range s3 {
		// fmt.Printf("we are processing: %v\n", element)
		sensorx, sensory, beaconx, beacony := retrieveElements(element)
		distance := calculateDistance(sensorx, sensory, beaconx, beacony)
		// fmt.Printf("Sensor at %v, %v, closest beacon at %v, %v\n", sensorx, sensory, beaconx, beacony)
		if sensorx < minElementX {
			minElementX = sensorx
		} else if sensorx > maxElementX {
			maxElementX = sensorx
		}

		// if beaconx < minElementX {
		// 	minElementX = beaconx
		// } else if beaconx > maxElementX {
		// 	maxElementX = beaconx
		// }

		if sensory < minElementY {
			minElementY = sensory
		} else if sensory > maxElementY {
			maxElementY = sensory
		}

		// if beacony < minElementY {
		// 	minElementY = beacony
		// } else if beacony > maxElementY {
		// 	maxElementY = beacony
		// }

		if distance > maxdistance {
			maxdistance = distance
		}
	}
	fmt.Printf("min element X at %v, max at %v. Min element Y at %v, max at %v.\n", minElementX, maxElementX, minElementY, maxElementY)
	fmt.Printf("max distance is %v\n", maxdistance)

	// Making array that is large enough using maxdistance and min and max elements

	xnullindex := minElementX - maxdistance
	ynullindex := minElementY - maxdistance

	// totalYaxisLength := maxElementY - minElementY + 2*maxdistance
	totalXaxisLength := maxElementX - minElementX + 2*maxdistance
	relevantYaxis := make([]uint8, totalXaxisLength)
	// totalSituation := make([][]uint8, totalYaxisLength)
	// for i := range totalSituation {
	// 	totalSituation[i] = make([]uint8, totalXaxisLength)
	// }

	for _, element := range s3 {
		// fmt.Printf("we are processing: %v\n", element)
		sensorx, sensory, beaconx, beacony := retrieveElements(element)
		// totalSituation = drawInMap(totalSituation, sensorx, sensory, xnullindex, ynullindex, "sensor")
		relevantYaxis = drawInMapYaxis(relevantYaxis, beaconx, beacony, xnullindex, ynullindex, indextograp, "beacon")

		distance := calculateDistance(sensorx, sensory, beaconx, beacony)
		// totalSituation = drawNonBeacons(totalSituation, sensorx, sensory, xnullindex, ynullindex, distance)
		relevantYaxis = drawNonBeaconsYaxis(relevantYaxis, sensorx, sensory, xnullindex, ynullindex, indextograp, distance)
	}

	// yindexrowtograp := indextograp - ynullindex
	// fmt.Printf("%v\n", relevantYaxis)
	totalSum := 0
	for _, element := range relevantYaxis {
		if element == 1 {
			totalSum += 1
		}
	}
	fmt.Printf("the total value is %v", totalSum)
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

func drawNonBeacons(totalSituation [][]uint8, sensorx int, sensory int, xnullindex int, ynullindex int, distance int) [][]uint8 {

	// fmt.Printf("Looping through list with distance %v\n", distance)
	for index := -1 * distance; index <= distance; index++ {
		ydistance := distance - abs(index)
		for yloop := -1 * ydistance; yloop <= ydistance; yloop++ {
			// fmt.Printf("At value x: %v, y: %v\n", index, yloop)
			xIndex := sensorx + index
			yIndex := sensory - ynullindex + yloop
			if totalSituation[yIndex][xIndex] == 0 {
				totalSituation[yIndex][xIndex] = 1
			}

		}
	}

	return totalSituation
}

func drawInMap(totalSituation [][]uint8, sensorx int, sensory int, xnullindex int, ynullindex int, objecttype string) [][]uint8 {

	xIndex := sensorx - xnullindex
	yIndex := sensory - ynullindex

	if objecttype == "sensor" {
		totalSituation[yIndex][xIndex] = 5
		return totalSituation
	} else if objecttype == "beacon" {
		totalSituation[yIndex][xIndex] = 2
		return totalSituation
	} else {
		panic("cannot find type")
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
