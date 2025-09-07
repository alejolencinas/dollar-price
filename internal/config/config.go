package config

import (
	"log"
	"os"
)

type Config struct {
	Port   string
	Env    string
	BnaUrl string
}

func Load() *Config {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // default
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	bnaUrl := os.Getenv("BNA_URL")
	if bnaUrl == "" {
		bnaUrl = "https://www.bna.com.ar/Personas"
	}

	log.Printf("Loaded config: PORT=%s, ENV=%s, BNA_URL=%s\n", port, env, bnaUrl)

	return &Config{
		Port:   port,
		Env:    env,
		BnaUrl: bnaUrl,
	}
}
