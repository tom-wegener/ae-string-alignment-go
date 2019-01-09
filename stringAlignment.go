package main

func main() {
	aeae := "/home/luca/Documents/data/aedes_aegypti_protein.fa"
	aeal := "/home/luca/Documents/data/aedes_albopictus_protein.fa"
	dataAEAE := parseFiles(aeae)
	dataAEAL := parseFiles(aeal)
	compareFiles(dataAEAE, dataAEAL)
}
