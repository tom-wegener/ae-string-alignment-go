package main

import (
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func plotIt(runTimesNoS, runTimesRaS, runTimesNoL, runTimesRaL []times) {

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "benötigte Zeit"
	p.X.Label.Text = "Datenlänge"
	p.Y.Label.Text = "benötigte Zeit in Sekunden"

	err = plotutil.AddLinePoints(p,
		"Fasta Short", plotTimes(runTimesNoS),
		"Zufall Short", plotTimes(runTimesRaS),
		"Fasta Long", plotTimes(runTimesNoL),
		"Zufall Long", plotTimes(runTimesRaL),
	)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func plotTimes(runTimes []times) plotter.XYs {
	//sortedTimes := sortTimes(runTimes)
	sort.Sort(runTimesArr(runTimes))
	sortedTimes := runTimes
	//sortedTimes := runTimes
	n := len(sortedTimes)
	pts := make(plotter.XYs, n)
	for i, row := range sortedTimes {
		pts[i].X = float64(row.strLen)
		pts[i].Y = float64(row.runTime)
	}
	return pts
}

/*func sortTimes(runTimes []times) (sortedTimes []times) {
	sortedTimes = runTimes
	sort.Slice(sortedTimes, func(i, j int) bool {
		return sortedTimes[i].len < sortedTimes[j].len
	})
	return sortedTimes
}*/
