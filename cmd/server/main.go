package main

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/typetrait/lit/internal"
)

const (
	EnvFile string = ".env"
)

func main() {
	if fileExists(EnvFile) {
		_ = godotenv.Load(EnvFile)
	}

	env := internal.Environment{
		IsDebugEnabled: strings.ToLower(os.Getenv("DEBUG")) == "true",
		DBHost:         os.Getenv("DB_HOST"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBPort:         os.Getenv("DB_PORT"),
		S3Bucket:       os.Getenv("S3_BUCKET"),
		LocalstackHost: os.Getenv("LOCALSTACK_HOST"),
	}

	app := internal.NewApp(&env)
	app.Start(":1323")
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
