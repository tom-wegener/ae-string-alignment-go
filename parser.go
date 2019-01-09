package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var rec record

func buildFasta(line string, data *[]record) {
	if line[:1] == ">" {
		if len(rec.key) > 0 {
			*data = append(*data, rec)
		}
		splittedLine := strings.Split(line, " ")
		rec.key = splittedLine[0]
		rec.name = splittedLine[1]
		rec.seq = ""
	} else {
		rec.seq = rec.seq + line
	}

	//return data
}

func parseFiles(path string) (data []record) {
	//data := make([]record, 0)
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
		buildFasta(sLine, &data)
		fmt.Print(data)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return data
}
