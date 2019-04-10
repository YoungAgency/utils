package utils

import "math"

func ReduceDecimalPrecision(v float64, p int) (r float64) {
	return math.Round(v*math.Pow10(p)) / math.Pow10(p)
}
