package main

import (
	"time"
)

func compareFiles(dataA []record, dataB []record) (runTimes [][]int) {
	runNums := []int{10, 15, 20, 25, 30, 35, 40}
	runTimes = make([][]int, 0)

	for _, runNum := range runNums {
		seqLen := 0
		for x := 0; x < runNum; x++ {
			for y := 0; y < runNum; y++ {
				startTime := time.Now()

				seqA := dataA[x].seq
				seqB := dataB[y].seq
				needlemanWunsch(seqA, seqB)
				seqLen = len(seqA) * len(seqB)

				endTime := time.Now()
				runTime := endTime.Sub(startTime)
				runTimesL := []int{seqLen, int(runTime.Seconds())}
				runTimes = append(runTimes, runTimesL)
			}
		}
		/*endTime := time.Now()
		runTime := endTime.Sub(startTime)
		runTimesL := []int{seqLen, int(runTime.Seconds())}
		runTimes = append(runTimes, runTimesL)
		//runTimes[i][1] = int(runTime.Seconds())*/
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
			numMat[i][j] = minA(resMatch, resDel, resIns)
		}
	}

}

func minA(is ...int) int {
	min := is[0]
	for _, i := range is[1:] {
		if i < min {
			min = i
		}
	}
	return min
}

func minB(a, b, c int) int {
	if a < b && b < c {
		return a
	} else if a > b && b > c {
		return c
	} else {
		return b
	}
}
