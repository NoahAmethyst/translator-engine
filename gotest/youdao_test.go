package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"os"
	"testing"
)

func TestYoudaoTrans(t *testing.T) {
	appKey, secretKey := getYoudaoCfg()
	cli := translator_engine.NewYoudaoTranslatorCli(appKey, secretKey)
	result, err := cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("youdao translate result:%+v", result)
	to = translator_engine.EN
	src = "这是一段测试文本。是中文的"
	result, err = cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("youdao translate result:%+v", result)

}

func TestYoudaoTransWithScene(t *testing.T) {
	appKey, secretKey := getYoudaoCfg()
	cli := translator_engine.NewYoudaoTranslatorCli(appKey, secretKey)
	result, err := cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("youdao translate result:%+v", result)
}

func getYoudaoCfg() (string, string) {
	return os.Getenv("YD_APP_KEY"), os.Getenv("TD_SECRET_KEY")
}
