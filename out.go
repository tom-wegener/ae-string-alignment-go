package main

import (
	"fmt"
	"os"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func plotIt(runTimesComp []runTimesArr, runNamesComp []string) {

	p, err := plot.New()
	check(err)

	p.Title.Text = "benötigte Zeit"
	p.X.Label.Text = "Datenlänge"
	p.Y.Label.Text = "benötigte Zeit in Nanosekunden"

	plotPoints := make([]plotter.XYs, 2)

	for i, arr := range runTimesComp {
		name := runNamesComp[i]
		plotPoints[i] = pnpTimes(arr, name)
		saveTimes(arr, name)
	}

	var args = make([]interface{}, 0, 2*len(plotPoints))

	for i, v := range plotPoints {
		args = append(args, runNamesComp[i], v)
	}
	err = plotutil.AddScatters(p, args...)
	check(err)
	// Save the plot to a PNG file.

	outputFile, err := cfg.String("files.scatter")
	check(err)
	err = p.Save(10*vg.Inch, 10*vg.Inch, outputFile)
	check(err)
}

func pnpTimes(runTimes []times, idStr string) plotter.XYs {
	//sorting
	sort.Sort(runTimesArr(runTimes))
	sortedTimes := runTimes

	//making plot
	n := len(sortedTimes)
	pts := make(plotter.XYs, n)
	for i, row := range sortedTimes {
		pts[i].X = float64(row.strLen)
		pts[i].Y = float64(row.runTime)
	}
	return pts
}

func saveScore(scoresColl scoresArr) {
	outputFile, err := cfg.String("files.scores")
	f, err := os.Create(outputFile)
	check(err)
	for i := range scoresColl {
		fmt.Fprintf(f, "%v , %v \n", scoresColl[i].key, scoresColl[i].score)
	}
	f.Close()

}

func saveTimes(timesColl runTimesArr, outputFile string) {
	//outputFile, err := cfg.String("files.times")
	f, err := os.Create(outputFile)
	check(err)
	for i := range timesColl {
		fmt.Fprintf(f, "%v , %v \n", timesColl[i].strLen, timesColl[i].runTime)
	}
	f.Close()
	check(err)

}
