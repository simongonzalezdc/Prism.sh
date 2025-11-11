package palette

import (
	"fmt"
	"math"
	"time"

	"github.com/kyanite/prism/internal/color"
	"github.com/kyanite/prism/internal/wcag"
)

// Generate creates a palette based on harmony rule
func Generate(baseColor color.Color, rule HarmonyRule) (Palette, error) {
	palette := Palette{
		ID:          generateID(),
		Name:        fmt.Sprintf("%s Palette", rule),
		BaseColor:   baseColor,
		HarmonyRule: string(rule),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Colors:      []color.Color{},
	}

	switch rule {
	case Monochromatic:
		palette.Colors = GenerateMonochromatic(baseColor)
	case Complementary:
		palette.Colors = GenerateComplementary(baseColor)
	case Analogous:
		palette.Colors = GenerateAnalogous(baseColor)
	case Triadic:
		palette.Colors = GenerateTriadic(baseColor)
	case Tetradic:
		palette.Colors = GenerateTetradic(baseColor)
	case SplitComplementary:
		palette.Colors = GenerateSplitComplementary(baseColor)
	case Square:
		palette.Colors = GenerateSquare(baseColor)
	default:
		return palette, fmt.Errorf("unknown harmony rule: %s", rule)
	}

	return palette, nil
}

// GenerateMonochromatic creates tints and shades of a single hue
func GenerateMonochromatic(base color.Color) []color.Color {
	// TODO: Implement monochromatic generation
	// Algorithm: Create 5 colors with different lightness values
	// - 2 tints: L + 20%, L + 40%
	// - Base color
	// - 2 shades: L - 20%, L - 40%
	colors := []color.Color{
		base.Lighten(40),
		base.Lighten(20),
		base,
		base.Darken(20),
		base.Darken(40),
	}
	return colors
}

// GenerateComplementary creates opposite hue pair
func GenerateComplementary(base color.Color) []color.Color {
	// TODO: Implement complementary generation
	// Algorithm: Base color + color at 180° hue rotation
	complement := color.NewFromHSL(
		math.Mod(base.HSL.H+180, 360),
		base.HSL.S,
		base.HSL.L,
	)

	return []color.Color{base, complement}
}

// GenerateAnalogous creates adjacent hues (±30°)
func GenerateAnalogous(base color.Color) []color.Color {
	// TODO: Implement analogous generation
	// Algorithm: Base + hues at -30° and +30°
	left := color.NewFromHSL(
		math.Mod(base.HSL.H-30+360, 360),
		base.HSL.S,
		base.HSL.L,
	)
	right := color.NewFromHSL(
		math.Mod(base.HSL.H+30, 360),
		base.HSL.S,
		base.HSL.L,
	)

	return []color.Color{left, base, right}
}

// GenerateTriadic creates 3 evenly spaced hues (120° apart)
func GenerateTriadic(base color.Color) []color.Color {
	// TODO: Implement triadic generation
	// Algorithm: Base + hues at +120° and +240°
	second := color.NewFromHSL(
		math.Mod(base.HSL.H+120, 360),
		base.HSL.S,
		base.HSL.L,
	)
	third := color.NewFromHSL(
		math.Mod(base.HSL.H+240, 360),
		base.HSL.S,
		base.HSL.L,
	)

	return []color.Color{base, second, third}
}

// GenerateTetradic creates 4 hues in complementary pairs
func GenerateTetradic(base color.Color) []color.Color {
	// TODO: Implement tetradic generation
	// Algorithm: Two pairs of complementary colors
	// Base, Base+90°, Base+180°, Base+270°
	second := color.NewFromHSL(
		math.Mod(base.HSL.H+90, 360),
		base.HSL.S,
		base.HSL.L,
	)
	third := color.NewFromHSL(
		math.Mod(base.HSL.H+180, 360),
		base.HSL.S,
		base.HSL.L,
	)
	fourth := color.NewFromHSL(
		math.Mod(base.HSL.H+270, 360),
		base.HSL.S,
		base.HSL.L,
	)

	return []color.Color{base, second, third, fourth}
}

// GenerateSplitComplementary creates complement + adjacent colors
func GenerateSplitComplementary(base color.Color) []color.Color {
	// TODO: Implement split-complementary generation
	// Algorithm: Base + colors at 180°±30° (150° and 210°)
	complement1 := color.NewFromHSL(
		math.Mod(base.HSL.H+150, 360),
		base.HSL.S,
		base.HSL.L,
	)
	complement2 := color.NewFromHSL(
		math.Mod(base.HSL.H+210, 360),
		base.HSL.S,
		base.HSL.L,
	)

	return []color.Color{base, complement1, complement2}
}

// GenerateSquare creates 4 evenly spaced hues (90° apart)
func GenerateSquare(base color.Color) []color.Color {
	// TODO: Implement square generation
	// Algorithm: Base + hues at +90°, +180°, +270°
	return GenerateTetradic(base) // Same as tetradic
}

// ValidatePaletteContrast checks if palette meets minimum contrast requirements
func ValidatePaletteContrast(palette Palette) (minContrast float64, ok bool) {
	// TODO: Implement palette contrast validation
	// Algorithm: Check contrast between adjacent colors
	// Must meet 3:1 minimum ratio
	if len(palette.Colors) < 2 {
		return 0, true
	}

	minContrast = 21.0 // Max possible contrast
	for i := 0; i < len(palette.Colors)-1; i++ {
		ratio := wcag.CalculateContrast(palette.Colors[i], palette.Colors[i+1])
		if ratio < minContrast {
			minContrast = ratio
		}
	}

	return minContrast, minContrast >= 3.0
}

// generateID creates a unique palette ID
func generateID() string {
	return fmt.Sprintf("palette_%s", time.Now().Format("20060102_150405"))
}
