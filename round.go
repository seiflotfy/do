package gogo

import (
	"math"
)

/*
Round ...
Rounds a float64 number up to a specific decimal point.
*/
func Round(value float64, comma int) float64 {
	var rounding float64
	pow := math.Pow(10, float64(comma))

	afterComma := value * pow
	_, fraction := math.Modf(afterComma)

	if fraction >= 0.5 {
		rounding = math.Ceil(afterComma)
	} else {
		rounding = math.Floor(afterComma)
	}

	return rounding / pow
}
