package gotest

import (
	translator_engine "github.com/NoahAmethyst/translator-engine"
	"os"
	"testing"
)

func TestVolcTrans(t *testing.T) {
	from := translator_engine.AUTO
	to := translator_engine.EN
	src := "这是一段用来测试的文本，它是中文的，将要翻译成英文"

	appKey, secretKey := getVolcConfig()
	cli := translator_engine.EngFactory.BuildVolcEngine(appKey, secretKey)
	result, err := cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("volc translate result:%+v", result)

	to = translator_engine.ZH
	src = result.Dst
	result, err = cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}
	t.Logf("volc translate result:%+v", result)

}

func getVolcConfig() (string, string) {
	return os.Getenv("VOLC_ACCESS_KEY"), os.Getenv("VOLC_SECRET_KEY")
}
