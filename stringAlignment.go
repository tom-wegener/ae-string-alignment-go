package main

func main() {
	aeae := "data/aedes_aegypti_protein.fa"
	aeal := "data/aedes_albopictus_protein.fa"
	dataAEAE := parseFiles(aeae)
	dataAEAL := parseFiles(aeal)
	compareFiles(dataAEAE, dataAEAL)
}
