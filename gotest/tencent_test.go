package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"os"
	"testing"
)

func TestTencentTrans(t *testing.T) {
	secretId, secretKey := getTencentCfg()
	cli, _ := translator_engine.NewTencentTransEngine(secretId, secretKey)
	result, err := cli.TransText(src, from, to)
	if err != nil {
		panic(err)
	}

	t.Logf("%+v", result)

}

func getTencentCfg() (string, string) {
	return os.Getenv("TC_SECRET_ID"), os.Getenv("TC_SECRET_KEY")
}
