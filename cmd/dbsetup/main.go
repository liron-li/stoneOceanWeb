package main

import (
	"context"
	"log"

	"stone-ocean-web/internal/config"
	"stone-ocean-web/internal/store"
)

func main() {
	cfg := config.Load()

	db, err := store.OpenDatabase(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.AutoMigrate(db); err != nil {
		log.Fatal(err)
	}
	if err := store.SeedDefaultPlans(context.Background(), db); err != nil {
		log.Fatal(err)
	}
	if err := store.SeedDemoData(context.Background(), store.New(db)); err != nil {
		log.Fatal(err)
	}

	log.Printf("database tables and demo data are ready using %s", cfg.Database.Driver)
}
