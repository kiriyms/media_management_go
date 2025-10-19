package common

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ADDR     string
	PORT     string
	ENV      string
	USER_KEY string
	JWT_KEY  string
}

var (
	cfg     *Config
	onceCfg sync.Once
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

func MustLoadConfig() {
	err := godotenv.Load("backend/.env")
	if err != nil {
		log.Printf("Failed to load .env file: %v; will look for variables in the environment\n", err)
	}

	env, ok := os.LookupEnv("ENV")
	if !ok {
		log.Fatal("ENV environment variable missing")
	}

	arrd, ok := os.LookupEnv("ADDR")
	if !ok {
		log.Fatal("ADDR environment variable missing")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT environment variable missing")
	}

	userKey, ok := os.LookupEnv("USER_KEY")
	if !ok {
		log.Fatal("USER_KEY environment variable missing")
	}

	jwtKey, ok := os.LookupEnv("JWT_KEY")
	if !ok {
		log.Fatal("JWT_KEY environment variable missing")
	}

	onceCfg.Do(func() {
		cfg = &Config{
			ADDR:     arrd,
			PORT:     port,
			ENV:      env,
			USER_KEY: userKey,
			JWT_KEY:  jwtKey,
		}
	})
}

func GetConfig() *Config {
	if cfg == nil {
		panic("Global config not initialized. Call MustLoadConfig() first.")
	}
	return cfg
}
