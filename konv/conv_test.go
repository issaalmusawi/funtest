package conv

import (
	"math"
	"testing"
)

// Test for fahrenheit to Celsius
func TestFahrenheitToCelsius(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{32, 0},
		{68, 20},
		{100, 37.77777777777778},
		{212, 100},
	}

	for _, tc := range testCases {
		output := FahrenheitToCelsius(tc.input)
		if math.Abs(output-tc.expected) > 0.0000001 {
			t.Errorf("FahrenheitToCelsius(%v) = %v; want %v", tc.input, output, tc.expected)
		}
	}
}

// test for celsius to fahrenheit
func TestCelsiusToFahrenheit(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{56.4, 133.52},
		{20, 68},
		{100, 212},
	}
	for _, tc := range testCases {
		output := CelsiusToFahrenheit(tc.input)
		if math.Abs(output-tc.expected) > 0.0000001 {
			t.Errorf("CelsiusToFahrenheit(%v) = %v; want %v", tc.input, output, tc.expected)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{10, -441.67},
		{270, 26.33},
		{0, -459.67},
	}
	for _, tc := range testCases {
		output := KelvinToFahrenheit(tc.input)
		if math.Abs(output-tc.expected) > 0.0000001 {
			t.Errorf("KelvinToFahrenheit(%v) = %v; want %v", tc.input, output, tc.expected)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{10, 283.15},
		{50, 323.15},
	}
	for _, tc := range testCases {
		output := CelsiusToKelvin(tc.input)
		if math.Abs(output-tc.expected) > 0.0000001 {
			t.Errorf("CelsiusToKelvin(%v) = %v; want %v", tc.input, output, tc.expected)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{50, 283.15},
		{20, 266.48333333},
	}
	for _, tc := range testCases {
		output := FahrenheitToKelvin(tc.input)
		if math.Abs(output-tc.expected) > 0.0000001 {
			t.Errorf("FahrenheitToKelvin(%v) = %v; want %v", tc.input, output, tc.expected)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{20, -253.15},
		{0, -273.15},
		{50, -223.15},
	}
	for _, tc := range testCases {
		output := KelvinToCelsius(tc.input)
		if math.Abs(output-tc.expected) > 0.0000001 {
			t.Errorf("KelvinToCelsius(%v) = %v; want%v", tc.input, output, tc.expected)
		}
	}
}
