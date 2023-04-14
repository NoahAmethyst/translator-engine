package translator_engine

import (
	"net/http"
	"time"
)

const (
	youdaoApi = "https://openapi.youdao.com/api"
)

type YoudaoTransEngine struct {
	client    *http.Client
	requester *youdaoRequest
	key       string
	secret    string
}

func (eng *YoudaoTransEngine) SetTimeOut(timeout time.Duration) *YoudaoTransEngine {
	eng.client.Timeout = timeout
	return eng
}

func (eng *YoudaoTransEngine) SetTransport(transport http.RoundTripper) *YoudaoTransEngine {
	eng.client.Transport = transport
	return eng
}

func (eng *YoudaoTransEngine) SetJar(jar http.CookieJar) *YoudaoTransEngine {
	eng.client.Jar = jar
	return eng
}

func (f *factory) BuildYoudaoEng(appKey, secretKey string) *YoudaoTransEngine {
	return &YoudaoTransEngine{
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       60 * time.Second,
		},
		requester: &youdaoRequest{},
		key:       appKey,
		secret:    secretKey,
	}
}
