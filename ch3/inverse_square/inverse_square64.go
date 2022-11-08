package main

import (
	"fmt"
	"math"
)

var threehalfs float64 = 1.5

// a "magic" constant used to approximate log2(x) across interval [0, 1] with the overall smallest margin of error
// this constant can be improved upon by deriving the inverse square root function, which we will not do here
const magic64 = 0x5fe6eb50c7b537a9

func newtonian_derivative(x2 float64, y float64) float64 {
	f := y * (threehalfs - (x2 * y * y))
	return f
}

func q_sqrt(number float64) float64 {
	x2 := number * 0.5

	// convert IEEE 754 floating point num to uint preserving bits, just intepret Mantissa + Exponent form as uint number
	b := math.Float64bits(number)
	// bitwise - bit shifting left one bit divides by two, circumenving need for slow / operator
	b = magic64 - (b >> 1)
	// convert uint to IEEE 754 floating point num preserving bits, just intepret uint number as Mantissa + Exponent form
	y := math.Float64frombits(b)

	// newtonian calculus - derivative approximation function of x - ( f(x) / f'(x) )
	y = newtonian_derivative(x2, y)
	y = newtonian_derivative(x2, y) // approximate a second time if margin of error is not low enough

	return y
}

func error_margin(known float64, calculated float64) float64 {
	e := math.Abs(float64((calculated/known)-1.0)) * 100.0
	return e
}

func main() {
	var num, known_inverse_sqrt float64 = 4.0, 0.5
	var inverse_sqrt float64 = q_sqrt(num)
	var err float64 = error_margin(known_inverse_sqrt, inverse_sqrt)

	fmt.Printf("num=%.2f, inverse_square=%.15f\n", num, inverse_sqrt)
	fmt.Printf("known_inverse_square=%.2f, error_margin=%.15f%%\n", known_inverse_sqrt, err)
}
