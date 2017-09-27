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
	var transit1 Edge

	fmt.Println("Give me an int for numPlanets please")
	fmt.Scan(&numPlanets)
	fmt.Println("you selected ", numPlanets)
	/* array of planets */
	var planets [numPlanets]string
	fmt.Println("Give me an int for numConnections please")
	fmt.Scan(&numConnections)
	/* array of edges */
	var transits [numConnections]Edge
	fmt.Println("you selected ", numConnections)

	fmt.Println("Give me a string to read please")
	fmt.Scan(&transit1.startPlanet)
	fmt.Scan(&transit1.endPlanet)
	transits[0] = transit1
	fmt.Println("You gave me ", transit1)

}

