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

func Load() Config {
	_ = godotenv.Load()

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
	}
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
