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

func setTransits(arrayPlanets []Planet, arrayTransits []Edge) []Planet{

	for i, p := range arrayPlanets{
		for _, p2 := range arrayPlanets{
			for _, t := range arrayTransits{
			//fmt.Println(p.name, t.startPlanet, t.endPlanet)
			if p.name == t.startPlanet && p2.name == t.endPlanet{
				arrayPlanets[i].numTransits = arrayPlanets[i].numTransits + 1
				arrayPlanets[i].children = append(arrayPlanets[i].children, p2)
			}


			}
		}

	}
	return arrayPlanets
}


func main() {

	var (
	successRead bool
	successFind []Planet
	root int
	)
	/*Tests begin */


	/* Can we read values from user? */
	var numPlanets, numConnections int

	fmt.Scan(&numPlanets)
	/* array of planets */
	var slicePlanets = make([]Planet,numPlanets)
	for i := 0; i < numPlanets; i++ {
		fmt.Scan(&slicePlanets[i].name)
		if (slicePlanets[i].name == "Scarif"){
			slicePlanets[i].cost = 0
			root = i
		} else if( slicePlanets[i].name == "Yavin") {
			slicePlanets[i].cost = 0
		} else{
			fmt.Scan(&slicePlanets[i].cost)
		}
	}

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

	successFind = setTransits(slicePlanets, sliceTransits)
	fmt.Println(successFind[root])

	target := searchTransits(sucessFind)
	/* Now trying to search and find the shortest path */
	//if successFind.name != "Scarif" {
	//	fmt.Println("Darth Blockades ", successFind.name)
	//} else{
	//	fmt.Println("Leia escapes with the plans!")
//	}
}

