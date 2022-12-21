package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var inputFile string = "trial.txt"

var inputFile string = "input.txt"

var monkeys = map[string]string{}

func main() {
	input, _ := os.ReadFile(inputFile)

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		// fmt.Printf("monkeys is %v\n", s)
		s3 := strings.Split(s, ": ")
		monkeys[s3[0]] = strings.TrimSpace(s3[1])
		// fmt.Printf("we are testing monkeys is %v.\n", monkeys)
	}

	fmt.Println(solve("root"))
}

func solve(expr string) int {
	if v, err := strconv.Atoi(monkeys[expr]); err == nil {
		return v
	}

	s := strings.Fields(monkeys[expr])
	return map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}[s[1]](solve(s[0]), solve(s[2]))
}
