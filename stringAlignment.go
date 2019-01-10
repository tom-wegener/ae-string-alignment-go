package main

import "fmt"

func main() {
	aeae := "data/aedes_aegypti_protein.fa"
	aeal := "data/aedes_albopictus_protein.fa"
	dataAEAE := parseFiles(aeae)
	dataAEAL := parseFiles(aeal)
	dataRA := genEntr()
	dataRB := genEntr()
	fmt.Println(dataRA[1].seq)
	runTimesA := compareFiles(dataAEAE, dataAEAL)
	runTimesB := compareFiles(dataRA, dataRB)
	plotIt(runTimesA, runTimesB)
}
