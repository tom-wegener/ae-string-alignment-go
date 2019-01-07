package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func parseFiles(path string) (data []record) {
	faFile, err := os.Open(path) //open file
	if err != nil {              //when errors occure....
		log.Fatal(err)
	}
	defer func() {
		if err = faFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(faFile)
	for scanner.Scan() {
		bLine := scanner.Text()
		sLine := string(bLine)
		buildFasta(sLine)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func buildFasta(line string) (data []record) {
	if line[:1] == ">" {
		var rec record
		splittedLine := strings.Split(line, " ")
		rec.key = splittedLine[0]
		rec.name = splittedLine[1]
		rec.seq = ""
		data = append(data, rec)
	} else {
		data[len(data)-1].seq = data[len(data)-1].seq + line
	}
	return data
}
