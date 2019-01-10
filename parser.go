package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseFiles(path string) (data []record) {
	//data := make([]record, 0)
	var rec record

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
		if sLine[:1] == ">" {
			splittedLine := strings.Split(sLine, "|")
			rec.key = splittedLine[0]
			rec.name = splittedLine[1]
			rec.seq = ""

			data = append(data, rec)
		} else {
			rec.seq = rec.seq + sLine
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("no error occured in Parsing")

	return data
}
