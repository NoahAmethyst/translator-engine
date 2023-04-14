package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"os"
	"testing"
)

func TestYoudaoTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它是中文的，将要翻译成英文"

	appKey, secretKey := getYoudaoCfg()
	cli := translator_engine.EngFactory.BuildYoudaoEng(appKey, secretKey)
	result, err := cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("youdao translate result:%+v", result)

	to = translator_engine.ZH
	src = result.Dst
	result, err = cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("youdao translate result:%+v", result)

}

func TestYoudaoTransWithScene(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"

	appKey, secretKey := getYoudaoCfg()
	cli := translator_engine.EngFactory.BuildYoudaoEng(appKey, secretKey)
	result, err := cli.TransText(src, from, to, translator_engine.Finance)
	if err != nil {
		panic(err)
	}
	t.Logf("youdao translate result:%+v", result)
}

func getYoudaoCfg() (string, string) {
	return os.Getenv("YD_APP_KEY"), os.Getenv("YD_SECRET_KEY")
}
