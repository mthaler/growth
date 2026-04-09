package main

import "math"

func exponential(c, b, d, t float64) float64 {
	return c * math.Exp((b-d)*t)
}

func logistic(nprev, r, k, deltat float64) float64 {
	n := nprev + deltat*r*nprev*(1-nprev/k)
	return n
}
