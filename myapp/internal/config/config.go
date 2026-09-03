package config

import "os"

type Config struct {
	Port string
}

// Load reads config from environment variables, falling back to
// sane defaults. Swap in `viper` or `envconfig` here if the app grows.
func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return Config{Port: port}
}
