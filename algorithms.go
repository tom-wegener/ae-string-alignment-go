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
	scoresColl := make(scoresArr, 0)

	if alg == "nwa" {
		for x := 0; x < runNum; x++ {
			for y := 0; y < runNum; y++ {
				var timeRow times
				var scoreRow scores

				startTime := time.Now()

				seqA := dataA[x].seq
				seqB := dataB[y].seq
				scoreRow.score = nwa(seqA, seqB)
				keyA := dataA[x].key
				keyB := dataB[y].key
				scoreRow.key = keyA + "," + keyB
				scoresColl = append(scoresColl, scoreRow)

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
				var scoreRow scores
				startTime := time.Now()

				seqA := dataA[x].seq
				seqB := dataB[y].seq

				scoreRow.score = hnwa(seqA, seqB)
				keyA := dataA[x].key
				keyB := dataB[y].key
				scoreRow.key = keyA + "," + keyB
				scoresColl = append(scoresColl, scoreRow)

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
				var scoreRow scores
				startTime := time.Now()

				seqA := dataA[x].seq
				seqB := dataB[y].seq

				scoreRow.score = snwa(seqA, seqB)
				keyA := dataA[x].name
				keyB := dataB[y].name
				scoreRow.key = keyA + "," + keyB
				scoresColl = append(scoresColl, scoreRow)

				endTime := time.Now()

				timeRow.strLen = (len(seqA) * len(seqB))
				timeRow.runTime = int(endTime.Sub(startTime))
				runTimes = append(runTimes, timeRow)
			}
		}
	} else {
		fmt.Println("The algorithm you configured in the config-file does not exist")
		fmt.Print("Please choose on of the following three: \n- nwa for Needleman-Wunsch-Algorithm \n- snwa for a splitted nwa where the value is not exact \n-hnwa for Hirschberg-Needleman-Wunsch (not done yet)\n")
		os.Exit(1)
	}

	saveScore(scoresColl)
	return runTimes
}

func nwa(seqA, seqB string) int {

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
			numMat[i][j], arrMat[i][j] = max(resMatch, resDel, resIns)

		}
	}
	return numMat[a-1][b-1]
}

func hnwa(seqA, seqB string) int {
	fmt.Println("not working at the moment, sorry")
	/*p := 1
	aux := seqA
	minLen := 10
	if (len(seqA) > minLen || len(seqB) > minLen) && p < P {

	} else {
		nwa(seqA, seqB)
	}
	*/
	os.Exit(1)
	return 1
}

func snwa(seqA, seqB string) int {
	//calculate a point where a letter is the same in both strings nearest to the middle
	splitterA, splitterB := findSame(seqA, seqB)

	var score int
	if splitterA == 0 || splitterB == 0 {
		score = nwa(seqA, seqB)
	} else {
		seqAA, seqAB := bisectString(seqA, splitterA)
		seqBA, seqBB := bisectString(seqB, splitterB)
		scoreA := nwa(seqAA, seqBA)
		scoreB := nwa(seqAB, seqBB)
		score = scoreA + scoreB
	}

	return score
}

func findSame(seqA, seqB string) (int, int) {
	for i := 0; i < len(seqA); i++ {
		for j := 0; j < len(seqB) && j < (i+5); j++ {
			k := snwaHelper(i, len(seqA))
			l := snwaHelper(j, len(seqB))
			if seqA[k] == seqB[l] {
				return k, l
			}
		}
	}
	return 0, 0
}

func snwaHelper(a, b int) int {
	c := a
	if a%2 == 0 {
		c = a / 2
	} else {
		c = int(-(a/2 + 1))
	}
	d := b/2 + c
	return d
}

func bisectString(a string, strLen int) (string, string) {
	//make the string into a Sclice, bisect it and return it
	strSli := strings.Split(a, "")
	firHalf := strings.Join(strSli[:strLen], "")
	secHalf := strings.Join(strSli[strLen:], "")
	return firHalf, secHalf

}

func max(a, b, c int) (int, int) {
	diagonal := 1
	up := 2
	left := 3
	if c < b && b < a {
		return a, diagonal
	} else if a < b && b < c {
		return c, up
	} else {
		return b, left
	}
}
