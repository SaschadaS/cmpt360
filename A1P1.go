package main

import (
	"fmt"
)

/* Darth Vader's least expensive blockade program */

func main() {
	/*Tests begin */

	/* Can we read values from user? */
	var numPlanets, numConnections int
	var edges string

	fmt.Println("Give me an int for numPlanets please\n")
	fmt.Scanln(numPlanets)
	fmt.Println("you selected %d", numPlanets)
	fmt.Println("Give me an int for numConnections please\n")
	fmt.Scanln(numConnections)
	fmt.Println("you selected %d", numConnections)

	fmt.Println("Give me a string to read please\n")
	fmt.Scanln(edges)

}

