package main

import (
	"flag"
	"fmt"

	"github.com/Johannekh/funtemps/conv"
)


var fahr float64
var celc float64
var kelv float64
var out string
var funfacts string


func init() {



	
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celc, "C", 0.0, "temperatur i grader celcius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	

}

func main() {

	flag.Parse()

	

	
	fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	if out == "K" && isFlagPassed("C") {
		Kevin := conv.CelsiusToKelvin(fahr)
		if fahr == float64(int(fahr)) {
			fmt.Printf("%.0f degrees celsius is %.2f degrees kevin.\n", fahr, Kevin)
		} else {
			fmt.Printf("%.3f degrees celsius is %.2f degrees kevin.\n", fahr, Kevin)
		}
	}

	if out == "K" && isFlagPassed("C") {
		Kevin := conv.CelsiusToKelvin(fahr)
		if fahr == float64(int(fahr)) {
			fmt.Printf("%.0f degrees celsius is %.2f degrees kevin.\n", fahr, Kevin)
		} else {
			fmt.Printf("%.3f degrees celsius is %.2f degrees kevin.\n", fahr, Kevin)
		}
	}

	if out == "C" && isFlagPassed("K") {
		Celsius := conv.KelvinToCelsius(fahr)
		if fahr == float64(int(fahr)) {
			fmt.Printf("%.0f degrees kevin is %.2f degrees celsius.\n", fahr, Celsius)
		} else {
			fmt.Printf("%.3f degrees kevin is %.2f degrees celsius.\n", fahr, Celsius)
		}
	}
	if out == "F" && isFlagPassed("C") {
		Farhrenheit := conv.CelsiusToFarhrenheit(fahr)
		if fahr == float64(int(fahr)) {
			fmt.Printf("%.0f degrees celsius is %.2f degrees Farhrenheit.\n", fahr, Farhrenheit)
		} else {
			fmt.Printf("%.3f degrees celsius is %.2f degrees Farhrenheit.\n", fahr, Farhrenheit)
		}
	}

	if out == "F" && isFlagPassed("K") {
		Farhrenheit := conv.KelvinToFarhrenheit(fahr)
		if fahr == float64(int(fahr)) {
			fmt.Printf("%.0f degrees kevin is %.2f degrees Farhrenheit.\n", fahr, Farhrenheit)
		} else {
			fmt.Printf("%.3f degrees kevin is %.2f degrees Farhrenheit.\n", fahr, Farhrenheit)
		}
	}
	if out == "K" && isFlagPassed("F") {
		Kevin := conv.FarhenheitToKelvin(fahr)
		if fahr == float64(int(fahr)) {
			fmt.Printf("%.0f degrees Farhrenheit is %.2f degrees Kevin.\n", fahr, Kevin)
		} else {
			fmt.Printf("%.3f degrees Farhrenheit is %.2f degrees Kevin.\n", fahr, Kevin)
		}

	}

	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		Celsius := conv.FarhenheitToCelsius(fahr)
		if fahr == float64(int(fahr)) {
			fmt.Printf("%.0f degrees Fahrenheit is %.2f degrees celsius\n", fahr, Celsius)
		} else {
			fmt.Printf("%.3f degrees Fahrenheit is %.2f degrees celsius\n", fahr, Celsius)
		}
	}

}


func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
