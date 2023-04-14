package translator_engine

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

const (
	tencentApiEndPoint = "tmt.tencentcloudapi.com"
)

type TencentTransEngine struct {
	client *tmt.Client
	key    string
	secret string
}

func (f *factory) BuildTencentEng(secretId, secretKey string) (*TencentTransEngine, error) {
	client, err := initTencentCli(secretId, secretKey, "ap-shanghai")
	return &TencentTransEngine{
		client: client,
		key:    secretId,
		secret: secretKey,
	}, err
}

func initTencentCli(secretId, secretKey, region string) (*tmt.Client, error) {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = tencentApiEndPoint
	return tmt.NewClient(credential, region, cpf)
}

func (eng *TencentTransEngine) SetRegion(region string) error {
	client, err := initTencentCli(eng.key, eng.secret, region)
	eng.client = client
	return err
}
