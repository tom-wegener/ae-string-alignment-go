package main

import (
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()

	aeae := "data/aedes_aegypti_protein.fa"
	aeal := "data/aedes_albopictus_protein.fa"

	var dataAEAE, dataAEAL, dataRA, dataRB []record

	dataAEAE = parseFiles(aeae)
	dataAEAL = parseFiles(aeal)

	dataRA = genEntr()
	dataRB = genEntr()

	runTimesNoS := compareFiles(dataAEAE, dataAEAL, false)
	runTimesRaS := compareFiles(dataRA, dataRB, false)

	runTimesNoL := compareFiles(dataAEAE, dataAEAL, true)
	runTimesRaL := compareFiles(dataRA, dataRB, true)

	plotIt(runTimesNoS, runTimesRaS, runTimesNoL, runTimesRaL)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
