package gogo

import (
	"fmt"
	"math"
)

func Round(value float64, comma int) float64 {
	var rounding float64
	afterComma := value * math.Pow(10, float64(comma))

	if value >= 0.5 {
		rounding = math.Ceil(afterComma)
	} else {
		rounding = math.Floor(afterComma)
	}

	return rounding / math.Pow(10, float64(comma))
}
