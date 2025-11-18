package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kyanite/prism/internal/color"
	"github.com/kyanite/prism/internal/palette"
	"github.com/kyanite/prism/internal/theme"
)

// GeneratorModel represents the palette generator screen
type GeneratorModel struct {
	themeManager   *theme.Manager
	styles         Styles
	baseColorInput string
	selectedRule   int
	rules          []palette.HarmonyRule
	generatedPalette *palette.Palette
	err            string
}

// NewGeneratorModel creates a new generator model
func NewGeneratorModel(tm *theme.Manager) GeneratorModel {
	return GeneratorModel{
		themeManager:   tm,
		styles:         NewStyles(tm.CurrentTheme()),
		baseColorInput: "#FF0080",
		selectedRule:   0,
		rules:          palette.AllRules(),
	}
}

// Init initializes the generator
func (m GeneratorModel) Init() tea.Cmd {
	return nil
}

// Update handles generator messages
func (m GeneratorModel) Update(msg tea.Msg) (GeneratorModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.selectedRule > 0 {
				m.selectedRule--
			}
		case "down", "j":
			if m.selectedRule < len(m.rules)-1 {
				m.selectedRule++
			}
		case "enter", " ":
			return m, m.generate()
		}
	}

	return m, nil
}

// View renders the generator
func (m GeneratorModel) View() string {
	styles := NewStyles(m.themeManager.CurrentTheme())

	var b strings.Builder

	// Title
	title := styles.Title.Render("Palette Generator")
	b.WriteString(title)
	b.WriteString("\n\n")

	// Base color
	b.WriteString(styles.Primary.Render("Base Color: "))
	b.WriteString(m.baseColorInput)
	b.WriteString("\n\n")

	// Harmony rules
	b.WriteString(styles.Secondary.Render("Select Harmony Rule:"))
	b.WriteString("\n")

	for i, rule := range m.rules {
		style := styles.Unselected
		cursor := "  "
		if i == m.selectedRule {
			style = styles.Selected
			cursor = "▸ "
		}

		line := fmt.Sprintf("%s%s", cursor, rule)
		b.WriteString(style.Render(line))
		b.WriteString("\n")
	}

	b.WriteString("\n")

	// Generated palette
	if m.generatedPalette != nil {
		b.WriteString(styles.Success.Render("Generated Palette:"))
		b.WriteString("\n")

		for i, c := range m.generatedPalette.Colors {
			swatch := lipgloss.NewStyle().
				Background(lipgloss.Color(c.Hex)).
				Padding(0, 2).
				Render("██")

			line := fmt.Sprintf("%d. %s %s", i+1, swatch, c.Hex)
			b.WriteString(line)
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	// Error
	if m.err != "" {
		b.WriteString(styles.Error.Render("Error: " + m.err))
		b.WriteString("\n\n")
	}

	// Help
	help := styles.Muted.Render("↑/↓: Select Rule • Enter: Generate • Esc: Menu")
	b.WriteString(help)

	// Wrap in border
	content := styles.Border.Width(ContentWidth).Render(b.String())

	return lipgloss.Place(ScreenWidth, ScreenHeight, lipgloss.Center, lipgloss.Center, content)
}

// generate generates a palette
func (m GeneratorModel) generate() tea.Cmd {
	return func() tea.Msg {
		baseColor, err := color.ParseHex(m.baseColorInput)
		if err != nil {
			m.err = "Invalid hex color"
			return nil
		}

		pal, err := palette.Generate(baseColor, m.rules[m.selectedRule])
		if err != nil {
			m.err = err.Error()
			return nil
		}

		m.generatedPalette = &pal
		m.err = ""
		return nil
	}
}
