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
	fmt.Scan(&numPlanets)
	fmt.Println("you selected ", numPlanets)
	fmt.Println("Give me an int for numConnections please\n")
	fmt.Scan(&numConnections)
	fmt.Println("you selected ", numConnections)

	fmt.Println("Give me a string to read please\n")
	fmt.Scan(&edges)
	fmt.Println("You gave me ", edges)

}

