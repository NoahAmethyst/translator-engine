package gotest

import (
	"github.com/NoahAmethyst/translator-engine"
	"os"
	"testing"
)

func TestAliTrans(t *testing.T) {
	accessId, accessKey := getAliCfg()
	cli, err := translator_engine.NewAliTransEngine(accessId, accessKey)
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
	id, secret := getAliCfg()
	cli, err := translator_engine.NewAliTransEngine(id, secret)
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
