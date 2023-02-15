package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	fahrenheit float64
	celsius    float64
	kelvin     float64
	out        string
)

func init() {
	flag.Float64Var(&fahrenheit, "F", 0.0, "temprature in fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temprature in celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temprature in kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - fahrenheit, K - kelvin")

}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func KelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func FahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9 + 273.15
}

func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func main() {
	flag.Parse()

	if out == "C" && isFlagPassed("F") {
		fmt.Println(fahrenheit, "°F is equal to °C", math.Round(FahrenheitToCelsius(fahrenheit)*100)/100)
	}

	if out == "F" && isFlagPassed("C") {
		fmt.Println(celsius, "°C is equal to °F", math.Round(CelsiusToFahrenheit(celsius)*100)/100)
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Println(celsius, "°C is equal to °K", math.Round(CelsiusToKelvin(celsius)*100)/100)
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Println(kelvin, "K is equal to °C", math.Round(KelvinToCelsius(kelvin)*100)/100)
	}

	if out == "F" && isFlagPassed("K") {
		fmt.Println(kelvin, "K is equal to °F", math.Round(KelvinToFahrenheit(kelvin)*100)/100)
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
