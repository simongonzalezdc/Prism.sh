# prism.sh - Product Requirements Document

**Version:** 1.0  
**Date:** November 2025  
**Owner:** Simon | Kyanite Suite  
**Status:** READY FOR IMPLEMENTATION  
**Target:** Independent GitHub Repository + v1.0 Release

---

## Executive Summary

**prism.sh** is a terminal-based color palette design and theory tool. It enables designers, artists, and creatives to explore color relationships, generate harmonious palettes, and understand color theory—entirely within the command line.

### Value Proposition

- **For Designers:** Quickly explore color relationships and generate harmonious palettes
- **For Creatives:** Match visual concepts to color choices instantly
- **For Developers:** Generate color palettes programmatically

### Target Users

1. **Primary:** Terminal-native designers and visual artists
2. **Secondary:** Developers building color systems
3. **Tertiary:** Creative professionals exploring terminal workflows

### Positioning

> "Color design in your terminal—beautiful, fast, and perfectly on-brand."

---

## Product Vision

### What It Is

An interactive color design tool with educational components. Create palettes, understand color theory, validate accessibility, and export in multiple formats—all without leaving the terminal.

### What It Is NOT

- Not a full design suite (use Figma/Illustrator for comprehensive design)
- Not a photo color extractor (use dedicated image analysis tools)
- Not a complete color accessibility suite (has basic WCAG checks only)

### Kyanite Suite Integration

**prism.sh is standalone but designed to integrate with future Kyanite tools:**
- Export palettes as themes for other tools
- Use as a shared library for color workflows
- Reference in documentation and tutorials

---

## Core Features

### Feature 1: Interactive Color Wheel / Circle of Colors

**Purpose:** Visual exploration of hue relationships

**Core Mechanics:**
- 360-degree hue circle displayed in terminal
- Navigate with arrow keys or vim keys (hjkl)
- Display current color: hex, RGB, HSL, HSV
- Show complementary, analogous, triadic colors automatically
- Animated highlight showing active hue
- Export current color to clipboard or file

**Technical Requirements:**
- ASCII/Unicode visualization (no graphics)
- Real-time color calculations
- Terminal-aware color rendering (256-color or truecolor detection)
- Responsive to terminal resize

**Success Criteria:**
- ✅ Displays full hue spectrum in <200 chars width
- ✅ Navigate 360 degrees smoothly
- ✅ Display all formats instantly (<100ms)
- ✅ Works in 80x24 terminal minimum

**Status:** Core feature, MUST ship v1.0

---

### Feature 2: Palette Generator

**Purpose:** Generate harmonious color palettes based on mathematical rules

**Core Mechanics:**
- **Input:** Base color (hex/name) + harmony rule
- **Harmony Rules:**
  - Monochromatic (tints & shades of single hue)
  - Complementary (opposite hue)
  - Analogous (adjacent hues ±30°)
  - Triadic (3 evenly spaced hues 120° apart)
  - Tetradic (4 hues, 2 pairs opposite 180° apart)
  - Split-Complementary (complementary + one adjacent)
  - Square (4 evenly spaced hues 90° apart)
- **Output:** 3-7 color palette with visual swatches
- **Export:** JSON, CSS variables, TOML, Kyanite theme format

**Technical Requirements:**
- HSL/HSV color space calculations
- Harmonic relationships precise to ±5° hue error
- Real-time generation (<500ms)
- Validation that colors are perceptually harmonious

**Success Criteria:**
- ✅ Generates palettes for all 7 rules
- ✅ Generated colors are visually harmonious
- ✅ Visual display shows all palette colors with hex codes
- ✅ Export works for all 4 formats

**Status:** Core feature, MUST ship v1.0

---

### Feature 3: Named Color Database

**Purpose:** Quick reference to named colors and variations

**Core Mechanics:**
- Search by name ("red", "ocean", "forest")
- Display exact color values and variations
- Show color family (warm/cool, saturated/desaturated)
- Generate tints (lighter), shades (darker), tones (desaturated)
- Support CSS color names + extended set
- Quick insertion into active palette

