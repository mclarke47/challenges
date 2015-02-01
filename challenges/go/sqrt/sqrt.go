package main

import (
	"fmt"
	"math"
)

// z = z - ((z^2-x)/(2z))

func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 20; i++ {

		z = z - (((z*z)-x)/(2*z))
	}

	return z
}

func formatMessage(x float64) string {

	return fmt.Sprintf("sqrt of %f  is approximatly %f and exactly %f\n", x, Sqrt(x), math.Sqrt(x))
}

func main() {
	fmt.Printf(formatMessage(1.0))
	fmt.Printf(formatMessage(2.0))
	fmt.Printf(formatMessage(3.0))
	fmt.Printf(formatMessage(4.0))
	fmt.Printf(formatMessage(5.0))
	fmt.Printf(formatMessage(6.0))
	fmt.Printf(formatMessage(7.0))
	fmt.Printf(formatMessage(8.0))
	fmt.Printf(formatMessage(9.0))
	fmt.Printf(formatMessage(654558.0))
}
