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
	runTimesComp[0] = compareFiles(dataA, dataB, alg)
	runTimesComp[1] = compareFiles(dataRA, dataRB, alg)

	//var runNamesComp []string
	runNamesComp := make([]string, 2)
	runNamesComp[0] = firstFile + secFile
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
