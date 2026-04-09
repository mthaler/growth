package main

import (
	"math"
)

func exponential(c, b, d, t float64) float64 {
	return c * math.Exp((b-d)*t)
}

func logistic(r, k, n0, t float64) float64 {
	n := k / (1 + (k-n0)/n0*math.Exp(-r*t))
	return n
}
