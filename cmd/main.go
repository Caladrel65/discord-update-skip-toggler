package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Expand ~ to home directory
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	settingsPath := filepath.Join(home, ".config", "discord", "settings.json")

	// Read the file
	data, err := os.ReadFile(settingsPath)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", settingsPath, err)
		return
	}

	// Parse JSON
	var settings map[string]any
	if err := json.Unmarshal(data, &settings); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Toggle the SKIP_HOST_UPDATE, defaulting to "false" if the field doesn't exist for safety
	shouldSkip := false
	if val, ok := settings["SKIP_HOST_UPDATE"]; ok {
		shouldSkip = !val.(bool)
	}

	// Update the settings
	settings["SKIP_HOST_UPDATE"] = shouldSkip

	// Marshal back to JSON
	updated, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	// Ensure directory exists
	err = os.MkdirAll(filepath.Dir(settingsPath), 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Write the file
	err = os.WriteFile(settingsPath, updated, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	if shouldSkip {
		fmt.Println("✅ Enabled SKIP_HOST_UPDATE. Discord will not check for updates.")
	} else {
		fmt.Println("❌ Disabled SKIP_HOST_UPDATE. Discord will check for updates again.")
	}
}
