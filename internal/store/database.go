package store

import (
	"fmt"

	"stone-ocean-web/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDatabase(cfg config.DatabaseConfig) (*gorm.DB, error) {
	switch cfg.Driver {
	case "mysql":
		return gorm.Open(mysql.Open(cfg.MySQLDSN()), &gorm.Config{})
	default:
		return nil, fmt.Errorf("%w: unsupported database driver %q", ErrInvalidInput, cfg.Driver)
	}
}
