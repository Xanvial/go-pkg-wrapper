package rest

import "time"

type Config struct {
	Timeout time.Duration
}

type RestParam struct {
	Url          string
	JsonBodyData any
	QueryParam   map[string]string
	Header       map[string][]string
}