**Technical Requirements:**
- Pre-loaded database (~500 named colors)
- Fast fuzzy search (<100ms)
- Variations calculated in HSL space
- Works offline

**Success Criteria:**
- ✅ Search finds 5-10 matching colors
- ✅ Display name, hex, RGB, HSL
- ✅ Show 5 tint/shade variations
- ✅ Fuzzy matching works intuitively

**Status:** Core feature, MUST ship v1.0

---

### Feature 4: Color Theory Education

**Purpose:** Teach color relationships and theory in context

**Core Mechanics:**
- **Interactive Lessons:** "What are complementary colors?", "Show analogous relationships", "Warm vs cool"
- **Quick Tips:** Display theory snippets inline while working
- **Visual Examples:** Show concepts with actual color swatches
- **Glossary:** Color theory terms with definitions
- **No Internet:** All content bundled (offline first)

**Content Required (v1.0):**
- 5+ foundational lessons
- 15+ glossary entries
- 50+ visual examples

**Success Criteria:**
- ✅ At least 5 lessons implemented
- ✅ Each lesson includes visual examples
- ✅ 2-3 minutes to read/interact per lesson
- ✅ Content is accurate and beginner-friendly

**Status:** Core feature, MUST ship v1.0

---

### Feature 5: WCAG Accessibility Validation

**Purpose:** Validate color combinations meet accessibility standards

**Core Mechanics:**
- **Input:** Two colors (foreground + background)
- **Output:** 
  - Contrast ratio (e.g., "4.5:1")
  - WCAG compliance level (AA/AAA/Fail)
  - Recommendation if failing
- **Batch Checking:** Validate multiple color pairs
- **Warnings:** Flag problematic combinations in palette

**Technical Requirements:**
- WCAG 2.1 contrast formula implementation
- Accurate to ±0.05 contrast ratio
- Works for all color formats

**Success Criteria:**
- ✅ Contrast ratio accurate within ±0.05
- ✅ WCAG AA/AAA determination correct
- ✅ Provides actionable recommendations
- ✅ Supports batch checking

**Status:** Core feature, MUST ship v1.0

---

### Feature 6: Palette Management

**Purpose:** Save, organize, and recall color palettes

**Core Mechanics:**
- **Save:** Store palette with name/description/tags
- **List:** Show saved palettes with color preview
- **Load:** Recall and edit previous palettes
- **Delete:** Remove saved palettes with confirmation
- **Export:** Save to file (JSON/CSS/TOML)
- **Import:** Load palette from file

**Storage:**
- Location: `~/.config/prism/palettes/`
- Format: JSON with metadata
- Auto-organize by creation date

**Technical Requirements:**
- Fast file I/O (<100ms)
- Graceful error handling
- Support multiple export formats
- Handle corrupted files gracefully

**Success Criteria:**
- ✅ Save/load cycle preserves colors exactly
- ✅ List shows 20+ previous palettes
- ✅ Export generates valid formats
- ✅ Import from external files works

**Status:** Core feature, MUST ship v1.0

---

### Feature 7: Kyanite Suite Integration & Export

**Purpose:** Export colors for use in other Kyanite tools

**Core Mechanics:**
- **Export to Theme:** Convert palette to Kyanite theme format
- **Standard Format:** All exports use consistent structure
- **Interoperability:** Exported colors work in other tools
- **Documentation:** Format is documented and versioned

**Export Format:**
```json
{
  "name": "My Palette",
  "kyanite_version": "1.0",
  "created_at": "2025-11-15T10:30:00Z",
  "palette": {
    "primary": "#FF0080",
    "secondary": "#00D4FF",
    "accent": "#FFE600",
    "background": "#0D0221",
    "text": "#F0F3FF",
    "success": "#39FF14"
  }
}
```

**Success Criteria:**
- ✅ Exports are valid JSON
- ✅ Format documented
- ✅ Can be imported by other tools (future)

