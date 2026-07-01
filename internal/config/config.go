package config

import (
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Email    EmailConfig
}

type AppConfig struct {
	Port string
}

type DatabaseConfig struct {
	Enabled      bool
	Driver       string
	Host         string
	Port         string
	User         string
	Password     string
	Name         string
	Charset      string
	ParseTime    bool
	Location     string
	AutoMigrate  bool
	SeedPlans    bool
	SeedDemoData bool
}

type EmailConfig struct {
	Enabled     bool
	Host        string
	Port        string
	Username    string
	Password    string
	FromName    string
	FromAddress string
	ReplyTo     string
}

func Load() Config {
	_ = godotenv.Load()
	email := loadEmailConfig()

	return Config{
		App: AppConfig{
			Port: env("APP_PORT", env("PORT", "8080")),
		},
		Database: DatabaseConfig{
			Enabled:      envBool("DB_ENABLED", false),
			Driver:       env("DB_DRIVER", "mysql"),
			Host:         env("DB_HOST", "127.0.0.1"),
			Port:         env("DB_PORT", "3306"),
			User:         env("DB_USER", "root"),
			Password:     env("DB_PASSWORD", ""),
			Name:         env("DB_NAME", "recoverease"),
			Charset:      env("DB_CHARSET", "utf8mb4"),
			ParseTime:    envBool("DB_PARSE_TIME", true),
			Location:     env("DB_LOC", "Local"),
			AutoMigrate:  envBool("DB_AUTO_MIGRATE", true),
			SeedPlans:    envBool("DB_SEED_PLANS", true),
			SeedDemoData: envBool("DB_SEED_DEMO_DATA", false),
		},
		Email: email,
	}
}

func loadEmailConfig() EmailConfig {
	username := env("EMAIL_USER", env("SMTP_USER", env("MAIL_USER", env("USER", ""))))
	cfg := EmailConfig{
		Host:        env("EMAIL_HOST", env("SMTP_HOST", env("MAIL_HOST", env("HOST", "")))),
		Port:        env("EMAIL_PORT", env("SMTP_PORT", env("MAIL_PORT", env("PORT", "587")))),
		Username:    username,
		Password:    env("EMAIL_PASSWORD", env("SMTP_PASSWORD", env("MAIL_PASSWORD", env("PASS", "")))),
		FromName:    env("EMAIL_FROM_NAME", "RecoverEase"),
		FromAddress: env("EMAIL_FROM", env("EMAIL_ALIAS", env("ALIAS", username))),
		ReplyTo:     env("EMAIL_REPLY_TO", env("REPLY_TO", "")),
	}
	configured := cfg.Host != "" && cfg.Port != "" && cfg.Username != "" && cfg.Password != "" && cfg.FromAddress != ""
	cfg.Enabled = envBool("EMAIL_ENABLED", configured)
	return cfg
}

func (c DatabaseConfig) MySQLDSN() string {
	loc, err := time.LoadLocation(c.Location)
	if err != nil {
		loc = time.Local
	}

	mysqlConfig := mysql.Config{
		User:                 c.User,
		Passwd:               c.Password,
		Net:                  "tcp",
		Addr:                 c.Host + ":" + c.Port,
		DBName:               c.Name,
		ParseTime:            c.ParseTime,
		Loc:                  loc,
		AllowNativePasswords: true,
		Params: map[string]string{
			"charset": c.Charset,
		},
	}

	return mysqlConfig.FormatDSN()
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func envBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return parsed
}
