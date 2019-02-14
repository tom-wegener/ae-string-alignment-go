package main

import (
	"time"
)

func compareFiles(dataA []record, dataB []record, alg string) (runTimes runTimesArr) {
	runNum := 100
	runTimes = make(runTimesArr, 0)

	if alg == "needleman-wunsch" {
		for x := 0; x < runNum; x++ {
			for y := 0; y < runNum; y++ {
				var timeRow times
				startTime := time.Now()

				seqA := dataA[x].seq
				seqB := dataB[y].seq
				needlemanWunsch(seqA, seqB)

				endTime := time.Now()

				timeRow.strLen = (len(seqA) * len(seqB))
				timeRow.runTime = int(endTime.Sub(startTime))
				runTimes = append(runTimes, timeRow)
			}
		}
	} else {
		for x := 0; x < runNum; x++ {
			for y := 0; y < runNum; y++ {
				var timeRow times
				startTime := time.Now()

				seqA := dataA[x].seq
				seqB := dataB[y].seq
				paralellNeedlemanWunsch(seqA, seqB)

				endTime := time.Now()

				timeRow.strLen = (len(seqA) * len(seqB))
				timeRow.runTime = int(endTime.Sub(startTime))
				runTimes = append(runTimes, timeRow)
			}
		}
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
			arrMat[i][j] = arrMin(resMatch, resDel, resIns)

		}
	}
}

func paralellNeedlemanWunsch(seqA, seqB string) {

}

func arrMin(a, b, c int) int {
	diagonal := 1
	up := 2
	left := 3
	if a < b && b < c {
		return diagonal
	} else if a > b && b > c {
		return up
	} else {
		return left
	}
}

func min(a, b, c int) int {
	if a < b && b < c {
		return a
	} else if a > b && b > c {
		return c
	} else {
		return b
	}
}
