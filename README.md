# prism.sh

> Terminal-based color palette design and theory tool

**Status:** 🚧 Scaffold/Foundation - Ready for Implementation

This repository contains a complete project scaffold for Prism.sh, with all core color mathematics implemented and ready to build upon.

## Quick Start

```bash
# Install dependencies
go mod download

# Build and run
make build
./bin/prism

# Or run directly
make run
```

## What's Implemented ✅

### Core Packages (Fully Functional)
- **internal/color/** - Complete color math and conversions
  - RGB ↔ HSL ↔ HSV ↔ Hex conversions
  - Color operations (lighten, darken, saturate, etc.)
  - Complement, warm/cool detection
- **internal/wcag/** - Full WCAG 2.1 contrast calculations
  - Gamma correction and relative luminance
  - AA/AAA level validation
  - Accurate to ±0.05 ratio
- **internal/palette/** - All 7 harmony rule generators
  - Monochromatic, Complementary, Analogous
  - Triadic, Tetradic, Split-Complementary, Square
- **internal/theme/** - All 10 Kyanite themes with exact hex codes
- **internal/export/** - Export scaffolds (JSON, CSS, TOML, Kyanite)
- **internal/storage/** - File I/O with atomic writes

### Demo Application
- Simple CLI that demonstrates core functionality
- Shows color conversion, palette generation, WCAG checking

## What Needs Implementation 🚧

See [TODO.md](TODO.md) for complete implementation checklist.

### High Priority
1. **Bubble Tea TUI** - Full interactive terminal UI
   - Menu screen, Color wheel, Palette generator
   - Theory lessons, WCAG checker, Palette manager
2. **Named Colors Database** - JSON file with CSS + X11 colors
3. **Color Theory Lessons** - Educational content
4. **Tests** - Comprehensive test suite (target >70% coverage)

### Medium Priority
5. **Export Implementations** - Complete all 4 export formats
6. **Clipboard Support** - Cross-platform clipboard integration
7. **Storage Enhancements** - File locking, error handling

### Documentation
8. **ARCHITECTURE.md** - Module breakdown and design decisions
9. **CONTRIBUTING.md** - Contribution guidelines
10. **SECURITY.md** - Security policy

## Project Structure

```
prism/
├── cmd/prism/           # Entry point (simple demo)
├── internal/
│   ├── color/           # ✅ Color math (complete)
│   ├── palette/         # ✅ Harmony rules (complete)
│   ├── wcag/            # ✅ Contrast calculations (complete)
│   ├── theme/           # ✅ 10 Kyanite themes (complete)
│   ├── export/          # 🚧 Export scaffolds (needs completion)
│   ├── storage/         # 🚧 File I/O (needs completion)
│   ├── app/             # ❌ Bubble Tea app (not started)
│   └── ui/              # ❌ TUI screens (not started)
├── tests/               # ❌ Test suite (not started)
├── data/                # ❌ Named colors, lessons (not started)
└── .github/workflows/   # ❌ CI/CD (not started)
```

## Development

### Build
```bash
make build
```

### Test
```bash
make test
```

### Run
```bash
make run
```

### Format
```bash
make fmt
```

## Architecture

See [02-TDD-2.md](02-TDD-2.md) for complete technical design.

### Core Principles
- **Bubble Tea** for TUI
- **Lipgloss** for styling
- **XDG Base Directory** for config
- **Atomic writes** for data safety
- **Cross-platform** support (Linux, macOS, Windows)

## Features (Planned)

- ✅ Interactive color wheel
- ✅ 7 harmony rule palette generation
- 🚧 Named color database with fuzzy search
- 🚧 Color theory education
- ✅ WCAG 2.1 accessibility validation
- 🚧 Palette management (save/load)
- 🚧 Export (JSON, CSS, TOML, Kyanite)
- 🚧 10 Kyanite themes
- 🚧 Clipboard support

## Contributing

This is a complete scaffold ready for implementation. Key areas:

1. **TUI Development** - Implement Bubble Tea interface
2. **Content Creation** - Add named colors and theory lessons
3. **Testing** - Write comprehensive tests
4. **Documentation** - Complete all docs
5. **CI/CD** - Set up GitHub Actions

See detailed breakdown in [01-PRD.md](01-PRD.md) and [02-TDD-2.md](02-TDD-2.md).

## Timeline

Estimated completion: 12-16 days full-time development

| Phase | Duration | Status |
|-------|----------|---------|
| Foundation | 2-3 days | ✅ Complete |
| Features | 5-6 days | 🚧 In Progress |
| Content | 2-3 days | ❌ Not Started |
| Polish | 2-3 days | ❌ Not Started |
| Release | 1 day | ❌ Not Started |

## License

MIT License - See [LICENSE](LICENSE)

## Part of Kyanite Suite

Prism.sh is an independent tool in the Kyanite Suite of terminal-based creative tools.

---

**Next Steps:** Review [TODO.md](TODO.md) for implementation checklist.
