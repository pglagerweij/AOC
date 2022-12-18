package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// var inputFile string = "trial.txt"

var inputFile string = "input.txt"

var maxSteps int = 1000000000000

type seen struct {
	rocks     int
	height    int
	situation struct {
		flow     byte
		element  int
		position [25][7]int
	}
}

func main() {
	totalbuffer := []seen{}
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	cave := make([][7]int, 1)
	cave[0] = [7]int{1, 1, 1, 1, 1, 1, 1}
	// fmt.Printf("%v\n", cave)
	timer := 0
	rockCounter := 0
	highestRock := 0
	var highestFaillingRock int
	var lowestFaillingRock int
	var leftIndex int
	var rightIndex int
	for rockCounter < maxSteps {
		inPlace := false
		// fmt.Printf("Rock to the %v with type %v\n", rockCounter, (rockCounter+1)%6)
		typeOfRock := (rockCounter) % 5
		cave, leftIndex, rightIndex, lowestFaillingRock, highestFaillingRock = enLargecave(cave, typeOfRock, highestRock)
		// if rockCounter%1 == 0 && rockCounter > 313 {
		// 	fmt.Printf("After rocks %v\n", rockCounter+1)
		// 	for index := len(cave) - 1; index >= maxofInts(len(cave)-21, 0); index-- {
		// 		fmt.Printf("%v\n", cave[index])
		// 	}
		// }
		for inPlace == false {
			windway := fileContent[timer%len(fileContent)]
			cave, leftIndex, rightIndex = windOnRock(cave, highestFaillingRock, lowestFaillingRock, leftIndex, rightIndex, windway)
			// if rockCounter%1 == 0 && rockCounter > 314 {
			// 	time.Sleep(1 * time.Second)
			// 	fmt.Printf("blown the rock to the %c\n", windway)
			// 	for index := len(cave) - 1; index >= maxofInts(len(cave)-21, 0); index-- {
			// 		fmt.Printf("%v\n", cave[index])
			// 	}
			// }

			cave, highestFaillingRock, lowestFaillingRock, inPlace = godownnRock(cave, highestFaillingRock, lowestFaillingRock, typeOfRock)
			timer += 1
			// if rockCounter%1 == 0 && rockCounter > 314 {
			// 	fmt.Printf("Down 1 positions\n")
			// 	for index := len(cave) - 1; index >= maxofInts(len(cave)-21, 0); index-- {
			// 		fmt.Printf("%v\n", cave[index])
			// 	}
			// }

		}

		for ind := len(cave) - 1; ind > 0; ind-- {
			if sumSlice(cave[ind]) != 0 {
				break
			}
			highestRock = ind - 1
		}

		if len(cave) > 25 {
			rockCounter += 1
			var bufferelem seen
			bufferelem.rocks = rockCounter
			bufferelem.height = highestRock
			bufferelem.situation.flow = fileContent[timer%len(fileContent)]
			bufferelem.situation.element = (rockCounter) % 5
			var currentCaveposition [25][7]int
			relevantCave := cave[len(cave)-25:]
			copy(currentCaveposition[:], relevantCave[:25])
			bufferelem.situation.position = currentCaveposition
			// res := false
			for _, element := range totalbuffer {
				if bufferelem.situation == element.situation {
					fmt.Printf("Hello we found a match! from %v to %v\n", element.rocks, rockCounter)
					drocks := rockCounter - element.rocks
					dheight := bufferelem.height - element.height
					drocksleft := maxSteps - rockCounter
					rockstoAdd := drocksleft / drocks * dheight
					timetoAdd := drocksleft / drocks * drocks
					// startingHeight := highestRock
					highestRock += rockstoAdd
					rockCounter += timetoAdd
					// res = true

					fmt.Printf("We have %v height at rocks dropped %v\n", highestRock, rockCounter)

					stepsLeft := maxSteps - rockCounter
					indextoFind := element.rocks + stepsLeft
					for _, secondelem := range totalbuffer {
						if secondelem.rocks == indextoFind {
							rocksadded := secondelem.height - element.height

							fmt.Printf("we have found the second step, at steps %v there is a height of %v\n", maxSteps, highestRock+rocksadded)
							os.Exit(3)
							break
						}
					}
					// fmt.Printf("We have bufffer %v\n", totalbuffer)

					break
				}
			}

			totalbuffer = append(totalbuffer, bufferelem)
		} else {
			rockCounter += 1
		}
	}

	// for index := len(cave) - 1; index >= 0; index-- {
	// 	fmt.Printf("%v\n", cave[index])
	// }
	fmt.Printf("The highest rock is on %v after %v rocks have fallen.\n", highestRock, rockCounter)
}

