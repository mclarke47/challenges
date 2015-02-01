package main

import (
	"fmt"
	"math"
)

// z = z - ((z^2-x)/(2z))

func Sqrt(number float32, next float32, prev float32) float32 {

	if next == prev {
		return next
	}

	z:= (next - (((next*next)-number)/(2*next)))

	return Sqrt(number, z, next)
}
func formatMessage(x float32) string {

	return fmt.Sprintf("sqrt of %f  is approximatly %f and exactly %f\n", x, Sqrt(x, 1.0 , 0.0), math.Sqrt(float64(x)))
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