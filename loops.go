package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	t, z := 0.0, 1.0
	for abs(t-z) > 1e-8 {
		t, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
