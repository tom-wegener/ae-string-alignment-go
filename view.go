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

func plotIt(runTimesNoS, runTimesRaS, runTimesNoL, runTimesRaL []times) {

	fmt.Println("starting to plot")
	p, err := plot.New()
	check(err)

	p.Title.Text = "benötigte Zeit"
	p.X.Label.Text = "Datenlänge"
	p.Y.Label.Text = "benötigte Zeit in Sekunden"

	err = plotutil.AddLinePoints(p,
		"Fasta Short", plotTimes(runTimesNoS, "NormalShort"),
		"Zufall Short", plotTimes(runTimesRaS, "RandomShort"),
		"Fasta Long", plotTimes(runTimesNoL, "NormalLangsam"),
		"Zufall Long", plotTimes(runTimesRaL, "RandomLangsam"),
	)
	check(err)

	// Save the plot to a PNG file.
	err = p.Save(10*vg.Inch, 10*vg.Inch, "points.png")
	check(err)
}

func plotTimes(runTimes []times, idStr string) plotter.XYs {
	sort.Sort(runTimesArr(runTimes))
	sortedTimes := runTimes
	n := len(sortedTimes)
	pts := make(plotter.XYs, n)
	for i, row := range sortedTimes {
		pts[i].X = float64(row.strLen)
		pts[i].Y = float64(row.runTime)
	}
	printTimes(pts, idStr)
	return pts
}

func printTimes(pts plotter.XYs, idStr string) {
	fmt.Println("algorithm is done")
	f, err := os.Create(idStr)
	check(err)
	for i := range pts {
		fmt.Fprintf(f, "%v , %v \n", pts[i].X, pts[i].Y)
	}
	f.Close()
}
