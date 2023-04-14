package translator_engine

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type baiduAuthRequest struct {
	req         *http.Request
	accessToken string
	expireAt    int64
}

type baiduRequest struct {
	req *http.Request
}

func (r *baiduAuthRequest) build(ctx context.Context, method, url string, headers map[string]string, reqBody any) (err error) {
	if reqBody == nil {
		r.req, err = http.NewRequestWithContext(ctx, method, url, nil)
		return
	}

	if r.req, err = http.NewRequestWithContext(
		ctx,
		method,
		url,
		strings.NewReader(reqBody.(string)),
	); err != nil {
		return
	}
	for k, v := range headers {
		r.req.Header.Set(k, v)
	}
	return
}

func (r *baiduRequest) build(ctx context.Context, method, url string, headers map[string]string, reqBody any) (err error) {
	if reqBody == nil {
		r.req, err = http.NewRequestWithContext(ctx, method, url, nil)
	}
	var reqBytes []byte
	reqBytes, err = json.Marshal(reqBody)
	if err != nil {
		return
	}

	if r.req, err = http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBytes),
	); err != nil {
		return
	}
	for k, v := range headers {
		r.req.Header.Set(k, v)
	}
	return
}

func (r *baiduAuthRequest) SetHeaders(headers map[string]string) {
	for k, v := range headers {
		r.req.Header.Set(k, v)
	}
}

func (r *baiduRequest) SetHeaders(headers map[string]string) {
	for k, v := range headers {
		r.req.Header.Set(k, v)
	}
}
