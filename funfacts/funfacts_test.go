package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	type test struct {
		input    string
		expected []Fact
	}
	tests := []test{
		{input: "Sun", expected: []Fact{
			{Text: "temperatur i solens kjerne", Value: 15000000, Type: "C"},
			{Text: "temperatur på ytre lag av solen", Value: 5778, Type: "K"}}},

		{input: "Terra", expected: []Fact{
			{Text: "høyeste temperatur målt på jordens overflate", Value: 56.7, Type: "C"},
			{Text: "laveste temperatur målt på jordens overflate", Value: -89.4, Type: "C"},
			{Text: "temperatur i jordens indre kjerne", Value: 9392, Type: "K"}}},

		{input: "Luna", expected: []Fact{
			{Text: "temperatur på månens overflate om natten", Value: -183, Type: "C"},
			{Text: "temperatur på månens overflate om dagen", Value: 106, Type: "C"}}},
	}

	for _, tc := range tests {
		output := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.expected, output) {
			t.Errorf("expected: %v, got: %v", tc.expected, output)
		}
	}

}

//"FahrenheitToCelsius(%v) = %v; want %v"
