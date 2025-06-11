package cmd

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"log"
	"os"
	"ya41-56/internal/shared/logger"
)

type Config struct {
	Address     string            `env:"ADDRESS" envDefault:"localhost:8080"`
	ModeLogger  logger.ModeLogger `env:"LOG_MODE" envDefault:"dev"`
	DatabaseDSN string            `env:"DATABASE_DSN" envDefault:""`
	CorsOrigins []string          `env:"CORS_ORIGINS" envDefault:"http://localhost:3000"`
}

func ParseFlags() Config {
	_ = godotenv.Load()
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Failed to parse env: %v\n", err)
		if err != nil {
			log.Println(err)
		}
		os.Exit(1)
	}

	pflag.StringVarP(&cfg.Address, "address", "a", cfg.Address, "HTTP server address")
	pflag.StringVarP(&cfg.DatabaseDSN, "dsn", "d", cfg.DatabaseDSN, "PostgresSQL DSN")

	pflag.Parse()

	if len(pflag.Args()) > 0 {
		_, err := fmt.Fprintf(os.Stderr, "Unknown flags: %v\n", pflag.Args())
		if err != nil {
			log.Println(err)
		}
		os.Exit(1)
	}

	return cfg
}
