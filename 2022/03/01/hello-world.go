package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	total_priority := 0
	for _, element := range s3 {
		// fmt.Printf("%v\n", element)
		element1 := element[:len(element)/2]
		element2 := element[len(element)/2:]
		// fmt.Printf("%v\n", element1)
		// fmt.Printf("%v\n", element2)
		var required1 []string
		var required2 []string
		for _, k := range element1 {
			// fmt.Printf("%c\n", k)
			required1 = append(required1, strconv.QuoteRune(k))
		}
		sort.Strings(required1)
		set_1 := removeDuplicateStr(required1)
		// fmt.Printf("%v\n", set_1)
		for _, k := range element2 {
			// fmt.Printf("%c\n", k)
			required2 = append(required2, strconv.QuoteRune(k))
		}
		sort.Strings(required2)
		set_2 := removeDuplicateStr(required2)
		// fmt.Printf("%v\n", set_2)
		intersection := intersect.Simple(set_1, set_2)
		stringRepresentation := fmt.Sprintf("%v", intersection[0])
		stringRepresentation2 := fmt.Sprintf("%c", stringRepresentation[1])
		priority := getPriority(stringRepresentation2)
		// fmt.Printf("The intersection of element %v and element %v = %v with priority %v\n", element1, element2, intersection[0], priority)

		total_priority += priority
	}
	fmt.Printf("The total priority is: %v\n", total_priority)
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
