package main

import "math"

func exponential(c, b, d, t float64) float64 {
	return c * math.Exp((b-d)*t)
}
