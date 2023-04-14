package translator_engine

import (
	alimt_cli "github.com/alibabacloud-go/alimt-20181012/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

const (
	aliApiEndpoint = "mt.cn-hangzhou.aliyuncs.com"
)

type AliTransEngine struct {
	client *alimt_cli.Client
	key    string
	secret string
}

func (f *factory) BuildAliEngine(accessId, accessSecret string) (*AliTransEngine, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessId),
		AccessKeySecret: tea.String(accessSecret),
		Endpoint:        tea.String(aliApiEndpoint),
	}
	cli, err := alimt_cli.NewClient(config)

	return &AliTransEngine{
		client: cli,
		key:    accessId,
		secret: accessSecret,
	}, err
}
