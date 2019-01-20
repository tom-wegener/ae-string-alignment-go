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

	runTimesA := compareFiles(dataAEAE, dataAEAL)
	runTimesB := compareFiles(dataRA, dataRB)

	plotIt(runTimesA, runTimesB)
}
