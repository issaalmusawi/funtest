package main

import (
	"flag"
	"fmt"
)

var (
	fahrenheitFlag = flag.Float64("f", 0.0, "temperature in Fahrenheit")
	celsiusFlag    = flag.Float64("c", 0.0, "temperature in Celsius")
	kelvinFlag     = flag.Float64("k", 0.0, "temperature in Kelvin")
)

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

	if *fahrenheitFlag != 0.0 {
		celsius := FahrenheitToCelsius(*fahrenheitFlag)
		fmt.Printf("%.2f°F is equal to %.2f°C\n", *fahrenheitFlag, celsius)
	}

	if *celsiusFlag != 0.0 {
		fahrenheit := CelsiusToFahrenheit(*celsiusFlag)
		kelvin := CelsiusToKelvin(*celsiusFlag)
		fmt.Printf("%.2f°C is equal to %.2f°F and %.2fK\n", *celsiusFlag, fahrenheit, kelvin)
	}

	if *kelvinFlag != 0.0 {
		fahrenheit := KelvinToFahrenheit(*kelvinFlag)
		celsius := KelvinToCelsius(*kelvinFlag)
		fmt.Printf("%.2fK is equal to %.2f°F and %.2f°C\n", *kelvinFlag, fahrenheit, celsius)
	}
}
