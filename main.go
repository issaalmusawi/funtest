package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/issaalmusawi/funtest/funfacts"
	conv "github.com/issaalmusawi/funtest/konv"
)

// definerer hoved variabler
var (
	fahrenheit      float64
	celsius         float64
	kelvin          float64
	out             string
	temperatureType string
	Facts           string
)

// initialiserer variabel-flagg, der parameterne "pointer" til hovedvariablene over,
// der første string referer til flagg, default verdien "0.0", og tilslutt info/hjelp til forbrukeren
func init() {
	flag.Float64Var(&fahrenheit, "F", 0.0, "temprature in fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temprature in celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temprature in kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - fahrenheit, K - kelvin")
	flag.StringVar(&Facts, "funfacts", "Sun", "\"fun-facts\" om sun - Solen, luna - Månene og terra - Jorden")
	flag.StringVar(&temperatureType, "t", "F", "temperatur i F- fahrenheit, K - kelvin, C - celsius")
}

// definerer hva diverse flagg skal utføre av funksjoner. Hvis "C" ut og "F" inn, gjør dette:...
func main() {
	flag.Parse()

	if out == "C" && isFlagPassed("F") {
		fmt.Println(fahrenheit, "°F is equal to", math.Round(conv.FahrenheitToCelsius(fahrenheit)*100)/100, "°C")
	}

	if out == "F" && isFlagPassed("C") {
		fmt.Println(celsius, "°C is equal to", math.Round(conv.CelsiusToFahrenheit(celsius)*100)/100, "°F")
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Println(celsius, "°C is equal to", math.Round(conv.CelsiusToKelvin(celsius)*100)/100, "°K")
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Println(kelvin, "K is equal to", math.Round(conv.KelvinToCelsius(kelvin)*100)/100, "°C")
	}

	if out == "F" && isFlagPassed("K") {
		fmt.Println(kelvin, "K is equal to", math.Round(conv.KelvinToFahrenheit(kelvin)*100)/100, "°F")
	}
	/*if Facts == "Sun" && isFlagPassed("funfacts") {

		fmt.Println(funfacts.GetFunFacts(Facts))
	}

	if Facts == "Luna" && isFlagPassed("funfacts") {
		fmt.Println(funfacts.GetFunFacts(Facts))
	}
	if Facts == "Terra" && isFlagPassed("funfacts") {
		fmt.Println(funfacts.GetFunFacts((Facts)))
	}*/

	if isFlagPassed("t") && isFlagPassed("funfacts") {
		for _, Fact := range funfacts.GetFunFacts(Facts) {
			if temperatureType == "C" && Fact.Type == "K" {
				fmt.Println(Fact.Text, Fact.Value, Fact.Type)
			}
			if temperatureType == "C" && Fact.Type == "C" {
				fmt.Println(Fact.Text, Fact.Value, Fact.Type)
			}
		}

	}

	/*

		/*if isFlagPassed("funfacts") && out == ("t") {
			for _, Fact := range funfacts.GetFunFacts(Facts) {
				if temperatureType == "C" && Fact.Type == "C" {
					fmt.Println(Fact.Text, Fact.Value, Fact.Type)
				}
				//	{
				//	if fFacts.Type == "K" {
				//	fmt.Println(fFacts.Fact, fFacts.Value, fFacts.Type)
			}
		}

		//{Fact: "temperatur i solens kjerne", Value: 15000000, Type: "C"},
		//	{Fact: "temperatur på ytre lag av solen", Value: 5778, Type: "K"},
		//},

	}*/

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
