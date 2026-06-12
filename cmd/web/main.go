package main

import (
	"log"
	"os"

	"stone-ocean-web/internal/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := router.New()

	log.Printf("stoneOceanWeb is running at http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
