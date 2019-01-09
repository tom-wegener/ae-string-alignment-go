package main

import (
	"time"
)

func compareFiles(dataA []record, dataB []record) {
	runNums := []int{10, 20, 30, 40, 50, 60, 70, 80}
	runTimes := make([][]int, 8)
	for i, runNum := range runNums {
		startTime := time.Now()
		for x := 0; x < runNum; x++ {
			for y := 0; y < runNum; y++ {
				seqA := dataA[x].seq
				seqB := dataB[y].seq
				needlemanWunsch(seqA, seqB)
			}
		}
		endTime := time.Now()
		runTime := startTime.Sub(endTime)
		runTimes[i][0] = runNum
		runTimes[i][1] = int(runTime)
	}
}

func needlemanWunsch(seqA, seqB string) {

}
