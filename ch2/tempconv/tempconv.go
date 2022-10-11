// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Kelvin float64
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroK	Kelvin = 0
	FreezingK			Kelvin = 273.15
	BoilingK			Kelvin = 373.15
)

const (
	AbsoluteZeroC Celsius = KToC(AbsoluteZeroK)
	FreezingC			Celsius = KToC(FreezingK)
	BoilingC			Celsius = KToC(BoilingK)
)

func (k Kelvin) String() string 		{ return fmt.Sprintf("%g°K") }
func (c Celsius) String() string		{ return fmt.Sprintf("%g°C") }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F") }