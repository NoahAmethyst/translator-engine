package translator_engine

import (
	"context"
	"errors"
)

type factory struct {
}

var EngFactory = factory{}

type IRequest interface {
	build(ctx context.Context, method, url string, headers map[string]string, requestBody any) error
	SetHeaders(map[string]string)
}

type TransResult struct {
	From string
	To   string
	Src  string
	Dst  string
}

type ITransEngine interface {
	LanCodeIn(lanCode string) string
	LanCodeOut(lanCode string) string
	TransTextDirect(src, from, to string, scene ...Scene) (*TransResult, error)
	TransText(src, from, to string, scene ...Scene) (*TransResult, error)
}

func TransText(src, from, to string, eng ITransEngine, scene ...Scene) (*TransResult, error) {

	if eng == nil {
		return nil, errors.New("unsupported engine")
	}

	return eng.TransText(src, from, to, scene...)
}
