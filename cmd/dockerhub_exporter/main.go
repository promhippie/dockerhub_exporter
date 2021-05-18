package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/promhippie/dockerhub_exporter/pkg/command"
)

func main() {
	if env := os.Getenv("DOCKERHUB_EXPORTER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
