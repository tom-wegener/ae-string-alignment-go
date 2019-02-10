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

	//sort and print Times
	//pts := []
	ptsNoS := pnpTimes(runTimesNoS, "NormalShort")
	ptsRaS := pnpTimes(runTimesRaS, "RandomShort")
	ptsNoL := pnpTimes(runTimesNoL, "NormalLangsam")
	ptsRaL := pnpTimes(runTimesRaL, "RandomLangsam")

	err = plotutil.AddScatters(p,
		"Fasta Short", ptsNoS,
		"Zufall Short", ptsRaS,
		"Fasta Long", ptsNoL,
		"Zufall Long", ptsRaL,
	)

	check(err)

	// Save the plot to a PNG file.
	err = p.Save(10*vg.Inch, 10*vg.Inch, "points.png")
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
	//printing
	f, err := os.Create(idStr)
	check(err)
	for i := range pts {
		fmt.Fprintf(f, "%v , %v \n", pts[i].X, pts[i].Y)
	}
	f.Close()
	return pts
}
