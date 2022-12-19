package main

import (
	"fmt"
	"os"
	"strings"
)

var inputFile string = "trial.txt"

// var inputFile string = "input.txt"

var maxSteps int = 24

type Materials struct {
	ore, clay, obsidian, geode int
}

type Choice struct {
	cost   Materials
	robots Materials
}

func (p Materials) Add(q Materials) Materials {
	return Materials{p.ore + q.ore, p.clay + q.clay, p.obsidian + q.obsidian, p.geode + q.geode}
}

func (p Materials) Min(q Materials) Materials {
	return Materials{p.ore - q.ore, p.clay - q.clay, p.obsidian - q.obsidian, p.geode - q.geode}
}

func main() {
	input, _ := os.ReadFile(inputFile)
	// convert the file binary into a string using string
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var oreRobotCost Materials = Materials{0, 0, 0, 0}
		var clayRobotCost Materials = Materials{0, 0, 0, 0}
		var obsidianRobotCost Materials = Materials{0, 0, 0, 0}
		var geodeRobotCost Materials = Materials{0, 0, 0, 0}

		var robots Materials = Materials{1, 0, 0, 0}
		var currentStock Materials = Materials{0, 0, 0, 0}
		var blueprint int
		var maxTime int = 24

		choices := make(map[string]Choice)

		fmt.Sscanf(s, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &blueprint, &oreRobotCost.ore, &clayRobotCost.ore, &obsidianRobotCost.ore, &obsidianRobotCost.clay, &geodeRobotCost.ore, &geodeRobotCost.obsidian)
		choices["nothing"] = Choice{cost: Materials{0, 0, 0, 0}, robots: Materials{0, 0, 0, 0}}
		choices["ore"] = Choice{cost: oreRobotCost, robots: Materials{1, 0, 0, 0}}
		choices["clay"] = Choice{cost: clayRobotCost, robots: Materials{0, 1, 0, 0}}
		choices["obsidian"] = Choice{cost: obsidianRobotCost, robots: Materials{0, 0, 1, 0}}
		choices["geode"] = Choice{cost: geodeRobotCost, robots: Materials{0, 0, 0, 1}}
		for ind, v := range choices {
			fmt.Printf("%v\n", ind)
			fmt.Printf("%v\n", v.cost)
			fmt.Printf("%v\n", v.robots)
		}

		var time int = 1
		fmt.Printf("This is blueprint %v\n", blueprint)
		choicemap := []string{}
		totalGeode, finalchoice := solveRecur(choices, currentStock, robots, time, maxTime, choicemap)

		fmt.Printf("for blueprint %v the maximum number of Geode is %v\n", blueprint, totalGeode)
		fmt.Printf("The choices are %v\n", finalchoice)
		os.Exit(3)
	}

}

func solveRecur(choices map[string]Choice, currentStock Materials, robots Materials, time int, limit int, choicemap []string) (int, []string) {

	geodeScore := currentStock.geode + (limit-time)*robots.geode
	max := geodeScore
	final_choice := choicemap
	// fmt.Printf("the time is : %v\n", time)
	if time < limit {
		for choice, v := range choices {
			// fmt.Printf("%v\n", choice)
			var result bool
			currentStock, result = buyRobot(v.cost, currentStock)
			if result == false {
				// fmt.Printf("Cannot do choice %v\n", choice)
				continue
			}
			// fmt.Printf("we bought robots %v\n", choice)
			currentStock = currentStock.Add(robots)
			// fmt.Printf("we are now collecting materials %v, current stock is %v\n", robots, currentStock)
			robots = robots.Add(v.robots)
			// fmt.Printf("we added  	 %v robots, we now have %v robots\n", v.robots, robots)
			var possibleScore int
			var finalchoice []string
			choicemap_run := append(choicemap, choice)
			possibleScore, finalchoice = solveRecur(choices, currentStock, robots, time+1, limit, choicemap_run)

			if possibleScore > max {
				fmt.Printf("the time is : %v with choicemap %v\n", time, finalchoice)
				max = possibleScore
				final_choice = finalchoice
			}
		}
	}

	return max, final_choice
}

func buyRobot(in Materials, currentStock Materials) (Materials, bool) {
	newStock := currentStock.Min(in)
	nothing := Materials{0, 0, 0, 0}
	result := newStock.clay >= 0 && newStock.ore >= 0 && newStock.obsidian >= 0
	if (in.clay == currentStock.clay && currentStock.clay != 0 && result) || (in.ore == currentStock.ore && currentStock.ore != 0 && result) || (in.obsidian == currentStock.obsidian && currentStock.obsidian != 0 && result) || (in == nothing && result) {
		return newStock, true
	} else {
		return currentStock, false
	}
}
