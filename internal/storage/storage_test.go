package storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	color "github.com/kyanite/prism/internal/color"
	palette "github.com/kyanite/prism/internal/palette"
	)

// TestStorageIntegration tests complete storage workflows
func TestStorageIntegration(t *testing.T) {
	// Create temporary test directory
	tmpDir, err := os.MkdirTemp("", "prism-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Override config directory for testing
	originalConfigDir := os.Getenv("XDG_CONFIG_HOME")
	os.Setenv("XDG_CONFIG_HOME", tmpDir)
	defer func() {
		if originalConfigDir == "" {
			os.Unsetenv("XDG_CONFIG_HOME")
		} else {
			os.Setenv("XDG_CONFIG_HOME", originalConfigDir)
		}
	}()

	t.Run("ConfigSaveAndLoad", func(t *testing.T) {
		// Create and save a config
		cfg := &Config{
			Theme:            "custom-theme",
			AutoSave:         true,
			AutoSaveInterval: 60,
		}

		err := SaveConfig(cfg)
		if err != nil {
			t.Fatalf("SaveConfig failed: %v", err)
		}

		// Load the config
		loadedCfg, err := LoadConfig()
		if err != nil {
			t.Fatalf("LoadConfig failed: %v", err)
		}

		// Verify loaded config matches saved config
		if loadedCfg.Theme != cfg.Theme {
			t.Errorf("Theme = %s, want %s", loadedCfg.Theme, cfg.Theme)
		}
		if loadedCfg.AutoSave != cfg.AutoSave {
			t.Errorf("AutoSave = %v, want %v", loadedCfg.AutoSave, cfg.AutoSave)
		}
		if loadedCfg.AutoSaveInterval != cfg.AutoSaveInterval {
			t.Errorf("AutoSaveInterval = %d, want %d", loadedCfg.AutoSaveInterval, cfg.AutoSaveInterval)
		}
	})

	t.Run("DefaultConfig", func(t *testing.T) {
		// Remove config file if it exists
		configDir, _ := GetConfigDir()
		os.RemoveAll(configDir)

		// Load config should return default when file doesn't exist
		cfg, err := LoadConfig()
		if err != nil {
			t.Fatalf("LoadConfig failed: %v", err)
		}

		defaultCfg := DefaultConfig()
		if cfg.Theme != defaultCfg.Theme {
			t.Errorf("Theme = %s, want %s", cfg.Theme, defaultCfg.Theme)
		}
	})

	t.Run("PaletteSaveLoadDelete", func(t *testing.T) {
		// Create a test palette
		baseColor, _ := color.ParseHex("#FF5733")
		pal, _ := palette.Generate(baseColor, palette.Triadic)
		pal.Name = "Test Triadic Palette"
		pal.Description = "Integration test palette"
		pal.Tags = []string{"test", "triadic"}

		// Save palette
		err := SavePalette(pal)
		if err != nil {
			t.Fatalf("SavePalette failed: %v", err)
		}

		// Load palette
		loadedPal, err := LoadPalette(pal.ID)
		if err != nil {
			t.Fatalf("LoadPalette failed: %v", err)
		}

		// Verify palette data
		if loadedPal.Name != pal.Name {
			t.Errorf("Name = %s, want %s", loadedPal.Name, pal.Name)
		}
		if loadedPal.Description != pal.Description {
			t.Errorf("Description = %s, want %s", loadedPal.Description, pal.Description)
		}
		if len(loadedPal.Colors) != len(pal.Colors) {
			t.Errorf("Colors count = %d, want %d", len(loadedPal.Colors), len(pal.Colors))
		}
		if len(loadedPal.Tags) != len(pal.Tags) {
			t.Errorf("Tags count = %d, want %d", len(loadedPal.Tags), len(pal.Tags))
		}

		// Delete palette
		err = DeletePalette(pal.ID)
		if err != nil {
			t.Fatalf("DeletePalette failed: %v", err)
		}

		// Verify palette is deleted
		_, err = LoadPalette(pal.ID)
		if err == nil {
			t.Error("LoadPalette should fail after delete")
		}
	})

	t.Run("ListPalettes", func(t *testing.T) {
		// Clean up any existing palettes
		configDir, _ := GetConfigDir()
		palettesDir := filepath.Join(configDir, "palettes")
		os.RemoveAll(palettesDir)

		// Create multiple test palettes
		palettes := []palette.Palette{}
		rules := []palette.HarmonyRule{
			palette.Monochromatic,
			palette.Complementary,
			palette.Analogous,
		}

		for i, rule := range rules {
			baseColor, _ := color.ParseHex("#FF5733")
			pal, _ := palette.Generate(baseColor, rule)
			pal.Name = string(rule)
			palettes = append(palettes, pal)

			err := SavePalette(pal)
			if err != nil {
				t.Fatalf("SavePalette %d failed: %v", i, err)
			}
		}

		// List all palettes
		loadedPalettes, err := ListPalettes()
		if err != nil {
			t.Fatalf("ListPalettes failed: %v", err)
		}

		// Verify count
		if len(loadedPalettes) != len(palettes) {
			t.Errorf("ListPalettes count = %d, want %d", len(loadedPalettes), len(palettes))
		}

		// Clean up
		for _, pal := range palettes {
			DeletePalette(pal.ID)
		}
	})

	t.Run("AtomicWrite", func(t *testing.T) {
		testFile := filepath.Join(tmpDir, "atomic-test.txt")
		testData := []byte("test data for atomic write")

		err := AtomicWrite(testFile, testData)
		if err != nil {
			t.Fatalf("AtomicWrite failed: %v", err)
		}

		// Verify file exists and has correct content
		data, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatalf("ReadFile failed: %v", err)
		}

		if string(data) != string(testData) {
			t.Errorf("File content = %s, want %s", string(data), string(testData))
		}
	})

	t.Run("ConcurrentPaletteSaves", func(t *testing.T) {
		// Test concurrent saves to ensure atomic writes work correctly
		done := make(chan bool)

		for i := 0; i < 5; i++ {
			go func(index int) {
				baseColor, _ := color.ParseHex("#FF5733")
				pal, _ := palette.Generate(baseColor, palette.Triadic)
				pal.Name = "Concurrent Test"
				err := SavePalette(pal)
				if err != nil {
					t.Errorf("Concurrent SavePalette %d failed: %v", index, err)
				}
				done <- true
			}(i)
		}

		// Wait for all goroutines to complete
		for i := 0; i < 5; i++ {
			<-done
		}
	})
}