func enLargecave(cave [][7]int, typeOfRock int, highestRock int) ([][7]int, int, int, int, int) {
	heightToAdd := highestRock + 4 - (len(cave))
	if heightToAdd > 0 {
		cavetoAdd := make([][7]int, heightToAdd)
		cave = append(cave, cavetoAdd...)
	}
	if typeOfRock == 0 {
		rockToAdd := make([][7]int, 1)
		rockToAdd[0] = [7]int{0, 0, 2, 2, 2, 2, 0}
		cave = append(cave[:highestRock+4], rockToAdd...)
		return cave, 2, 5, highestRock + 4, highestRock + 4
	} else if typeOfRock == 1 {
		rockToAdd := make([][7]int, 3)
		rockToAdd[0] = [7]int{0, 0, 0, 2, 0, 0, 0}
		rockToAdd[1] = [7]int{0, 0, 2, 2, 2, 0, 0}
		rockToAdd[2] = [7]int{0, 0, 0, 2, 0, 0, 0}
		cave = append(cave[:highestRock+4], rockToAdd...)
		return cave, 2, 4, highestRock + 4, highestRock + 6
	} else if typeOfRock == 2 {
		rockToAdd := make([][7]int, 3)
		rockToAdd[0] = [7]int{0, 0, 2, 2, 2, 0, 0}
		rockToAdd[1] = [7]int{0, 0, 0, 0, 2, 0, 0}
		rockToAdd[2] = [7]int{0, 0, 0, 0, 2, 0, 0}
		cave = append(cave[:highestRock+4], rockToAdd...)
		return cave, 2, 4, highestRock + 4, highestRock + 6
	} else if typeOfRock == 3 {
		rockToAdd := make([][7]int, 4)
		rockToAdd[0] = [7]int{0, 0, 2, 0, 0, 0, 0}
		rockToAdd[1] = [7]int{0, 0, 2, 0, 0, 0, 0}
		rockToAdd[2] = [7]int{0, 0, 2, 0, 0, 0, 0}
		rockToAdd[3] = [7]int{0, 0, 2, 0, 0, 0, 0}
		cave = append(cave[:highestRock+4], rockToAdd...)
		return cave, 2, 2, highestRock + 4, highestRock + 7
	} else if typeOfRock == 4 {
		rockToAdd := make([][7]int, 2)
		rockToAdd[0] = [7]int{0, 0, 2, 2, 0, 0, 0}
		rockToAdd[1] = [7]int{0, 0, 2, 2, 0, 0, 0}
		cave = append(cave[:highestRock+4], rockToAdd...)
		return cave, 2, 3, highestRock + 4, highestRock + 5
	} else {
		panic("help")
	}

}

