package main

import (
	"fmt"
)

type Edge struct {
	startPlanet string
	endPlanet string
}


/* Darth Vader's least expensive blockade program */

func main() {
	/*Tests begin */

	/* Can we read values from user? */
	var numPlanets, numConnections int

	fmt.Println("Give me an int for numPlanets please")
	fmt.Scan(&numPlanets)
	fmt.Println("you selected ", numPlanets)
	/* array of planets */
	var slicePlanets = make([]string,numPlanets)
	for i := 0; i < numPlanets; i++ {
		fmt.Scan(&slicePlanets[i])
	}

	fmt.Println("Give me an int for numConnections please")
	fmt.Scan(&numConnections)
	/* array of edges */
	var sliceTransits = make([]Edge,numConnections)
	fmt.Println("you selected ", numConnections)
	for j := 0; j < numConnections; j++ {
		fmt.Scan(&sliceTransits[j].startPlanet)
		fmt.Scan(&sliceTransits[j].endPlanet)
	}

	fmt.Println(slicePlanets)
	fmt.Println(sliceTransits)
}

