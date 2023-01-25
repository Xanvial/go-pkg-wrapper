package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

type fastHttp struct {
	timeout time.Duration
}

func NewFastHttp(cfg Config) Connection {
	return &fastHttp{
		timeout: cfg.Timeout,
	}
}

func (fh *fastHttp) Get(ctx context.Context, param RestParam) (int, []byte, error) {
	return fh.connect(ctx, http.MethodGet, param)
}

func (fh *fastHttp) Put(ctx context.Context, param RestParam) (int, []byte, error) {
	return fh.connect(ctx, http.MethodPut, param)
}

func (fh *fastHttp) Post(ctx context.Context, param RestParam) (int, []byte, error) {
	return fh.connect(ctx, http.MethodPost, param)
}

func (fh *fastHttp) Delete(ctx context.Context, param RestParam) (int, []byte, error) {
	return fh.connect(ctx, http.MethodDelete, param)
}

func (fh *fastHttp) connect(ctx context.Context, method string, param RestParam) (int, []byte, error) {
	var (
		req  = fasthttp.AcquireRequest()
		resp = fasthttp.AcquireResponse()
	)
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(param.Url)
	req.Header.SetMethod(method)

	if param.JsonBodyData != nil {
		jsonByte, err := json.Marshal(param.JsonBodyData)
		if err != nil {
			return 0, nil, errors.Wrap(err, "invalid json format")
		}
		req.SetBody(jsonByte)
	}

	for k, v := range param.QueryParam {
		req.URI().QueryArgs().Add(k, v)
	}

	for key, headerList := range param.Header {
		for _, v := range headerList {
			req.Header.Add(key, v)
		}
	}

	err := fasthttp.DoTimeout(req, resp, fh.timeout)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}
