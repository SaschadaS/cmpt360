package main

import (
	"fmt"
)

type Edge struct {
	startPlanet string
	endPlanet string
}

type Planet struct {
	name string
	cost int
	parent Planet
	numTransits int
}

type Tree struct{
	root Planet
	curCost int
}

/* Darth Vader's least expensive blockade program */

func readValues() bool{

	/* Can we read values from user? */
	var numPlanets, numConnections int

	fmt.Println("Give me an int for numPlanets please")
	fmt.Scan(&numPlanets)
	/* array of planets */
	var slicePlanets = make([]Planet,numPlanets)
	for i := 0; i < numPlanets; i++ {
		fmt.Scan(&slicePlanets.name[i])
		fmt.Scan(&slicePlanets.cost[i])
	}

	fmt.Println("Give me an int for numConnections please")
	fmt.Scan(&numConnections)
	/* array of edges */
	var sliceTransits = make([]Edge,numConnections)
	for j := 0; j < numConnections; j++ {
		fmt.Scan(&sliceTransits[j].startPlanet)
		fmt.Scan(&sliceTransits[j].endPlanet)
	}
}

func main() {
	var sucessRead
	var successFind Planet
	var targetPlanet Planet
	/*Tests begin */
	successRead = readValues()

	fmt.Println(slicePlanets)
	fmt.Println(sliceTransits)

	success2 = searchTransits()


	/* Now trying to search and find the shortest path */
	if successFind = nil {
	fmt.Println("Darth Blockades ", targetPlanet.name)
	} else{
	fmt.Println("Leia escapes with the plans!")
}

