package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormDB struct {
	db *gorm.DB
}

func NewGorm(cfg Config) (Database, error) {
	var (
		cli *gorm.DB
		err error
	)
	switch cfg.DatabaseType {
	case DBTypePostgresql:
		sslmode := cfg.SSLMode
		if sslmode == "" {
			sslmode = "disable"
		}
		port := cfg.Port
		if port <= 0 {
			port = 5432
		}
		var timezone string
		if cfg.TimeZone != "" {
			timezone = " TimeZone=" + cfg.TimeZone
		}
		cli, err = gorm.Open(
			postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s%s",
				cfg.Host,
				cfg.Username,
				cfg.Password,
				cfg.DatabaseName,
				port,
				sslmode,
				timezone,
			)),
			&gorm.Config{},
		)
	case DBTypeMySql:
		port := cfg.Port
		if port <= 0 {
			port = 3306
		}
		timezone := cfg.TimeZone
		if timezone == "" {
			timezone = "Local"
		}
		cli, err = gorm.Open(
			mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=%s",
				cfg.Username,
				cfg.Password,
				cfg.Host,
				port,
				cfg.DatabaseName,
				timezone,
			)),
			&gorm.Config{},
		)
	default:
		return nil, errors.New("invalid database type")
	}

	if err != nil {
		return nil, err
	}

	return &gormDB{
		db: cli,
	}, nil
}

func (gd *gormDB) QueryRow(query string, result any, params ...any) error {
	dbTx := gd.db.Raw(query, params...).Scan(result)
	if dbTx.Error != nil {
		return dbTx.Error
	}
	return nil
}

func (gd *gormDB) Query(query string, result any, params ...any) error {
	dbTx := gd.db.Raw(query, params...).Scan(result)
	if dbTx.Error != nil {
		return dbTx.Error
	}
	return nil
}

func (gd *gormDB) Exec(query string, params ...any) error {
	dbTx := gd.db.Exec(query, params...)
	if dbTx.Error != nil {
		return dbTx.Error
	}
	return nil
}

func (gd *gormDB) ExecReturn(query string, params ...any) (int, error) {
	var returning int
	dbTx := gd.db.Raw(query, params...).Scan(&returning)
	if dbTx.Error != nil {
		return 0, dbTx.Error
	}
	return returning, nil
}
