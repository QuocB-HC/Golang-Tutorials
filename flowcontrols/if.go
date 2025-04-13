package flowcontrols

import (
	"fmt"
	"math"
)

func SqrtByIf(x float64) string {
	if x < 0 {
		return SqrtByIf(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func PowByIf(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
