
/*
Sascha da Silva
snd809
11040574

Assignment 1 Programming question 2

Identifying virus spread in a list of computer interactions.
 */
package main

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

var numComputers, numTrace, numQueries int
fmt.Scanln(&numComputers)
fmt.Scanln(&numTrace)
isInfected := make([]bool, numComputers)
traceData := make([]trace, numTrace)
for i, a := range traceData {
	fmt.Scanln(&a.source, &a.destination, &a.time)
	traceData[i] = a
}

fmt.Scanln(&numQueries)
queryData := make([]query, numQueries)
for j, b := range queryData {
	fmt.Scanln(&b.startComp, &b.startTime, &b.endComp, &b.endTime)
	queryData[j] = b
	isInfected[b.startComp-1] = true
	for _, d := range traceData {
		if b.startTime <= d.time && d.time <= b.endTime {
			if isInfected[d.source-1] == true || isInfected[d.destination-1] == true {
				isInfected[d.destination-1] = true
				isInfected[d.source-1] = true
			}
		}
	}
	if isInfected[b.endComp-1] == true {
		fmt.Println("Y")
	} else {
		fmt.Println("N")
	}
	for k := 0; k < numComputers; k++{
		isInfected[k] = false
	}
}

fmt.Println(isInfected, traceData, queryData)
fmt.Println("Compiled successfully!")

}

