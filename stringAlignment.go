package main

import (
	"io/ioutil"

	"github.com/olebedev/config"
	"github.com/pkg/profile"
)

var cfg *config.Config

func main() {
	cfg = readConf()
	profileBool, err := cfg.Bool("profile")
	check(err)
	if profileBool == true {
		defer profile.Start().Stop()
	}

	test, err := cfg.Bool("test")
	check(err)
	if test == true {
		testAlgorithms()
	} else {
		profileAlgorithms()
	}

}

func testAlgorithms() {
	firstFile, err := cfg.String("files.first")
	check(err)
	secFile, err := cfg.String("files.second")
	check(err)

	runTimesComp := make([]runTimesArr, 2)

	dataA := parseFiles(firstFile)
	dataB := parseFiles(secFile)

	alg := "snwa"
	runTimesComp[0] = compareFilesTest(dataA, dataB, alg)

	alg = "nwa"
	runTimesComp[1] = compareFilesTest(dataA, dataB, alg)

	runNamesComp := make([]string, 2)
	runNamesComp[0] = "snwa"
	runNamesComp[1] = "nwa"

	saveTimes(runTimesComp[0], runNamesComp[0])
	saveTimes(runTimesComp[1], runNamesComp[1])
}

func profileAlgorithms() {
	firstFile, err := cfg.String("files.first")
	check(err)
	secFile, err := cfg.String("files.second")
	check(err)

	alg, err := cfg.String("algorithm")
	check(err)

	dataA := parseFiles(firstFile)
	dataB := parseFiles(secFile)

	dataRA := genEntr()
	dataRB := genEntr()

	//var runTimesComp []runTimesArr
	runTimesComp := make([]runTimesArr, 2)
	runTimesComp[0] = compareFiles(dataRA, dataRB, alg)
	runTimesComp[1] = compareFiles(dataA, dataB, alg)

	//var runNamesComp []string
	runNamesComp := make([]string, 2)
	runNamesComp[0] = "fasta"
	runNamesComp[1] = "Random"

	plotIt(runTimesComp, runNamesComp)
}

func readConf() *config.Config {
	var confFile, err = ioutil.ReadFile("config.yml")
	check(err)
	yamlString := string(confFile)

	cfg, err := config.ParseYaml(yamlString)
	check(err)
	return cfg
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
