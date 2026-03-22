package app

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	AppEnv          string
	Port            string
	DBDSN           string
	APIKey          string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

func LoadConfig() Config {
	loadDotEnv(".env")

	return Config{
		AppEnv:          getEnv("APP_ENV", "development"),
		Port:            getEnv("PORT", "3000"),
		DBDSN:           getEnv("DB_DSN", ""),
		APIKey:          getEnv("API_KEY", "adalahpokoknya"),
		ReadTimeout:     getDuration("READ_TIMEOUT", 5*time.Second),
		WriteTimeout:    getDuration("WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:     getDuration("IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout: getDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
	}
}

func (c Config) Address() string {
	if strings.HasPrefix(c.Port, ":") {
		return c.Port
	}

	return ":" + c.Port
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	return value
}

func getDuration(key string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return duration
}

func loadDotEnv(filename string) {
	path := filepath.Clean(filename)
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
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

		if _, exists := os.LookupEnv(key); exists {
			continue
		}

		_ = os.Setenv(key, value)
	}
}
