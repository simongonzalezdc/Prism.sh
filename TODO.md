# Prism.sh Implementation TODO

This file tracks all remaining implementation tasks. Check off items as you complete them.

## Phase 1: Foundation ✅ COMPLETE

- [x] Project structure
- [x] Go module initialization
- [x] Color package (types, conversions)
- [x] WCAG package (contrast calculations)
- [x] Palette package (harmony rules)
- [x] Theme package (10 Kyanite themes)
- [x] Export package scaffolds
- [x] Storage package scaffolds
- [x] Basic main.go demo

## Phase 2: Bubble Tea TUI (High Priority)

### Core App Structure
- [ ] internal/app/model.go - Root Bubble Tea model
- [ ] internal/app/nav.go - Navigation router between screens
- [ ] internal/app/keys.go - Global keybindings (Ctrl+Q, Ctrl+H, etc.)

### UI Screens
- [ ] internal/ui/menu.go - Main menu screen
- [ ] internal/ui/wheel.go - Color wheel screen
  - [ ] Color circle visualization
  - [ ] Arrow key navigation
  - [ ] Display current color info
  - [ ] Show complementary/analogous/triadic
- [ ] internal/ui/generator.go - Palette generator screen
  - [ ] Select harmony rule
  - [ ] Generate palette
  - [ ] Display color swatches
- [ ] internal/ui/theory.go - Color theory education screen
  - [ ] Lesson browser
  - [ ] Visual examples
  - [ ] Glossary
- [ ] internal/ui/checker.go - WCAG checker screen
  - [ ] Input two colors
  - [ ] Display contrast ratio
  - [ ] Show AA/AAA status
  - [ ] Provide recommendations
- [ ] internal/ui/manager.go - Palette manager screen
  - [ ] List saved palettes
  - [ ] Load/delete palettes
  - [ ] Export options

### UI Components
- [ ] internal/ui/components/palette_display.go
- [ ] internal/ui/components/color_swatch.go
- [ ] internal/ui/components/help.go

## Phase 3: Data & Content

### Named Colors Database
- [ ] data/colors.json - CSS Level 4 + X11 colors
  - [ ] Parse CSS color names
  - [ ] Add X11 color names
  - [ ] Include hex values and metadata
- [ ] internal/color/names.go - Named color search
  - [ ] Load colors.json
  - [ ] Fuzzy search implementation
  - [ ] Color lookup by name

### Color Theory Content
- [ ] data/lessons.json - Theory lessons
  - [ ] Lesson 1: "Complementary Colors"
  - [ ] Lesson 2: "Analogous Colors"
  - [ ] Lesson 3: "Warm vs Cool"
  - [ ] Lesson 4: "Saturation & Desaturation"
  - [ ] Lesson 5: "Tints, Shades, Tones"
- [ ] internal/theory/lessons.go - Lesson management
- [ ] internal/theory/glossary.go - Color terms glossary

## Phase 4: Storage & Export

### Complete Export Implementations
- [ ] internal/export/json.go - Finalize JSON export
- [ ] internal/export/css.go - Test CSS variable export
- [ ] internal/export/toml.go - Test TOML export
- [ ] internal/export/theme.go - Test Kyanite format

### Complete Storage Implementations
- [ ] internal/storage/config.go - TOML config parsing
- [ ] internal/storage/palettes.go - Test save/load
- [ ] internal/storage/lock.go - Cross-platform file locking
  - [ ] Linux/macOS: Use golang.org/x/sys/unix
  - [ ] Windows: Use LockFileEx
- [ ] internal/storage/history.go - Recent colors tracking

### Clipboard Support
- [ ] internal/clipboard/clipboard.go - Main interface
- [ ] internal/clipboard/platform_linux.go - xclip/xsel/wl-clipboard
- [ ] internal/clipboard/platform_darwin.go - pbcopy
- [ ] internal/clipboard/platform_windows.go - clip.exe

## Phase 5: Testing

### Unit Tests (Target >70% coverage)
- [ ] tests/color_test.go
  - [ ] Test all color conversions
  - [ ] Test color operations
  - [ ] Test edge cases (black, white, gray)
