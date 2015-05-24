package gogo

import (
	"math"
)

/*
Round returns value rounded on the decimal value.
*/
func Round(value float64, decimal int) float64 {
	var rounding float64
	pow := math.Pow(10, float64(decimal))

	afterComma := value * pow
	_, fraction := math.Modf(afterComma)

	if fraction >= 0.5 {
		rounding = math.Ceil(afterComma)
	} else {
		rounding = math.Floor(afterComma)
	}

	return rounding / pow
}
