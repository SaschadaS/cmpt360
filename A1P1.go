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
	children []Planet
	numTransits int
}

type Tree struct{
	root Planet
	curCost int
}

/* Darth Vader's least expensive blockade program */

func readPlanets() []Planet{

	var numPlanets int

	fmt.Println("Give me an int for numPlanets please")
	fmt.Scan(&numPlanets)
	/* array of planets */
	var slicePlanets = make([]Planet,numPlanets)
	for i := 0; i < numPlanets; i++ {
		fmt.Scan(&slicePlanets[i].name)
		fmt.Scan(&slicePlanets[i].cost)
	}
	return slicePlanets
}

func readTransits() []Edge{

	var numConnections int
	fmt.Println("Give me an int for numConnections please")
	fmt.Scan(&numConnections)
	/* array of edges */
	var sliceTransits = make([]Edge,numConnections)
	for j := 0; j < numConnections; j++ {
		fmt.Scan(&sliceTransits[j].startPlanet)
		fmt.Scan(&sliceTransits[j].endPlanet)
	}
	return sliceTransits

}

func searchTransits(arrayPlanets []Planet, arrayTransits []Edge) Planet{
	var target Planet

	for _, p := range arrayPlanets{
		for _, p2 := range arrayPlanets{
			for _, t := range arrayTransits{

			if p.name == t.startPlanet && p2.name == t.endPlanet{
				p.numTransits++
				p.children = append(p.children,p2)
			}


			}
		}
	}


	return target
}


func main() {

	var (
	successRead bool
	successFind Planet
	)
	/*Tests begin */


	/* Can we read values from user? */
	var numPlanets, numConnections int

//	fmt.Println("Give me an int for numPlanets please")
	fmt.Scan(&numPlanets)
	/* array of planets */
	var slicePlanets = make([]Planet,numPlanets)
	for i := 0; i < numPlanets; i++ {
		fmt.Scan(&slicePlanets[i].name)
		if (slicePlanets[i].name == "Scarif" || slicePlanets[i].name == "Yavin") {
			slicePlanets[i].cost = 0
		} else{
		fmt.Scan(&slicePlanets[i].cost)
		}
	}

//	fmt.Println("Give me an int for numConnections please")
	fmt.Scan(&numConnections)
	/* array of edges */
	var sliceTransits = make([]Edge,numConnections)
	for j := 0; j < numConnections; j++ {
		fmt.Scan(&sliceTransits[j].startPlanet)
		fmt.Scan(&sliceTransits[j].endPlanet)
	}
	if successRead == true {
	fmt.Println("filler")
	}

	fmt.Println(slicePlanets)
	fmt.Println(sliceTransits)

	successFind = searchTransits(slicePlanets, sliceTransits)

	/* Now trying to search and find the shortest path */
	if successFind.name != "Scarif" {
	fmt.Println("Darth Blockades ", successFind.name)
	} else{
	fmt.Println("Leia escapes with the plans!")
	}
}

