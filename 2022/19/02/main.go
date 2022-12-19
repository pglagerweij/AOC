package main

import (
	"fmt"
	"os"
	"strings"
)

// var inputFile string = "trial.txt"

var inputFile string = "input.txt"

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

	totalResult := 1
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var oreRobotCost Materials = Materials{0, 0, 0, 0}
		var clayRobotCost Materials = Materials{0, 0, 0, 0}
		var obsidianRobotCost Materials = Materials{0, 0, 0, 0}
		var geodeRobotCost Materials = Materials{0, 0, 0, 0}
		var maxCost Materials
		var robots Materials = Materials{1, 0, 0, 0}
		var currentStock Materials = Materials{0, 0, 0, 0}
		var blueprint int
		var maxTime int = 33

		choices := make(map[string]Choice)

		fmt.Sscanf(s, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &blueprint, &oreRobotCost.ore, &clayRobotCost.ore, &obsidianRobotCost.ore, &obsidianRobotCost.clay, &geodeRobotCost.ore, &geodeRobotCost.obsidian)
		choices["nothing"] = Choice{cost: Materials{0, 0, 0, 0}, robots: Materials{0, 0, 0, 0}}
		choices["ore"] = Choice{cost: oreRobotCost, robots: Materials{1, 0, 0, 0}}
		maxCost = oreRobotCost
		choices["clay"] = Choice{cost: clayRobotCost, robots: Materials{0, 1, 0, 0}}
		maxCost = Materials{Max(maxCost.ore, clayRobotCost.ore), Max(maxCost.clay, clayRobotCost.clay), Max(maxCost.obsidian, clayRobotCost.obsidian), 0}
		choices["obsidian"] = Choice{cost: obsidianRobotCost, robots: Materials{0, 0, 1, 0}}
		maxCost = Materials{Max(maxCost.ore, obsidianRobotCost.ore), Max(maxCost.clay, obsidianRobotCost.clay), Max(maxCost.obsidian, obsidianRobotCost.obsidian), 0}
		choices["geode"] = Choice{cost: geodeRobotCost, robots: Materials{0, 0, 0, 1}}
		maxCost = Materials{Max(maxCost.ore, geodeRobotCost.ore), Max(maxCost.clay, geodeRobotCost.clay), Max(maxCost.obsidian, geodeRobotCost.obsidian), 0}

		for ind, v := range choices {
			fmt.Printf("%v\n", ind)
			fmt.Printf("%v\n", v.cost)
			fmt.Printf("%v\n", v.robots)
		}

		fmt.Printf("maxcost are %v\n", maxCost)

		var time int = 1
		fmt.Printf("This is blueprint %v\n", blueprint)
		choicemap := []string{}
		totalGeode, _ := solveRecur2(choices, currentStock, robots, time, maxTime, choicemap, Materials{0, 0, 0, 0}, maxCost)

		fmt.Printf("for blueprint %v the maximum number of Geode is %v\n", blueprint, totalGeode)
		// fmt.Printf("The choices are %v\n", finalchoice)
		totalResult = totalResult * totalGeode
	}
	fmt.Printf("The total result is %v", totalResult)
}

func solveRecur2(choices map[string]Choice, currentStock Materials, robots Materials, time int, limit int, choicemap []string, lastStock Materials, maxCost Materials) (int, []string) {

	geodeScore := currentStock.geode + (limit-time)*robots.geode
	max := geodeScore
	finalChoice := choicemap
	// fmt.Printf("the time is : %v\n", time)
	if time < limit {
		for choice, v := range choices {
			checkmax := robots.Add(v.robots).Min(maxCost)
			if checkmax.clay > 0 || checkmax.ore > 0 || checkmax.obsidian > 0 {
				continue
			}
			newStock, result := buyRobot(v.cost, currentStock, choicemap, lastStock)
			if result == false {
				continue
			}
			// fmt.Printf("We bough robot %v and have now a stock of %v \n", choice, newStock)
			stock := newStock.Add(robots)
			// fmt.Printf("we are now collecting materials %v, current stock is %v\n", robots, stock)
			robotsinloop := robots.Add(v.robots)
			// fmt.Printf("we added %v robots, we now have %v robots\n", v.robots, robotsinloop)
			choicesMade := append(choicemap, choice)
			// fmt.Printf("the time is : %v with choicemap %v\n", time, choicesMade)
			possibleScore, choice := solveRecur2(choices, stock, robotsinloop, time+1, limit, choicesMade, currentStock, maxCost)
			if possibleScore > max {
				max = possibleScore
				finalChoice = choice
			}
		}
	}
	// fmt.Printf("at the end of the time limit %v\nCurrent stock of geodes is %v we added %v we now have %v\n", time, currentStock.ore, (limit-time)*robots.ore, geodeScore)
	return max, finalChoice
}

func buyRobot(in Materials, currentStock Materials, choicemap []string, lastStock Materials) (Materials, bool) {
	newStock := currentStock.Min(in)
	LatestoptionStock := lastStock.Min(in)
	nothing := Materials{0, 0, 0, 0}
	result := newStock.clay >= 0 && newStock.ore >= 0 && newStock.obsidian >= 0
	resultLatest := LatestoptionStock.clay >= 0 && LatestoptionStock.ore >= 0 && LatestoptionStock.obsidian >= 0
	var lastturn string
	if len(choicemap) != 0 {
		lastturn = choicemap[len(choicemap)-1]
		// Don't take an action if you could take it last turn and it is not nothing
		if resultLatest && lastturn == "nothing" && in != nothing {
			return currentStock, false
		} else if result {
			return newStock, true
		} else {
			return currentStock, false
		}
	} else {
		if result {
			return newStock, true
		} else {
			return currentStock, false
		}
	}

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
