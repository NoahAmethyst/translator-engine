package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"os"
	"testing"
)

func TestTencentTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它的语言是中文，将要翻译为英文"

	secretId, secretKey := getTencentCfg()
	cli, _ := translator_engine.EngFactory.BuildTencentEng(secretId, secretKey)
	result, err := cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}

	t.Logf("%+v", result)

}

func getTencentCfg() (string, string) {
	return os.Getenv("TC_SECRET_ID"), os.Getenv("TC_SECRET_KEY")
}
