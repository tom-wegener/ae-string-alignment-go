package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

func parseFiles(path string) (data []record) {
	var rec record

	faFile, err := os.Open(path) //open file
	check(err)

	scanner := bufio.NewScanner(faFile)
	for scanner.Scan() {
		bLine := scanner.Text()
		sLine := string(bLine)
		if sLine[:1] == ">" {
			splittedLine := strings.Split(sLine, "|")
			rec.key = splittedLine[0]
			for _, part := range splittedLine {
				if part[:2] == "XP" {
					rec.name = part
					break
				}
			}
			rec.name = splittedLine[1]
			rec.seq = ""

			data = append(data, rec)
		} else {
			data[len(data)-1].seq = data[len(data)-1].seq + sLine
		}
	}

	err = scanner.Err()
	check(err)
	faFile.Close()
	check(err)
	return data
}

func genEntr() (data []record) {
	for i := 0; i < 101; i++ {
		var rec record
		rec.key = ">" + randomString(9)
		rec.name = randomString(10)
		strLen := 300
		if i > 50 {
			strLen = i * 50
		} else {
			strLen = i * 100
		}
		rec.seq = randomString(strLen)

		data = append(data, rec)
	}
	return data
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
