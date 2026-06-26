package store

import (
	"fmt"

	"stone-ocean-web/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase(cfg config.DatabaseConfig) (*gorm.DB, error) {
	switch cfg.Driver {
	case "mysql":
		return gorm.Open(mysql.Open(cfg.MySQLDSN()), &gorm.Config{})
	case "sqlite":
		return gorm.Open(sqlite.Open(cfg.SQLitePath), &gorm.Config{})
	default:
		return nil, fmt.Errorf("%w: unsupported database driver %q", ErrInvalidInput, cfg.Driver)
	}
}
