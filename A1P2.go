
/*
Sascha da Silva
snd809
11040574

Assignment 1 Programming question 2

Identifying virus spread in a list of computer interactions.
 */

import "fmt"

type trace struct{
	source int
	destination int
	time int
}

type query struct{
	startComp int
	startTime int
	endComp int
	endTime int
}

func main(){

var isInfected []bool
var numComputers, numTrace, numQueries int

var traceData []trace =readTrace()
var queryData []query =readQuery()

checkQueries(comp1 int, trace1 int, comp2 int, trace2 int)

fmt.Println("Compiled successfully!")

}
