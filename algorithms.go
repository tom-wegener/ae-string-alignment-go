package main

import (
	"fmt"
	"time"
)

func compareFiles(dataA []record, dataB []record) (runTimes [][]int) {
	fmt.Println(dataA[1].seq)
	fmt.Println("-------------------------------------------------")
	runNums := []int{10, 15, 20, 25, 30}
	runTimes = make([][]int, 0)
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
		runTime := endTime.Sub(startTime)
		runTimesL := []int{runNum, int(runTime)}
		runTimes = append(runTimes, runTimesL)
		runTimes[i][1] = int(runTime.Seconds())
	}
	return runTimes
}

func needlemanWunsch(seqA, seqB string) {

	a := len(seqA) + 1
	b := len(seqB) + 1
	match := 1
	mismatch := -1
	gap := -1

	numMat := make([][]int, a)
	for i := range numMat {
		numMat[i] = make([]int, b)
	}
	arrMat := make([][]int, a)
	for i := range arrMat {
		arrMat[i] = make([]int, b)
	}

	for i := 0; i < a; i++ {
		numMat[i][0] = mismatch * i
	}
	for j := 0; j < b; j++ {
		numMat[0][j] = mismatch * j
	}

	for i := 1; i < a; i++ {
		for j := 1; j < b; j++ {
			score := 0
			if seqA[i-1] == seqB[j-1] {
				score = match
			} else {
				score = mismatch
			}
			resMatch := numMat[i-1][j-1] + score
			resDel := numMat[i-1][j] + gap
			resIns := numMat[i][j-1] + gap
			numMat[i][j] = min(resMatch, resDel, resIns)
		}
	}

}

func min(is ...int) int {
	min := is[0]
	for _, i := range is[1:] {
		if i < min {
			min = i
		}
	}
	return min
}
