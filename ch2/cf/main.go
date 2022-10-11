// Cf converts its numeric argument to Celsius and Fahrenheit
package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/benjamindburke/go-studies/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n%s = %s, %s = %s\n", f, tempconv.FToC(f), f, tempconv.FToK(f), c, tempconv.CToF(c), c, tempconv.CToK(c), k, tempconv.KToF(k), k, tempconv.KToC(k))
	}
}