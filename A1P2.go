
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

func readTrace(numTrace int) []trace {
	var TraceData []trace
	for i := 0; i<numTrace; i++ {
	fmt.Scan(TraceData[i].source, TraceData[i].destination, TraceData[i].time)
	}
	return TraceData
}

func readQuery(numQueries int) []query {
	var QueryData []query
	for i := 0; i<numQueries; i++ {
	fmt.Scan(QueryData[i].startComp, QueryData[i].startTime, QueryData[i].endComp, QueryData[i].endTime)
	}
	return QueryData
}

func checkQueries(queryData []query, traceData []trace, isInfected []bool){
for _, y := range queryData {

	if isInfected[y.endComp+1] == true{
	fmt.Println("Y")
	}else {
	fmt.Println("N")
	}
}
}

func main(){

var numComputers, numTrace, numQueries int
fmt.Scan(numComputers, numTrace)
var isInfected []bool
var traceData []trace =readTrace(numTrace)
fmt.Scan(numQueries)
var queryData []query =readQuery(numQueries)

checkQueries(queryData, traceData, isInfected)

fmt.Println("Compiled successfully!")

}
