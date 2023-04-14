package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"os"
	"testing"
)

func TestAliTrans(t *testing.T) {

	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"

	accessId, accessKey := getAliCfg()

	cli, err := translator_engine.EngFactory.BuildAliEngine(accessId, accessKey)
	if err != nil {
		panic(err)
	}
	result, err := cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("ali translate result:%+v", result)
}

func TestAliTransWithScene(t *testing.T) {

	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"

	id, secret := getAliCfg()
	cli, err := translator_engine.EngFactory.BuildAliEngine(id, secret)
	if err != nil {
		panic(err)
	}
	result, err := cli.TransText(src, from, to, translator_engine.Medicine)
	if err != nil {
		panic(err)
	}
	t.Logf("ali translate result:%+v", result)
}

func getAliCfg() (string, string) {
	return os.Getenv("ALI_ACCESS_ID"), os.Getenv("ALI_ACCESS_SECRET")
}
