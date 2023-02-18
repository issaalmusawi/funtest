package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	type test struct {
		input    string
		expected []Facts
	}
	tests := []test{
		{input: "Sun", expected: []Facts{
			{Interesting: "temperatur i solens kjerne", Value: 15000000, temperatureType: "C"},
			{Interesting: "temperatur på ytre lag av solen", Value: 5778, temperatureType: "K"}}},

		{input: "Terra", expected: []Facts{
			{Interesting: "høyeste temperatur målt på jordens overflate", Value: 56.7, temperatureType: "C"},
			{Interesting: "laveste temperatur målt på jordens overflate", Value: -69.7, temperatureType: "C"},
			{Interesting: "temperatur i jordens indre kjerne", Value: 9118.85, temperatureType: "C"}}},

		{input: "Luna", expected: []Facts{
			{Interesting: "temperatur på månens overflate om natten", Value: -183, temperatureType: "C"},
			{Interesting: "temperatur på månens overflate om dagen", Value: 106, temperatureType: "C"}}},
	}

	for _, tc := range tests {
		output := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.expected, output) {
			t.Errorf("expected: %v, got: %v", tc.expected, output)
		}
	}

}

//"FahrenheitToCelsius(%v) = %v; want %v"
