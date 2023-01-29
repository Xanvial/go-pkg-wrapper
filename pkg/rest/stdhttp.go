package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type stdHttp struct {
	client *http.Client
}

func NewStdHttp(cfg Config) Connection {
	return &stdHttp{
		client: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

func (sh *stdHttp) Get(ctx context.Context, param RestParam) (int, []byte, error) {
	return sh.connect(ctx, http.MethodGet, param)
}

func (sh *stdHttp) Put(ctx context.Context, param RestParam) (int, []byte, error) {
	return sh.connect(ctx, http.MethodPut, param)
}

func (sh *stdHttp) Post(ctx context.Context, param RestParam) (int, []byte, error) {
	return sh.connect(ctx, http.MethodPost, param)
}

func (sh *stdHttp) Delete(ctx context.Context, param RestParam) (int, []byte, error) {
	return sh.connect(ctx, http.MethodDelete, param)
}

func (sh *stdHttp) connect(ctx context.Context, method string, param RestParam) (int, []byte, error) {
	var (
		reqBody io.Reader
	)

	if param.JsonBodyData != nil {
		jsonByte, err := json.Marshal(param.JsonBodyData)
		if err != nil {
			return 0, nil, fmt.Errorf("invalid json format: %w", err)
		}

		reqBody = bytes.NewBuffer(jsonByte)
		// override content-type
		if param.Header == nil {
			param.Header = make(map[string][]string)
		}
		param.Header["Content-Type"] = []string{"application/json"}
	}

	req, err := http.NewRequestWithContext(ctx, method, param.Url, reqBody)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header = param.Header

	if len(param.QueryParam) > 0 {
		tmpQuery := req.URL.Query()
		for k, v := range param.QueryParam {
			tmpQuery.Add(k, v)
		}
		req.URL.RawQuery = tmpQuery.Encode()
	}

	resp, err := sh.client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to execute http request: %w", err)
	}

	statusCode := resp.StatusCode

	if resp.Body == nil || resp.Body == http.NoBody {
		return statusCode, nil, nil
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return statusCode, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return statusCode, content, nil
}
