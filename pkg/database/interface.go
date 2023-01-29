package database

type Database interface {
	QueryRow(query string, result any, params ...any) error
	Query(query string, result any, params ...any) error
	Exec(query string, params ...any) error
	ExecReturn(query string, params ...any) (int, error)
}
