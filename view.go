package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func plotIt(runTimesA, runTimesB [][]int) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "benötigte Zeit"
	p.X.Label.Text = "Anzahl der Vergleiche pro Durchgang"
	p.Y.Label.Text = "benötigte Zeit in Nanosekunden"

	err = plotutil.AddLinePoints(p,
		"Fasta", plotTimes(runTimesA),
		"Zufall", plotTimes(runTimesB))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func getTimes(runTimes [][]int) (column []int) {
	column = make([]int, 0)
	for _, row := range runTimes {
		column = append(column, row[1])
	}
	return
}

func plotTimes(runTimes [][]int) plotter.XYs {
	n := len(runTimes)
	pts := make(plotter.XYs, n)
	for i, row := range runTimes {
		pts[i].X = float64(row[0])
		pts[i].Y = float64(row[1])
	}
	return pts
}
