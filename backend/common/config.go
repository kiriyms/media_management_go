package common

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ADDR string
	PORT string
	ENV  string
	KEY  string
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
	err := godotenv.Load()
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

	key, ok := os.LookupEnv("KEY")
	if !ok {
		log.Fatal("KEY environment variable missing")
	}

	onceCfg.Do(func() {
		cfg = &Config{
			ADDR: arrd,
			PORT: port,
			ENV:  env,
			KEY:  key,
		}
	})
}

func GetConfig() *Config {
	if cfg == nil {
		panic("Global config not initialized. Call MustLoadConfig() first.")
	}
	return cfg
}
