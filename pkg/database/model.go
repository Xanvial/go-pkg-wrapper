package database

type DBType int

const (
	DBTypeUnknown = iota
	DBTypePostgresql
	DBTypeMySql
)

type Config struct {
	DatabaseType DBType
	Host         string
	Username     string
	Password     string
	DatabaseName string
	SSLMode      string // default: disable, only for postgres
	Port         int    // default: 5432 for postgres, 3306 for mysql
	TimeZone     string // default: Local
}