- [ ] tests/palette_test.go
  - [ ] Test all 7 harmony rules
  - [ ] Test angle accuracy (±5°)
  - [ ] Test palette contrast validation
- [ ] tests/wcag_test.go
  - [ ] Test WCAG formula accuracy
  - [ ] Test white on black = 21:1
  - [ ] Test AA/AAA thresholds
- [ ] tests/export_test.go
  - [ ] Test JSON export validity
  - [ ] Test CSS export formatting
  - [ ] Test TOML export formatting
- [ ] tests/storage_test.go
  - [ ] Test save/load cycle
  - [ ] Test atomic writes
  - [ ] Test file locking
- [ ] tests/ui_test.go
  - [ ] Test Bubble Tea model state transitions
  - [ ] Test navigation
  - [ ] Test key handling

### Integration Tests
- [ ] End-to-end palette creation workflow
- [ ] Export and import workflow
- [ ] Theme switching workflow

## Phase 6: Documentation

### Required Files
- [ ] ARCHITECTURE.md - Module breakdown and design
- [ ] CONTRIBUTING.md - How to contribute
- [ ] LICENSE - MIT License (already created)
- [ ] SECURITY.md - Security policy
- [ ] CHANGELOG.md - Version history

### Code Documentation
- [ ] Add godoc comments to all public functions
- [ ] Add package-level documentation
- [ ] Add usage examples in docs

## Phase 7: CI/CD

### GitHub Actions Workflows
- [ ] .github/workflows/test.yml
  - [ ] Run tests on push/PR
  - [ ] Check coverage >70%
  - [ ] Upload to codecov
- [ ] .github/workflows/build.yml
  - [ ] Build for Linux, macOS, Windows
  - [ ] Build for amd64 and arm64
  - [ ] Upload artifacts
- [ ] .github/workflows/release.yml
  - [ ] Trigger on version tags
  - [ ] Build all platforms
  - [ ] Create GitHub release
  - [ ] Upload binaries

## Phase 8: Polish

### Features
- [ ] 80x24 terminal minimum support
- [ ] Help system (Ctrl+H)
- [ ] Theme switcher (Ctrl+Shift+T)
- [ ] Error recovery (no panics)
- [ ] Performance optimization (<1s startup, <500ms operations)

### User Experience
- [ ] Comprehensive help text
- [ ] User-friendly error messages
- [ ] Loading indicators for slow operations
- [ ] Confirmation dialogs for destructive actions

## Phase 9: Release

- [ ] Finalize README.md
- [ ] Create demo GIF/screenshots
- [ ] Tag v1.0.0
- [ ] Create GitHub release
- [ ] Publish announcement

## Performance Targets

- [ ] Startup: <1s
- [ ] Palette generation: <500ms
- [ ] UI interaction: <100ms
- [ ] Search: <100ms
- [ ] Memory idle: <50MB

## Acceptance Criteria (v1.0)

From PRD - all must be checked before release:

- [ ] Color wheel displays all 360 hues visually
- [ ] Arrow key navigation smooth (<50ms response)
- [ ] All 7 harmony rules work correctly (±5° accuracy)
- [ ] Generated palettes maintain 3:1 minimum contrast
- [ ] Save/load preserves colors exactly
- [ ] WCAG checker accurate (±0.05 ratio)
- [ ] Export to JSON, CSS, TOML, Kyanite works
- [ ] 10 Kyanite app themes applied
- [ ] Ctrl+Shift+T theme switcher works
- [ ] All universal shortcuts implemented
- [ ] Help system complete
- [ ] No panics in application code
- [ ] Functional UI in 80x24 terminal
- [ ] README complete
- [ ] ARCHITECTURE.md complete
- [ ] LICENSE included
- [ ] CONTRIBUTING.md included
- [ ] Clipboard support works on Linux/macOS/Windows

---

**Priority Order:**
1. Bubble Tea TUI (Phase 2)
2. Named colors & theory content (Phase 3)
3. Complete storage & export (Phase 4)
4. Testing (Phase 5)
5. Documentation (Phase 6)
6. CI/CD (Phase 7)
7. Polish & Release (Phases 8-9)
