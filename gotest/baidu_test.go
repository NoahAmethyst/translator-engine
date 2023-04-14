package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"os"
	"testing"
)

func TestBaiduTrans(t *testing.T) {
	apiKey, secretKey := getBaiduCfg()
	cli := translator_engine.NewBaiduTransEngine(apiKey, secretKey)
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
