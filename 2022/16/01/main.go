package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// var inputFile string = "trial.txt"

var inputFile string = "input.txt"

var maxSteps int = 30

type graph struct {
	to string
	wt int
}

func main() {

	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	s := strings.Split(fileContent, "\n")

	var flows map[string]int = make(map[string]int)
	valves := make([]string, len(s))

	graf := make(map[string][]graph, len(valves))
	for index, element := range s {
		splitted := strings.Split(element, "; ")
		l := splitted[0]
		lsplit := strings.Split(l, "=")
		flowrate, _ := strconv.Atoi(strings.TrimSpace(lsplit[1]))
		valvenumber := l[6:8]
		valves[index] = valvenumber
		// fmt.Printf("in valve %v the flow is %v\n", valvenumber, flowrate)
		flows[valvenumber] = flowrate

		// Right side
		r := splitted[1]
		graphss := []graph{}
		if strings.HasPrefix(r, "tunnels lead to") {
			rsplit := strings.Split(r, "valves ")
			to_valves := strings.Split(strings.TrimSpace(rsplit[1]), ", ")
			// fmt.Printf("%v\n", to_valves)

			for _, elem := range to_valves {
				graphss = append(graphss, graph{elem, 1})
			}
			graf[valvenumber] = graphss
		} else if strings.HasPrefix(r, "tunnel leads to") {
			rsplit := strings.Split(r, "valve ")
			to_valve := strings.TrimSpace(rsplit[1])
			graphss = append(graphss, graph{to_valve, 1})
			graf[valvenumber] = graphss
		} else {
			panic("canot find match")
		}

	}

	// Algorithm to compute shortest part between all matrix elements
	g := graf
	dist := make(map[string]map[string]int, len(g))
	for i := range g {
		di := make(map[string]int, len(g))
		for j := range g {
			di[j] = 900000
		}
		di[i] = 0
		dist[i] = di
	}
	// fmt.Printf("%v\n", dist)
	for u, graphs := range g {
		for _, v := range graphs {
			dist[u][v.to] = v.wt
		}
	}
	// fmt.Printf("%v\n", dist)
	for k, dk := range dist {
		for _, di := range dist {
			for j, dij := range di {
				if d := di[k] + dk[j]; dij > d {
					di[j] = d
				}
			}
		}
	}

	// Get all flows that have presuure
	relevantFlows := []string{}
	for key, elem := range flows {
		if elem != 0 {
			relevantFlows = append(relevantFlows, key)
		}
	}
	fmt.Printf("%v\n", relevantFlows)

	// Solve recursive
	res := solveRecur(dist, flows, 0, 0, 0, "AA", relevantFlows, maxSteps)
	fmt.Printf("%v\n", res)
}

func solveRecur(matrix map[string]map[string]int, pressures map[string]int, currentTime int, currentPressure int, currentFlow int, currentTunnel string, remaining []string, limit int) int {
	// The score if no other valves are being opened before the $limit time
	nScore := currentPressure + (limit-currentTime)*currentFlow
	max := nScore

	for _, v := range remaining {
		distanceAndOpen := matrix[currentTunnel][v] + 1
		if currentTime+distanceAndOpen < limit {
			newTime := currentTime + distanceAndOpen
			newPressure := currentPressure + distanceAndOpen*currentFlow
			newFlow := currentFlow + pressures[v]
			possibleScore := solveRecur(matrix, pressures, newTime, newPressure, newFlow, v, removeFromList(remaining, v), limit)
			if possibleScore > max {
				max = possibleScore
			}
		}
	}

	return max
}

func removeFromList(in []string, v string) []string {
	new := []string{}
	for _, i := range in {
		if i != v {
			new = append(new, i)
		}
	}
	return new
}
