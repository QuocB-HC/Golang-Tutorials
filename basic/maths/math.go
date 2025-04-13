package maths

import (
	"math"
	"math/rand"
)

func RandomNumber() int {
	return rand.Intn(100)
}

func Add(x, y int) int {
	return x + y
}

func SqrtNumber(x float64) float64 {
	return math.Sqrt(x)
}

func Pi() float64 {
	return math.Pi
}
