package main

import (
	"log"
	"os"

	db "github.com/Lee266/op-ai-server/models"
	"github.com/Lee266/op-ai-server/router"
)

const PORT = ":8002"

func main() {
	dbEnv := db.DbEnv{
		Host:     os.Getenv("DATABASE_HOST"),
		User:     os.Getenv("DATABASE_USER"),
		Pass:     os.Getenv("DATABASE_PASSWORD"),
		DB:       os.Getenv("DATABASE_DB"),
		Port:     os.Getenv("HOST_MACHINE_DATABASE_PORT"),
		SslMode:  os.Getenv("SSL_MODE"),
		TimeZone: os.Getenv("TIME_ZONE"),
	}
	if err := dbEnv.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	_, err := dbEnv.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	if err := router.Router().Run(PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
