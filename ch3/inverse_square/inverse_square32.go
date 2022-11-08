package main

import (
	"fmt"
	"math"
)

var threehalfs float32 = 1.5

// a "magic" constant used to approximate log2(x) across interval [0, 1] with the overall smallest margin of error
// this constant can be improved upon by deriving the inverse square root function, which we will not do here
const magic32 = 0x5f375a86

func newtonian_derivative(x2 float32, y float32) float32 {
	f := y * (threehalfs - (x2 * y * y))
	return f
}

func q_sqrt(number float32) float32 {
	x2 := number * 0.5

	// convert IEEE 754 floating point num to uint preserving bits, just intepret Mantissa + Exponent form as uint number
	b := math.Float32bits(number)
	// bitwise - bit shifting left one bit divides by two, circumenving need for slow / operator
	b = magic32 - (b >> 1)
	// convert uint to IEEE 754 floating point num preserving bits, just intepret uint number as Mantissa + Exponent form
	y := math.Float32frombits(b)

	// newtonian calculus - derivative approximation function of x - ( f(x) / f'(x) )
	y = newtonian_derivative(x2, y)
	y = newtonian_derivative(x2, y) // approximate a second time if margin of error is not low enough

	return y
}

func error_margin(known float32, calculated float32) float32 {
	e := float32(math.Abs(float64((calculated/known)-1.0)) * 100.0)
	return e
}

func main() {
	var num, known_inverse_sqrt float32 = 4.0, 0.5
	var inverse_sqrt float32 = q_sqrt(num)
	var err float32 = error_margin(known_inverse_sqrt, inverse_sqrt)

	fmt.Printf("num=%.2f, inverse_square=%.15f\n", num, inverse_sqrt)
	fmt.Printf("known_inverse_square=%.2f, error_margin=%.15f%%\n", known_inverse_sqrt, err)
}
