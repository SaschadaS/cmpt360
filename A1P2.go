
/*
Sascha da Silva
snd809
11040574

Assignment 1 Programming question 2

Identifying virus spread in a list of computer interactions.
 */
package main

import (
	"fmt"
	"bufio"
	"os"
)

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
var tempLine string
fmt.Scanln(&numComputers)
fmt.Scanln(&numTrace)
isInfected := make([]bool, numComputers)
traceData := make([]trace, numTrace)
scanner := bufio.NewReader(os.Stdin)

for i, a := range traceData {
	tempLine, _ = scanner.ReadString('\n')
	fmt.Sscanf(tempLine, "(C%d, C%d, %d)", &a.source, &a.destination, &a.time)
	traceData[i] = a
}
//fmt.Scanln(&numQueries)
tempLine, _ = scanner.ReadString('\n')
fmt.Sscanf(tempLine,"%d", &numQueries)
queryData := make([]query, numQueries)
queryOutput := make([]string, numQueries)
for j, b := range queryData {
	tempLine, _ = scanner.ReadString('\n')
	fmt.Sscanf(tempLine, "C%d, %d, C%d, %d", &b.startComp, &b.startTime, &b.endComp, &b.endTime)
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
		queryOutput[j] = "Y"
	} else {
		queryOutput[j] = "N"
	}
	for k := 0; k < numComputers; k++{
		isInfected[k] = false
	}
}

for l := 0; l < numQueries; l++ {
	fmt.Println(queryOutput[l])
}

}

