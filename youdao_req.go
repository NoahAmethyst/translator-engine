package translator_engine

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
)

type youdaoRequest struct {
	req *http.Request
}

func (r *youdaoRequest) build(ctx context.Context, method, apiUrl string, headers map[string]string, reqBody any) (err error) {
	if reqBody == nil {
		r.req, err = http.NewRequestWithContext(ctx, method, apiUrl, nil)
		return
	}

	if r.req, err = http.NewRequestWithContext(
		ctx,
		method,
		apiUrl,
		bytes.NewBufferString(reqBody.(url.Values).Encode()),
	); err != nil {
		return
	}
	for k, v := range headers {
		r.req.Header.Set(k, v)
	}
	return
}

func (r *youdaoRequest) SetHeaders(headers map[string]string) {
	for k, v := range headers {
		r.req.Header.Set(k, v)
	}

}
