package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const defaultPort = 8080

type Config struct {
	Environment string
	Port        int
}

func Load() (Config, error) {
	// Load .env if present, but do not fail when the file is absent.
	_ = loadDotEnv(".env")

	cfg := Config{
		Environment: getEnv("APP_ENV", "development"),
		Port:        defaultPort,
	}

	if portStr := os.Getenv("PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return Config{}, fmt.Errorf("invalid PORT value %q: %w", portStr, err)
		}
		if port <= 0 {
			return Config{}, fmt.Errorf("invalid PORT value %q: must be > 0", portStr)
		}
		cfg.Port = port
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func loadDotEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
		if key == "" {
			continue
		}
		if _, exists := os.LookupEnv(key); !exists {
			_ = os.Setenv(key, value)
		}
	}

	return scanner.Err()
}
