package database

type Database interface {
	QueryRow(query string, result interface{}, params ...interface{}) error
	Query(query string, result interface{}, params ...interface{}) error
	Exec(query string, params ...interface{}) error
	ExecReturn(query string, params ...interface{}) (int, error)
}
