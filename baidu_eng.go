package translator_engine

import (
	"net/http"
	"time"
)

type BaiduTransEngine struct {
	client      *http.Client
	requester   *baiduRequest
	authRequest *baiduAuthRequest
	key         string
	secret      string
}

func (eng *BaiduTransEngine) SetTimeOut(timeout time.Duration) *BaiduTransEngine {
	eng.client.Timeout = timeout
	return eng
}

func (eng *BaiduTransEngine) SetTransport(transport http.RoundTripper) *BaiduTransEngine {
	eng.client.Transport = transport
	return eng
}

func (eng *BaiduTransEngine) SetJar(jar http.CookieJar) *BaiduTransEngine {
	eng.client.Jar = jar
	return eng
}

func (f *factory) BuildBaiduEng(apiKey, secretKey string) *BaiduTransEngine {
	return &BaiduTransEngine{
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       60 * time.Second,
		},
		requester:   &baiduRequest{},
		authRequest: &baiduAuthRequest{},
		key:         apiKey,
		secret:      secretKey,
	}
}
