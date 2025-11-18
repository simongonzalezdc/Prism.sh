package tests

import (
	"math"
	"testing"

	"github.com/kyanite/prism/internal/color"
)

func TestParseHex(t *testing.T) {
	tests := []struct {
		hex     string
		wantRGB color.RGB
		wantErr bool
	}{
		{"#FF0000", color.RGB{R: 255, G: 0, B: 0}, false},
		{"#00FF00", color.RGB{R: 0, G: 255, B: 0}, false},
		{"#0000FF", color.RGB{R: 0, G: 0, B: 255}, false},
		{"#FFFFFF", color.RGB{R: 255, G: 255, B: 255}, false},
		{"#000000", color.RGB{R: 0, G: 0, B: 0}, false},
		{"FF0000", color.RGB{R: 255, G: 0, B: 0}, false}, // Without #
		{"#GGGGGG", color.RGB{}, true},          // Invalid
		{"#FFF", color.RGB{}, true},             // Too short
	}

	for _, tt := range tests {
		got, err := color.ParseHex(tt.hex)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseHex(%q) error = %v, wantErr %v", tt.hex, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && got.RGB != tt.wantRGB {
			t.Errorf("ParseHex(%q) = %v, want %v", tt.hex, got.RGB, tt.wantRGB)
		}
	}
}

func TestRGBToHex(t *testing.T) {
	tests := []struct {
		r, g, b int
		want    string
	}{
		{255, 0, 0, "#FF0000"},
		{0, 255, 0, "#00FF00"},
		{0, 0, 255, "#0000FF"},
		{255, 255, 255, "#FFFFFF"},
		{0, 0, 0, "#000000"},
		{128, 128, 128, "#808080"},
	}

	for _, tt := range tests {
		got := color.RGBToHex(tt.r, tt.g, tt.b)
		if got != tt.want {
			t.Errorf("RGBToHex(%d, %d, %d) = %q, want %q", tt.r, tt.g, tt.b, got, tt.want)
		}
	}
}

func TestRGBToHSL(t *testing.T) {
	tests := []struct {
		name string
		rgb  color.RGB
		want color.HSL
	}{
		{"Red", color.RGB{R: 255, G: 0, B: 0}, color.HSL{H: 0, S: 100, L: 50}},
		{"Green", color.RGB{R: 0, G: 255, B: 0}, color.HSL{H: 120, S: 100, L: 50}},
		{"Blue", color.RGB{R: 0, G: 0, B: 255}, color.HSL{H: 240, S: 100, L: 50}},
		{"White", color.RGB{R: 255, G: 255, B: 255}, color.HSL{H: 0, S: 0, L: 100}},
		{"Black", color.RGB{R: 0, G: 0, B: 0}, color.HSL{H: 0, S: 0, L: 0}},
		{"Gray", color.RGB{R: 128, G: 128, B: 128}, color.HSL{H: 0, S: 0, L: 50}},
	}

	for _, tt := range tests {
		got := color.RGBToHSL(tt.rgb)

		// Allow small tolerance for floating point
		if math.Abs(got.H-tt.want.H) > 1 || math.Abs(got.S-tt.want.S) > 1 || math.Abs(got.L-tt.want.L) > 1 {
			t.Errorf("%s: RGBToHSL(%v) = %v, want %v", tt.name, tt.rgb, got, tt.want)
		}
	}
}

func TestHSLToRGB(t *testing.T) {
	tests := []struct {
		name string
		hsl  color.HSL
		want color.RGB
	}{
		{"Red", color.HSL{H: 0, S: 100, L: 50}, color.RGB{R: 255, G: 0, B: 0}},
		{"Green", color.HSL{H: 120, S: 100, L: 50}, color.RGB{R: 0, G: 255, B: 0}},
		{"Blue", color.HSL{H: 240, S: 100, L: 50}, color.RGB{R: 0, G: 0, B: 255}},
		{"White", color.HSL{H: 0, S: 0, L: 100}, color.RGB{R: 255, G: 255, B: 255}},
		{"Black", color.HSL{H: 0, S: 0, L: 0}, color.RGB{R: 0, G: 0, B: 0}},
	}

	for _, tt := range tests {
		got := color.HSLToRGB(tt.hsl.H, tt.hsl.S, tt.hsl.L)

		// Allow small tolerance for rounding
		if abs(got.R-tt.want.R) > 1 || abs(got.G-tt.want.G) > 1 || abs(got.B-tt.want.B) > 1 {
			t.Errorf("%s: HSLToRGB(%v) = %v, want %v", tt.name, tt.hsl, got, tt.want)
		}
	}
}

func TestColorConversionRoundtrip(t *testing.T) {
	// Test that RGB -> HSL -> RGB is reversible
	tests := []color.RGB{
		{R: 255, G: 0, B: 0},
		{R: 0, G: 255, B: 0},
		{R: 0, G: 0, B: 255},
		{R: 128, G: 64, B: 192},
		{R: 255, G: 128, B: 64},
	}

	for _, original := range tests {
		hsl := color.RGBToHSL(original)
		recovered := color.HSLToRGB(hsl.H, hsl.S, hsl.L)

		if abs(original.R-recovered.R) > 2 || abs(original.G-recovered.G) > 2 || abs(original.B-recovered.B) > 2 {
			t.Errorf("Roundtrip failed: %v -> %v -> %v", original, hsl, recovered)
		}
	}
}

func TestColorOperations(t *testing.T) {
	c := color.NewFromHSL(180, 100, 50) // Cyan

	// Test Lighten
	lighter := c.Lighten(20)
	if lighter.HSL.L < c.HSL.L {
		t.Errorf("Lighten() did not increase lightness")
	}

	// Test Darken
	darker := c.Darken(20)
	if darker.HSL.L > c.HSL.L {
		t.Errorf("Darken() did not decrease lightness")
	}

	// Test Complement
	complement := c.Complement()
	expectedHue := 0.0 // Cyan (180°) complement is Red (0°/360°)
	if math.Abs(complement.HSL.H-expectedHue) > 1 {
		t.Errorf("Complement() = %.0f°, want %.0f°", complement.HSL.H, expectedHue)
	}

	// Test Saturate
	saturated := c.Saturate(10)
	if saturated.HSL.S < c.HSL.S && c.HSL.S < 100 {
		t.Errorf("Saturate() did not increase saturation")
	}

	// Test Desaturate
	desaturated := c.Desaturate(10)
	if desaturated.HSL.S > c.HSL.S && c.HSL.S > 0 {
		t.Errorf("Desaturate() did not decrease saturation")
	}
}

func TestColorTemperature(t *testing.T) {
	warm := color.NewFromHSL(30, 100, 50)  // Orange
	cool := color.NewFromHSL(210, 100, 50) // Blue

	if !warm.IsWarm() {
		t.Error("Orange should be warm")
	}
	if !cool.IsCool() {
		t.Error("Blue should be cool")
	}
	if warm.IsCool() {
		t.Error("Orange should not be cool")
	}
	if cool.IsWarm() {
		t.Error("Blue should not be warm")
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
