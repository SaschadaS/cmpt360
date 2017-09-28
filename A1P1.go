package main

import (
	"fmt"
)

type Edge struct {
	startPlanet string
	endPlanet string
}

type Planet struct {
	id int
	name string
	cost int
	children []int
	parents []int
	numTransits int
	found bool
}



/* Darth Vader's least expensive blockade program */

func setTransits(arrayPlanets []Planet, arrayTransits []Edge) []Planet{
	//var source, dest Planet

	for i, p := range arrayPlanets{
		for j, p2 := range arrayPlanets{
			for _, t := range arrayTransits{

			if p.name == t.startPlanet && p2.name == t.endPlanet{
				arrayPlanets[i].numTransits = arrayPlanets[i].numTransits + 1
				arrayPlanets[j].numTransits = arrayPlanets[j].numTransits + 1
				//source = arrayPlanets[i]
				//dest = arrayPlanets[j]
				arrayPlanets[i].children = append(arrayPlanets[i].children, p2.id)
				arrayPlanets[j].children = append(arrayPlanets[j].children, p.id)
			}

			}

		}
	}
	return arrayPlanets
}


func DFS(root int, arrayPlanets []Planet)int{
	var goal int
	var isArtPoint bool
	var x int

	arrayPlanets[root].found = true
	for len(arrayPlanets[root].children) > 0 {

	x = arrayPlanets[root].children[0]
	arrayPlanets[root].children = arrayPlanets[root].children[1:]

	if (arrayPlanets[x].found == false) {
	isArtPoint = checkArticulation(arrayPlanets, x)

	if isArtPoint == true && arrayPlanets[x].cost < arrayPlanets[goal].cost{
		goal = x
	}

	DFS(x, arrayPlanets)
	}
	}


	return goal
}

func checkArticulation(arrayPlanets []Planet,  id int)bool{
	for i := 0; i < arrayPlanets[id].numTransits; i++ {
		if arrayPlanets[arrayPlanets[id].children[i]].found==true {
		return false
		}else{
		checkArticulation(arrayPlanets, arrayPlanets[id].children[i])
		return true
		}
	}
	return false
}

func main(){

	var (
	successFind []Planet
	root, numPlanets, numConnections, target int
	// slicePlanets will be a slice of all given planets
	// sliceTransits will be a slice of all given transits
	)

	/*Tests begin */
	/* read number of planets*/
	fmt.Scan(&numPlanets)
	/* set array of planets to size numPlanets */
	var slicePlanets = make([]Planet,numPlanets)
	/* reads in n planets into slicePlanets, where n is
	size of numPlanets  */
	for i := 0; i < numPlanets; i++ {
		fmt.Scan(&slicePlanets[i].name)
		slicePlanets[i].id = i
		if (slicePlanets[i].name == "Scarif"){
			slicePlanets[i].cost = 0
			root = i
		} else if( slicePlanets[i].name == "Yavin") {
			slicePlanets[i].cost = 0
		} else{
			fmt.Scan(&slicePlanets[i].cost)
		}
	}

	/* reads in numConnections from stdin */
	fmt.Scan(&numConnections)
	/* set array of transits to size numConnections */
	var sliceTransits = make([]Edge,numConnections)
	/* reads in n transits into sliceTransits, where n
	is of numConnections */
	for j := 0; j < numConnections; j++ {
		fmt.Scan(&sliceTransits[j].startPlanet)
		fmt.Scan(&sliceTransits[j].endPlanet)
	}

	/* sets the values of the Planets for searching */
	successFind = setTransits(slicePlanets, sliceTransits)

	/* prints values for each planet */
	for l := 0; l < numPlanets; l++ {
	fmt.Println("Record:", successFind[l])
	}

	/* sets int to planet.id of best articulation point, 
	or 0 if no point exists. */
	target = DFS(root, successFind)

	if target != root {
		fmt.Println("Darth Blockades ", slicePlanets[target].name)
	} else{
		fmt.Println("Leia escapes with the plans!")
	}
}

