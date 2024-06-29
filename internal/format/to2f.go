package format

import "math"

func To2f(f float64) float64 {
	return float64(math.Round(f*100)) / 100
}