func godownnRock(cave [][7]int, highestpoint int, lowestpoint int, typeofrock int) ([][7]int, int, int, bool) {
	result := false
	if typeofrock == 1 {
		index2 := lowestpoint + 1
		for indRow, element := range cave[index2] {
			// fmt.Printf("row %v, index %v, element %v\n", index, indRow, element)
			if element == 2 && cave[index2-1][indRow] == 1 {
				for indexcleanup := lowestpoint; indexcleanup <= highestpoint; indexcleanup++ {
					for indRow, element := range cave[indexcleanup] {
						if element == 2 {
							cave[indexcleanup][indRow] = 1
						}
					}
				}
				return cave, highestpoint, lowestpoint, true
			}
		}
	}
	for index := lowestpoint; index <= highestpoint; index++ {
		newRow := remove2s(cave[index-1])
		for indRow, element := range cave[index] {
			// fmt.Printf("row %v, index %v, element %v\n", index, indRow, element)
			if element == 2 {
				newRow[indRow] = 2
			}
			if element == 2 && cave[index-1][indRow] == 1 {
				result = true

			}
		}
		if result == false {
			cave[index-1] = newRow
		} else if result == true {
			for indRow, element := range cave[index] {
				if element == 2 {
					cave[index][indRow] = 1
				}
			}
		}

		if index == highestpoint && result == false {
			cave[index] = remove2s(cave[index])
		}

	}

	return cave, highestpoint - 1, lowestpoint - 1, result
}

func remove2s(input [7]int) [7]int {
	for ind, elem := range input {
		if elem == 2 {
			input[ind] = 0
		}
	}
	return input
}

func windOnRock(cave [][7]int, highestRock int, lowestRock int, leftIndex int, rightIndex int, way byte) ([][7]int, int, int) {
	relevantCave := cave[lowestRock : highestRock+1]
	// fmt.Printf("the cave is: \n %v\n", relevantCave)
	if (way == '<' && leftIndex == 0) || (way == '>' && rightIndex == 6) {
		// fmt.Printf("Stone cannot move\n")
		return cave, leftIndex, rightIndex
	} else if canItMove(relevantCave, way) == false {
		// fmt.Printf("Stone cannot move\n")
		return cave, leftIndex, rightIndex
	}

	if way == '<' && leftIndex != 0 {
		for caveHeight := lowestRock; caveHeight <= highestRock; caveHeight++ {
			newRow := [7]int{0, 0, 0, 0, 0, 0, 0}
			for index := range newRow {
				if cave[caveHeight][index] == 2 {
					newRow[index-1] = 2
				} else if cave[caveHeight][index] == 1 {
					newRow[index] = 1
				}
			}
			// fmt.Printf("old row is %v, new row is %v\n.", cave[caveHeight], newRow)
			cave[caveHeight] = newRow
		}
		return cave, leftIndex - 1, rightIndex - 1
	} else if way == '>' && rightIndex != 6 {
		for caveHeight := lowestRock; caveHeight <= highestRock; caveHeight++ {
			newRow := [7]int{0, 0, 0, 0, 0, 0, 0}
			for index := range newRow {
				if cave[caveHeight][index] == 2 {
					newRow[index+1] = 2
				} else if cave[caveHeight][index] == 1 {
					newRow[index] = 1
				}
			}
			// fmt.Printf("old row is %v, new row is %v\n", cave[caveHeight], newRow)
			cave[caveHeight] = newRow
		}
		return cave, leftIndex + 1, rightIndex + 1
	} else {
		return cave, leftIndex, rightIndex
	}

}

func canItMove(cave [][7]int, way byte) bool {
	result := true
	if way == '<' {
		for _, caveRow := range cave {
			for ind := 0; ind < len(caveRow); ind++ {
				if caveRow[ind] == 2 {
					if caveRow[ind-1] == 1 {
						result = false
						return result
					}
					break
					// return result
				}
			}
		}
		return result
	} else if way == '>' {
		for _, caveRow := range cave {
			for ind := len(caveRow) - 1; ind >= 0; ind-- {
				if caveRow[ind] == 2 {
					if caveRow[ind+1] == 1 {
						result = false
						return result
					}
					break

				}
			}
		}
		return result
	} else {
		panic("I do not know this way")
	}

}

func maxofInts(a int, b int) int {
	if a > b {
		return a
	} else if a <= b {
		return b
	} else {
		panic("didnt get ints")
	}
}

func sumSlice(numarray [7]int) int {
	var arrSum int = 0

	for i := 0; i < len(numarray); i++ {
		arrSum = arrSum + numarray[i]
	}
	return arrSum
}
