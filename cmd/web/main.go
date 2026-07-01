package main

import (
	"context"
	"log"

	"stone-ocean-web/internal/config"
	"stone-ocean-web/internal/events"
	"stone-ocean-web/internal/mailer"
	"stone-ocean-web/internal/router"
	"stone-ocean-web/internal/store"
)

func main() {
	cfg := config.Load()
	var appStore *store.Store

	if cfg.Database.Enabled {
		db, err := store.OpenDatabase(cfg.Database)
		if err != nil {
			log.Fatal(err)
		}
		if cfg.Database.AutoMigrate {
			if err := store.AutoMigrate(db); err != nil {
				log.Fatal(err)
			}
		}
		if cfg.Database.SeedPlans {
			if err := store.SeedDefaultPlans(context.Background(), db); err != nil {
				log.Fatal(err)
			}
		}
		appStore = store.New(db)
		if cfg.Database.SeedDemoData {
			if err := store.SeedDemoData(context.Background(), appStore); err != nil {
				log.Fatal(err)
			}
		}
		log.Printf("database connected using %s", cfg.Database.Driver)
	}

	eventBus := events.NewBus(log.Default())
	if cfg.Email.Enabled {
		smtpMailer := mailer.NewSMTPMailer(cfg.Email)
		eventBus.AddPaymentPaidListener(events.NewLicenseEmailListener(smtpMailer))
		eventBus.AddLicenseRecoveryCodeListener(events.NewLicenseRecoveryCodeEmailListener(smtpMailer))
		log.Printf("license email listeners enabled for %s:%s", cfg.Email.Host, cfg.Email.Port)
	} else {
		log.Printf("license email listeners disabled; set EMAIL_* SMTP variables to enable delivery")
	}

	r := router.NewWithEvents(appStore, eventBus)

	log.Printf("stoneOceanWeb is running at http://localhost:%s", cfg.App.Port)
	if err := r.Run(":" + cfg.App.Port); err != nil {
		log.Fatal(err)
	}
}
