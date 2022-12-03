package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/juliangruber/go-intersect"
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
	s3 := strings.Split(s, ";")
	// fmt.Printf("%v", len(s3)/3)
	total_priority := 0
	for index := 0; index < len(s3)/3; index++ {
		arrayElements := s3[index*3 : index*3+3]
		fmt.Printf("all elements are: %v\n", arrayElements)
		totalSet := make([][]string, 0)
		for _, element := range arrayElements {
			set := getSet(element)
			uniqueSet := removeDuplicateStr(set)
			totalSet = append(totalSet, uniqueSet)
		}
		intersectionOne := intersect.Simple(totalSet[0], totalSet[1])
		intersectionTwo := intersect.Simple(intersectionOne, totalSet[2])
		stringRepresentation := fmt.Sprintf("%v", intersectionTwo[0])
		stringRepresentation2 := fmt.Sprintf("%c", stringRepresentation[1])
		priority := getPriority(stringRepresentation2)
		total_priority += priority
	}
	fmt.Printf("The total priority is: %v\n", total_priority)
}

func getSet(inputString string) []string {
	var required1 []string
	for _, k := range inputString {
		// fmt.Printf("%c\n", k)
		required1 = append(required1, strconv.QuoteRune(k))
	}
	return required1
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func getPriority(inputChar string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	integerReturn := strings.Index(alphabet, inputChar) + 1
	return integerReturn
}
