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

	for i, p := range arrayPlanets{
		for j, p2 := range arrayPlanets{
			for _, t := range arrayTransits{

			if p.name == t.startPlanet && p2.name == t.endPlanet{
				arrayPlanets[i].numTransits = arrayPlanets[i].numTransits + 1
				arrayPlanets[j].numTransits = arrayPlanets[j].numTransits + 1
				arrayPlanets[i].children = append(arrayPlanets[i].children, p2.id)
				arrayPlanets[j].children = append(arrayPlanets[j].children, p.id)
			}

			}

		}
	}
	return arrayPlanets
}


func DFS(root int, arrayPlanets []Planet)int{
	goal := -1
	var isArtPoint bool
	var x int

	fmt.Println("got to dfs")

	arrayPlanets[root].found = true
	for len(arrayPlanets[root].children) > 0 {
	x = arrayPlanets[root].children[0]
	arrayPlanets[root].children = arrayPlanets[root].children[1:]

	fmt.Println("Dfsing: ",arrayPlanets[root], "with: ", arrayPlanets[x])

	if (arrayPlanets[x].found == false) {
	isArtPoint = checkArticulation(arrayPlanets, x, root)

	if isArtPoint == true && goal == -1 {
		goal = x
	} else if arrayPlanets[x].cost < arrayPlanets[goal].cost{
		goal = x
	}

	DFS(x, arrayPlanets)
	}
	}


	return goal
}

func checkArticulation(arrayPlanets []Planet,  id int, parent int)bool{
	var isArticulation bool
	fmt.Println("got to articulation. id =", id, " parent=", parent)

	if(arrayPlanets[id].numTransits <= 1) {
		fmt.Println("hit leaf")
		return false
	}

	for i := 0; i < arrayPlanets[id].numTransits; i++ {
		fmt.Println("Articulating: ", arrayPlanets[id], "with: ", arrayPlanets[id].children[i], "i =",i, " parent= ", parent)

		if (parent == i) {
			fmt.Println("hit Parent")
		}

		if (arrayPlanets[arrayPlanets[id].children[i]].found==true && i != parent) {
			return false
		}else if (i != parent) {
			isArticulation = checkArticulation(arrayPlanets, arrayPlanets[id].children[i], id)
		}

		if (i == arrayPlanets[id].numTransits-1) {
			return true
		}
	}
	return isArticulation
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

	/* prints values for each planet */
	for l := 0; l < numPlanets; l++ {
	fmt.Println("Record:", successFind[l])
	}


	if target != root {
		fmt.Println("Darth Blockades ", slicePlanets[target].name)
	} else{
		fmt.Println("Leia escapes with the plans!")
	}
}