// TestStorageErrorHandling tests error conditions
func TestStorageErrorHandling(t *testing.T) {
	t.Run("LoadNonexistentPalette", func(t *testing.T) {
		_, err := LoadPalette("nonexistent-id-12345")
		if err == nil {
			t.Error("LoadPalette should fail for nonexistent palette")
		}
	})

	t.Run("DeleteNonexistentPalette", func(t *testing.T) {
		err := DeletePalette("nonexistent-id-12345")
		if err == nil {
			t.Error("DeletePalette should fail for nonexistent palette")
		}
	})

	t.Run("ListPalettesEmptyDirectory", func(t *testing.T) {
		// Create temporary test directory
		tmpDir, err := os.MkdirTemp("", "prism-empty-*")
		if err != nil {
			t.Fatalf("Failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		// Override config directory
		originalConfigDir := os.Getenv("XDG_CONFIG_HOME")
		os.Setenv("XDG_CONFIG_HOME", tmpDir)
		defer func() {
			if originalConfigDir == "" {
				os.Unsetenv("XDG_CONFIG_HOME")
			} else {
				os.Setenv("XDG_CONFIG_HOME", originalConfigDir)
			}
		}()

		palettes, err := ListPalettes()
		if err != nil {
			t.Fatalf("ListPalettes should not fail on empty directory: %v", err)
		}

		if len(palettes) != 0 {
			t.Errorf("ListPalettes should return empty slice, got %d palettes", len(palettes))
		}
	})
}

// TestPaletteVersioning tests palette updates
func TestPaletteVersioning(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "prism-version-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalConfigDir := os.Getenv("XDG_CONFIG_HOME")
	os.Setenv("XDG_CONFIG_HOME", tmpDir)
	defer func() {
		if originalConfigDir == "" {
			os.Unsetenv("XDG_CONFIG_HOME")
		} else {
			os.Setenv("XDG_CONFIG_HOME", originalConfigDir)
		}
	}()

	// Create and save a palette
	baseColor, _ := color.ParseHex("#FF5733")
	pal, _ := palette.Generate(baseColor, palette.Triadic)
	originalName := "Original Name"
	pal.Name = originalName

	err = SavePalette(pal)
	if err != nil {
		t.Fatalf("SavePalette failed: %v", err)
	}

	// Simulate time passing
	time.Sleep(10 * time.Millisecond)

	// Update the palette
	pal.Name = "Updated Name"
	pal.UpdatedAt = time.Now()
	err = SavePalette(pal)
	if err != nil {
		t.Fatalf("SavePalette (update) failed: %v", err)
	}

	// Load and verify update
	loadedPal, err := LoadPalette(pal.ID)
	if err != nil {
		t.Fatalf("LoadPalette failed: %v", err)
	}

	if loadedPal.Name != "Updated Name" {
		t.Errorf("Updated name = %s, want %s", loadedPal.Name, "Updated Name")
	}

	if !loadedPal.UpdatedAt.After(loadedPal.CreatedAt) {
		t.Error("UpdatedAt should be after CreatedAt")
	}
}
