package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"testing"
)

func TestIBaiduTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
	apiKey, secretKey := getBaiduCfg()
	baiduEng := translator_engine.EngFactory.BuildBaiduEng(apiKey, secretKey)
	baiduResult, err := translator_engine.TransText(src, from, to, baiduEng)
	if err != nil {
		panic(err)
	}
	t.Logf("baidu trans result:%+v", baiduResult)
}

func TestIYoudaoTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
	appKey, secretKey := getYoudaoCfg()
	youdaoEng := translator_engine.EngFactory.BuildYoudaoEng(appKey, secretKey)
	youdaoResult, err := translator_engine.TransText(src, from, to, youdaoEng, translator_engine.Finance)
	if err != nil {
		panic(err)
	}
	t.Logf("youdao trans result:%+v", youdaoResult)
}

func TestITencentTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
	secretId, secretKey := getTencentCfg()
	tencentEng, err := translator_engine.EngFactory.BuildTencentEng(secretId, secretKey)
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
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
	accessId, accessKey := getAliCfg()
	aliEng, err := translator_engine.EngFactory.BuildAliEngine(accessId, accessKey)
	if err != nil {
		panic(err)
	}

	aliResult, err := translator_engine.TransText(src, from, to, aliEng, translator_engine.Finance)
	if err != nil {
		panic(err)
	}
	t.Logf("ali trans result:%+v", aliResult)
}

func TestIVolcTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"
	accessId, accessKey := getVolcConfig()
	volcEng := translator_engine.EngFactory.BuildVolcEngine(accessId, accessKey)

	volcResult, err := translator_engine.TransText(src, from, to, volcEng)
	if err != nil {
		panic(err)
	}
	t.Logf("ali trans result:%+v", volcResult)
}
