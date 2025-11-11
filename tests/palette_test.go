package tests

import (
	"math"
	"testing"

	"github.com/kyanite/prism/internal/color"
	"github.com/kyanite/prism/internal/palette"
)

func TestPaletteGeneration(t *testing.T) {
	baseColor := color.NewFromHSL(180, 100, 50) // Cyan

	tests := []struct {
		rule          palette.HarmonyRule
		expectedCount int
	}{
		{palette.Monochromatic, 5},
		{palette.Complementary, 2},
		{palette.Analogous, 3},
		{palette.Triadic, 3},
		{palette.Tetradic, 4},
		{palette.SplitComplementary, 3},
		{palette.Square, 4},
	}

	for _, tt := range tests {
		t.Run(string(tt.rule), func(t *testing.T) {
			pal, err := palette.Generate(baseColor, tt.rule)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			if len(pal.Colors) != tt.expectedCount {
				t.Errorf("Generate() produced %d colors, want %d", len(pal.Colors), tt.expectedCount)
			}

			if pal.HarmonyRule != string(tt.rule) {
				t.Errorf("HarmonyRule = %s, want %s", pal.HarmonyRule, tt.rule)
			}
		})
	}
}

func TestTriadicAngles(t *testing.T) {
	baseColor := color.NewFromHSL(0, 100, 50) // Red at 0°

	pal, err := palette.Generate(baseColor, palette.Triadic)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	// Triadic should be 120° apart
	expected := []float64{0, 120, 240}
	for i, c := range pal.Colors {
		if math.Abs(c.HSL.H-expected[i]) > 5 {
			t.Errorf("Color %d hue = %.0f°, want %.0f° (±5°)", i, c.HSL.H, expected[i])
		}
	}
}

func TestComplementaryAngles(t *testing.T) {
	baseColor := color.NewFromHSL(30, 100, 50) // Orange at 30°

	pal, err := palette.Generate(baseColor, palette.Complementary)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if len(pal.Colors) != 2 {
		t.Fatalf("Expected 2 colors, got %d", len(pal.Colors))
	}

	// Complementary should be 180° apart
	hue1 := pal.Colors[0].HSL.H
	hue2 := pal.Colors[1].HSL.H

	diff := math.Abs(hue1 - hue2)
	if diff > 360-diff {
		diff = 360 - diff
	}

	if math.Abs(diff-180) > 5 {
		t.Errorf("Complementary colors are %.0f° apart, want 180° (±5°)", diff)
	}
}

func TestAnalogousAngles(t *testing.T) {
	baseColor := color.NewFromHSL(180, 100, 50) // Cyan at 180°

	pal, err := palette.Generate(baseColor, palette.Analogous)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if len(pal.Colors) != 3 {
		t.Fatalf("Expected 3 colors, got %d", len(pal.Colors))
	}

	// Should be base-30°, base, base+30°
	expected := []float64{150, 180, 210}
	for i, c := range pal.Colors {
		if math.Abs(c.HSL.H-expected[i]) > 5 {
			t.Errorf("Color %d hue = %.0f°, want %.0f° (±5°)", i, c.HSL.H, expected[i])
		}
	}
}

func TestMonochromaticLightness(t *testing.T) {
	baseColor := color.NewFromHSL(0, 100, 50) // Red

	pal, err := palette.Generate(baseColor, palette.Monochromatic)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	// All colors should have same hue
	baseHue := pal.Colors[0].HSL.H
	for i, c := range pal.Colors {
		if math.Abs(c.HSL.H-baseHue) > 5 {
			t.Errorf("Color %d hue = %.0f°, want %.0f° (same hue)", i, c.HSL.H, baseHue)
		}
	}

	// Should have varying lightness values
	lightnesses := make(map[int]bool)
	for _, c := range pal.Colors {
		lightnesses[int(c.HSL.L)] = true
	}

	if len(lightnesses) < 3 {
		t.Errorf("Monochromatic palette should have at least 3 distinct lightness values, got %d", len(lightnesses))
	}
}

func TestAllRules(t *testing.T) {
	rules := palette.AllRules()

	if len(rules) != 7 {
		t.Errorf("AllRules() returned %d rules, want 7", len(rules))
	}

	// Check all expected rules are present
	expected := map[palette.HarmonyRule]bool{
		palette.Monochromatic:      true,
		palette.Complementary:      true,
		palette.Analogous:          true,
		palette.Triadic:            true,
		palette.Tetradic:           true,
		palette.SplitComplementary: true,
		palette.Square:             true,
	}

	for _, rule := range rules {
		if !expected[rule] {
			t.Errorf("Unexpected rule: %s", rule)
		}
		delete(expected, rule)
	}

	if len(expected) > 0 {
		t.Errorf("Missing rules: %v", expected)
	}
}
