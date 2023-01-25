package rest

import "context"

type Connection interface {
	Get(ctx context.Context, data RestParam) (statusCode int, response []byte, err error)
	Put(ctx context.Context, data RestParam) (statusCode int, response []byte, err error)
	Post(ctx context.Context, data RestParam) (statusCode int, response []byte, err error)
	Delete(ctx context.Context, data RestParam) (statusCode int, response []byte, err error)
}