**Status:** Core feature, MUST ship v1.0

---

## Feature Matrix

| Feature | v1.0 | v1.1 | v2.0 |
|---------|------|------|------|
| **Color Wheel** | ✅ | | |
| **Palette Generator** | ✅ | | |
| **Named Colors** | ✅ | | |
| **Color Theory** | ✅ | | |
| **WCAG Checker** | ✅ | | |
| **Palette Manager** | ✅ | | |
| **Suite Export** | ✅ | | |
| **Color Mixing** | | ✅ | |
| **Mood-to-Palette** | | ✅ | |
| **Colorblind Sim** | | | ✅ |
| **Image Analysis** | | | ✅ |

---

## User Workflows

### Workflow 1: Designer Creating Brand Palette

```
1. Start prism
2. Navigate to hue (electric blue)
3. Generate tetradic palette
4. View 4-color palette
5. Check WCAG accessibility
6. Export as JSON
Time: 2-3 minutes
```

---

### Workflow 2: Learning Color Theory

```
1. Open lessons menu
2. Read "Complementary Colors"
3. See interactive example
4. Explore with color wheel
5. Understand concept
Time: 5-10 minutes per lesson
```

---

### Workflow 3: Checking Accessibility

```
1. Open WCAG checker
2. Input foreground: #FF0080
3. Input background: #0D0221
4. See: "5.2:1 - WCAG AAA ✓"
5. Adjust if needed
Time: 30 seconds
```

---

## Out of Scope

Explicitly NOT v1.0:

- Photo/image analysis
- Gradient generation
- Colorblind simulation
- AI naming ("poet mode")
- Real-time collaboration
- Cloud sync
- Animation or transitions
- Grammar checking

---

## Success Metrics

### Technical Metrics

- **Startup:** <1 second
- **Color wheel render:** <100ms
- **Palette generation:** <500ms
- **Search:** <100ms
- **Memory:** <30MB idle
- **File I/O:** <100ms

### Feature Completion

- ✅ All 7 core features implemented
- ✅ All acceptance criteria met
- ✅ 0 known critical bugs

### Code Quality

- ✅ No TODO comments
- ✅ Error handling complete
- ✅ Tests for critical paths
- ✅ Documentation comprehensive

---

## Acceptance Criteria

### Must Have (v1.0)

- [ ] Color wheel displays all 360 hues visually
- [ ] Arrow key navigation smooth
- [ ] All 7 harmony rules work correctly
- [ ] Generated palettes are harmonious
- [ ] Save/load preserves colors exactly
- [ ] WCAG checker accurate
- [ ] Export to JSON, CSS, TOML works
- [ ] 10 Kyanite themes applied
- [ ] Ctrl+Shift+T theme switcher works
- [ ] All universal shortcuts implemented
- [ ] Help system complete
- [ ] No panics or crashes
- [ ] Works on 80x24 terminal
- [ ] README complete
- [ ] ARCHITECTURE complete

### Performance Targets

- [ ] Startup: <1s
- [ ] Operations: <500ms
- [ ] Memory: <30MB idle

### Quality Targets

- [ ] 0 panic conditions
- [ ] 100% error handling
- [ ] All workflows <5 minutes

---

## Timeline

| Phase | Duration | Deliverables |
|-------|----------|--------------|
| **Foundation** | 2 days | Project structure, theme system |
| **Features** | 3 days | All 7 features with UI |
| **Polish** | 1 day | Docs, testing, refinement |
| **Release** | 1 day | GitHub repo, v1.0 release |
| **TOTAL** | ~7 days | Launch ready |

---

## Approval

**Product Owner:** Simon  
**Status:** ✅ APPROVED FOR IMPLEMENTATION  
**Start Date:** When ready  
**Target Launch:** v1.0 on GitHub

---

**This is a standalone, independent product. prism.sh can be built and released completely independently from syntax.sh or any other Kyanite tool.**

**Next step:** Review Technical Design Document (TDD)
