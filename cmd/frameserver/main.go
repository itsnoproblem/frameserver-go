package main

import (
	"fmt"
	"github.com/itsnoproblem/frameserver-go/tile"
	"log"
	"net/http"
	"os"

	"github.com/itsnoproblem/frameserver-go/greeting"
	internalhttp "github.com/itsnoproblem/frameserver-go/http"
	"github.com/itsnoproblem/frameserver-go/providers/hubble"
)

type AppConfig struct {
	ListenPort string // "8080"
	AppURL     string // "http://127.0.0.1:8080"
	HubURL     string // "https://nemes.farcaster.xyz:2281"
	StaticDir  string
}

func main() {
	config := mustLoadConfig()

	webClient := internalhttp.NewClient()
	validator := hubble.NewProvider(webClient, config.HubURL)

	outputDir := fmt.Sprintf("%s/tiles", config.StaticDir)
	fontsDir := fmt.Sprintf("%s/fonts", config.StaticDir)
	tileMaker, err := tile.NewTileMaker(config.AppURL+"/static/tiles", outputDir, fontsDir)
	if err != nil {
		log.Fatal(err)
	}

	greetingService := greeting.NewService(validator, tileMaker, config.StaticDir, config.AppURL)
	greeetingTransport := greeting.NewTransporter(internalhttp.MakeHandler, greetingService)

	router := internalhttp.NewRouter(
		greeetingTransport,
	)
	router.StaticFS("static", http.Dir(config.StaticDir))
	if err := router.Run(":" + config.ListenPort); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func mustLoadConfig() AppConfig {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("PORT environment variable is required")
	}

	appURL, exists := os.LookupEnv("APP_URL")
	if !exists {
		log.Fatal("APP_URL environment variable is required")
	}

	hubURL, exists := os.LookupEnv("HUB_URL")
	if !exists {
		log.Fatal("HUB_URL environment variable is required")
	}

	staticDir, exists := os.LookupEnv("STATIC_DIR")
	if !exists {
		log.Fatal("STATIC_DIR environment variable is required")
	}

	return AppConfig{
		ListenPort: port,
		AppURL:     appURL,
		HubURL:     hubURL,
		StaticDir:  staticDir,
	}
}