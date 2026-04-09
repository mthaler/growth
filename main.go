package main

import (
	"fmt"

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
		deltat := 1.0 / 100.0
		if t == 0 {
			points2 = append(points2, plotter.XY{
				X: t,
				Y: 0,
			})
		} else {
			points2 = append(points2, plotter.XY{
				X: t,
				Y: logistic(points2[i-1].Y, 1, 1, deltat),
			})
		}
	}
	fmt.Println(logistic(0, 1, 1, 1/100.0))
	CreateLineplotPlot(points2, "t - growth", l, b, "logistic.png")
}
