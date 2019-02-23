#! /usr/bin/env python3

import subprocess
import time
import csv


def printTimes(runTimes):
  with open("runOften.csv", 'w') as myfile:
    wr = csv.writer(myfile, quoting=csv.QUOTE_ALL)
    wr.writerow(runTimes)



def main():
  runTimes = []
  for i in range(0, 10):
    startTime = time.time()
    process = subprocess.Popen("./ae-string-alignment-go")
    process.wait()
    endTime = time.time()
    runTime = endTime - startTime
    runTimes.append([runTime])

  printTimes(runTimes)

main()