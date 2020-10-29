package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
)

// Load loads config and secret
func Load() {
	// Get current workdir.
	WorkDir = os.Getenv("INVENTORY_WORKDIR")

	// Path to config photos.
	confPath := filepath.Join(WorkDir, "conf", "config.ini")

	// Create default config if needed.
	err := createDefaultConfig(confPath)
	if err != nil {
		log.Fatalf("Could not create default config: %v\n", err)
	}

	// Load config file.
	cfg, err := ini.Load(confPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load server config.
	err = cfg.Section("server").MapTo(&Server)
	if err != nil {
		log.Fatal(err)
	}

	// Load database config.
	err = cfg.Section("database").MapTo(&Database)
	if err != nil {
		log.Fatal(err)
	}

	// Load logs section.
	err = cfg.Section("logs").MapTo(&Logs)
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.Section("microservices").MapTo(&Microservices)
}

// createDefaultConfig creates default settings file at `path`
// if settings file does not yet exist.
func createDefaultConfig(path string) error {
	// Check if config photos already exists.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("No config found. Creating default config at \"%s\"\n", path)

		// Create new config file.
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()

		// Get default config file to be written.
		data, err := Asset("config.ini")
		if err != nil {
			return err
		}

		// Write config file.
		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}
