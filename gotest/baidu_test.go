package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"os"
	"testing"
)

func TestBaiduTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"

	apiKey, secretKey := getBaiduCfg()
	cli := translator_engine.EngFactory.BuildBaiduEng(apiKey, secretKey)
	result, err := cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("baidu translate result:%+v", result)

	to = translator_engine.ZH
	src = "this is a sentence for testing.Which is in English"
	result, err = cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("baidu translate result:%+v", result)
}

func getBaiduCfg() (string, string) {
	return os.Getenv("BAIDU_API_KEY"), os.Getenv("BAIDU_SECRET_KEY")
}
