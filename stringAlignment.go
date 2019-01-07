package main

func main() {
	aeae := "/home/luca/Dosuments/data/aedes_aegypti_protein.fa"
	aeal := "/home/luca/Dosuments/data/aedes_albopictus_protein.fa"
	dataAEAE := parseFiles(aeae)
	dataAEAL := parseFiles(aeal)
	compareFiles(dataAEAE, dataAEAL)
}
