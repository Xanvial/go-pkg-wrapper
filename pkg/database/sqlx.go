package database

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type sqlxDB struct {
	db *sqlx.DB
}

func NewSqlxDB(cfg Config) (Database, error) {
	var (
		cli *sqlx.DB
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
		cli, err = sqlx.Connect("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s%s",
			cfg.Host,
			cfg.Username,
			cfg.Password,
			cfg.DatabaseName,
			port,
			sslmode,
			timezone,
		))
	case DBTypeMySql:
		port := cfg.Port
		if port <= 0 {
			port = 3306
		}
		timezone := cfg.TimeZone
		if timezone == "" {
			timezone = "Local"
		}
		cli, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=%s",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			port,
			cfg.DatabaseName,
			timezone,
		))
	default:
		return nil, errors.New("invalid database type")
	}

	if err != nil {
		return nil, err
	}

	return &sqlxDB{
		db: cli,
	}, nil
}

func (sd *sqlxDB) QueryRow(query string, result any, params ...any) error {
	return sd.db.QueryRowx(query, params...).StructScan(result)
}

func (sd *sqlxDB) Query(query string, result any, params ...any) error {
	return sd.db.Select(result, query, params...)
}

func (sd *sqlxDB) Exec(query string, params ...any) error {
	_, err := sd.db.Exec(query, params...)
	if err != nil {
		return err
	}
	return nil
}

func (sd *sqlxDB) ExecReturn(query string, params ...any) (returning int, err error) {
	err = sd.db.QueryRow(query, params...).Scan(&returning)
	return
}
