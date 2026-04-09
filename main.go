package main

import "gonum.org/v1/plot/plotter"

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
		ymin: -0.5,
		ymax: 0.5,
	}
}
