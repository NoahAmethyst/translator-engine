package translator_engine

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

const (
	volcApiEndPoint = "open.volcengineapi.com"
	kServiceVersion = "2020-06-01"
)

type VolcTransEngine struct {
	client *base.Client
	key    string
	secret string
}

func (f *factory) BuildVolcEngine(accessKey, secretKey string) *VolcTransEngine {
	ServiceInfo := &base.ServiceInfo{
		Timeout: 60 * time.Second,
		Host:    volcApiEndPoint,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "translate"},
	}
	ApiInfoList := map[string]*base.ApiInfo{
		"TranslateText": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TranslateText"},
				"Version": []string{kServiceVersion},
			},
		},
	}
	client := base.NewClient(ServiceInfo, ApiInfoList)
	client.SetAccessKey(accessKey)
	client.SetSecretKey(secretKey)
	return &VolcTransEngine{
		client: client,
		key:    accessKey,
		secret: secretKey,
	}
}

func (eng *VolcTransEngine) SetTimeout(timeout time.Duration) {
	eng.client.ServiceInfo.Timeout = timeout
}
