package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"testing"
)

var from = translator_engine.AUTO
var to = translator_engine.EN
var src = "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"

func TestIBaiduTrans(t *testing.T) {
	apiKey, secretKey := getBaiduCfg()
	baiduEng := translator_engine.NewBaiduTransEngine(apiKey, secretKey)
	baiduResult, err := translator_engine.TransText(src, from, to, baiduEng)
	if err != nil {
		panic(err)
	}
	t.Logf("baidu trans result:%+v", baiduResult)
}

func TestIYoudaoTrans(t *testing.T) {
	appKey, secretKey := getYoudaoCfg()
	youdaoEng := translator_engine.NewYoudaoTranslatorCli(appKey, secretKey)
	youdaoResult, err := translator_engine.TransText(src, from, to, youdaoEng, translator_engine.Finance)
	if err != nil {
		panic(err)
	}
	t.Logf("youdao trans result:%+v", youdaoResult)
}

func TestITencentTrans(t *testing.T) {
	secretId, secretKey := getTencentCfg()
	tencentEng, err := translator_engine.NewTencentTransEngine(secretId, secretKey)
	if err != nil {
		panic(err)
	}
	tencentResult, err := translator_engine.TransText(src, from, to, tencentEng)
	if err != nil {
		panic(err)
	}
	t.Logf("tencent trans result:%+v", tencentResult)
}

func TestIAliTrans(t *testing.T) {
	accessId, accessKey := getAliCfg()
	aliEng, err := translator_engine.NewAliTransEngine(accessId, accessKey)
	if err != nil {
		panic(err)
	}

	aliResult, err := translator_engine.TransText(src, from, to, aliEng, translator_engine.Finance)
	if err != nil {
		panic(err)
	}
	t.Logf("ali trans result:%+v", aliResult)
}
