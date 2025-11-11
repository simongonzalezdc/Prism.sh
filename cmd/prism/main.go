package main

import (
	"fmt"
	"os"

	"github.com/kyanite/prism/internal/color"
	"github.com/kyanite/prism/internal/palette"
	"github.com/kyanite/prism/internal/wcag"
)

func main() {
	// TODO: Implement full Bubble Tea TUI
	// For now, this is a simple CLI demo showing core functionality

	fmt.Println("╭────────────────────────────────────╮")
	fmt.Println("│         PRISM.SH v0.1.0            │")
	fmt.Println("│    Color Palette Design Tool       │")
	fmt.Println("╰────────────────────────────────────╯")
	fmt.Println()

	// Demo: Create a color
	baseColor, err := color.ParseHex("#FF0080")
	if err != nil {
		fmt.Printf("Error parsing color: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Base Color: %s\n", baseColor.Hex)
	fmt.Printf("RGB: (%d, %d, %d)\n", baseColor.RGB.R, baseColor.RGB.G, baseColor.RGB.B)
	fmt.Printf("HSL: (%.0f°, %.0f%%, %.0f%%)\n", baseColor.HSL.H, baseColor.HSL.S, baseColor.HSL.L)
	fmt.Println()

	// Demo: Generate triadic palette
	fmt.Println("Generating Triadic Palette...")
	pal, err := palette.Generate(baseColor, palette.Triadic)
	if err != nil {
		fmt.Printf("Error generating palette: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Palette: %s (%d colors)\n", pal.Name, len(pal.Colors))
	for i, c := range pal.Colors {
		fmt.Printf("  %d. %s - RGB(%d, %d, %d)\n", i+1, c.Hex, c.RGB.R, c.RGB.G, c.RGB.B)
	}
	fmt.Println()

	// Demo: Check WCAG contrast
	if len(pal.Colors) >= 2 {
		ratio := wcag.CalculateContrast(pal.Colors[0], pal.Colors[1])
		result := wcag.Validate(pal.Colors[0], pal.Colors[1])
		fmt.Printf("Contrast between %s and %s: %s\n",
			pal.Colors[0].Hex, pal.Colors[1].Hex, result.Summary())
		fmt.Printf("Contrast Ratio: %.2f:1\n", ratio)
	}
	fmt.Println()

	fmt.Println("TODO: Full TUI interface coming soon!")
	fmt.Println("This scaffold demonstrates core color math, palette generation, and WCAG calculations.")
	fmt.Println()
	fmt.Println("To build on this:")
	fmt.Println("  1. Implement Bubble Tea TUI in internal/app/")
	fmt.Println("  2. Add all 6 screen UIs in internal/ui/")
	fmt.Println("  3. Implement export formats")
	fmt.Println("  4. Add named colors database")
	fmt.Println("  5. Create color theory lessons")
	fmt.Println("  6. Write comprehensive tests")
}
