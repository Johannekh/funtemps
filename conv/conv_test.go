package conv

import (
	"math"
	"testing"
)

// / 1
func TestKelvinToFarhrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 500, want: 440.33},
	}

	for _, tc := range tests {
		got := KelvinToFarhrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her
// ...
/// 2

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 500, want: 226.85},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

// / 3
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

// / 4
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 500, want: 773, 15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

/// 5

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 500, want: 533.15},
	}

	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}

/// 6

func TestCelsiusToFarhrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 500, want: 932},
	}

	for _, tc := range tests {
		got := CelsiusToFarhrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-12) {
			t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
		}
	}
}
func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
		return difference < error
	}

	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference / math.Abs(b)) < error
}
