package main

import (
	"gonum.org/v1/plot/plotter"
)

func main() {
	points := plotter.XYs{}
	for i := 0; i <= 500; i++ {
		t := float64(i) / 100.0
		points = append(points, plotter.XY{
			X: t,
			Y: exponential(1, 2, 1, t),
		})
	}
	b := bounds{
		xmin: 0,
		xmax: 5,
		ymin: 0,
		ymax: 20,
	}
	l := labels{
		x: "t",
		y: "growth",
	}
	CreateLineplotPlot(points, "t - growth", l, b, "exponential.png")
	var points2 plotter.XYs
	for i := 0; i <= 4000; i++ {
		t := float64(i) / 100.0
		points2 = append(points2, plotter.XY{
			X: t,
			Y: logistic(1, 1, 1000, t),
		})
	}
	CreateLineplotPlot(points2, "t - growth", l, b, "logistic.png")
}
