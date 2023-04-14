package translator_engine

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

func (eng *TencentTransEngine) TransText(src, from, to string, _ ...Scene) (data *TransResult, err error) {
	from = eng.LanCodeIn(from)
	to = eng.LanCodeIn(to)
	return eng.TransTextDirect(src, from, to)
}

func (eng *TencentTransEngine) TransTextDirect(src, from, to string, _ ...Scene) (data *TransResult, err error) {

	req := tmt.NewTextTranslateRequest()
	req.SourceText = common.StringPtr(src)
	req.Source = common.StringPtr(from)
	req.Target = common.StringPtr(to)
	req.ProjectId = common.Int64Ptr(0)

	response, err := eng.client.TextTranslate(req)

	if err != nil {
		return
	}

	data = &TransResult{
		From: eng.LanCodeIn(*response.Response.Source),
		To:   eng.LanCodeOut(*response.Response.Target),
		Src:  src,
		Dst:  *response.Response.TargetText,
	}
	return
}
