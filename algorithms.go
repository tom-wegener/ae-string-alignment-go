package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func compareFiles(dataA []record, dataB []record, alg string) (runTimes runTimesArr) {
	runNum := 100
	runTimes = make(runTimesArr, 0)

	if alg == "nwa" {
		for x := 0; x < runNum; x++ {
			for y := 0; y < runNum; y++ {
				var timeRow times
				startTime := time.Now()

				seqA := dataA[x].seq
				seqB := dataB[y].seq
				nwa(seqA, seqB)

				endTime := time.Now()

				timeRow.strLen = (len(seqA) * len(seqB))
				timeRow.runTime = int(endTime.Sub(startTime))
				runTimes = append(runTimes, timeRow)
			}
		}
	} else if alg == "hnwa" {
		for x := 0; x < runNum; x++ {
			for y := 0; y < runNum; y++ {
				var timeRow times
				startTime := time.Now()

				seqA := dataA[x].seq
				seqB := dataB[y].seq
				hnwa(seqA, seqB)

				endTime := time.Now()

				timeRow.strLen = (len(seqA) * len(seqB))
				timeRow.runTime = int(endTime.Sub(startTime))
				runTimes = append(runTimes, timeRow)
			}
		}
	} else if alg == "snwa" {
		for x := 0; x < runNum; x++ {
			for y := 0; y < runNum; y++ {
				var timeRow times
				startTime := time.Now()

				seqA := dataA[x].seq
				seqB := dataB[y].seq
				hnwa(seqA, seqB)

				endTime := time.Now()

				timeRow.strLen = (len(seqA) * len(seqB))
				timeRow.runTime = int(endTime.Sub(startTime))
				runTimes = append(runTimes, timeRow)
			}
		}
	} else {
		fmt.Println("The algorithm in the config-file does not exist")
		fmt.Print("Please choose on of the following two: \n- nwa for Needleman-Wunsch-Algorithm \n- snwa for a splitted nwa where the value is not exact \n-hnwa for Hirschberg-Needleman-Wunsch \n")
		os.Exit(1)
	}

	return runTimes
}

func nwa(seqA, seqB string) {

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
			numMat[i][j], arrMat[i][j] = min(resMatch, resDel, resIns)

		}
	}
}

func hnwa(seqA, seqB string) {
	fmt.Println("not working")
	/*p := 1
	aux := seqA
	minLen := 10
	if (len(seqA) > minLen || len(seqB) > minLen) && p < P {

	} else {
		nwa(seqA, seqB)
	}
	*/
}

func snwa(seqA, seqB string) {

}

func splitString(a string) (string, string) {
	strLen := int(len(a) / 2)
	//make the string into a Sclice, bisect it and return it
	strSli := strings.Split(a, "")
	firHalf := strings.Join(strSli[:strLen], "")
	secHalf := strings.Join(strSli[strLen:], "")
	return firHalf, secHalf

}

func min(a, b, c int) (int, int) {
	diagonal := 1
	up := 2
	left := 3
	if a < b && b < c {
		return a, diagonal
	} else if a > b && b > c {
		return c, up
	} else {
		return b, left
	}
}
