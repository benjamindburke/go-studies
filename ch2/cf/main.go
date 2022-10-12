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
		fmt.Printf("%s\t= %s\n\t= %s\n\n", f, tempconv.FToC(f), tempconv.FToK(f))
		fmt.Printf("%s\t= %s\n\t= %s\n\n", c, tempconv.CToF(c), tempconv.CToK(c))
		fmt.Printf("%s\t= %s\n\t= %s\n\n", k, tempconv.KToF(k), tempconv.KToC(k))
	}
}