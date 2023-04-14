package translator_engine

import (
	"context"
	"reflect"
)

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

	_eng := reflect.ValueOf(eng)

	fLanProxy := _eng.MethodByName("LanCodeIn")

	lanFrom := fLanProxy.Call([]reflect.Value{reflect.ValueOf(from)})

	lanTo := fLanProxy.Call([]reflect.Value{reflect.ValueOf(to)})

	var results []reflect.Value

	if len(scene) > 0 {
		results = _eng.MethodByName("TransTextDirect").
			Call([]reflect.Value{
				reflect.ValueOf(src),
				lanFrom[0],
				lanTo[0],
				reflect.ValueOf(scene[0]),
			})
	} else {
		results = _eng.MethodByName("TransTextDirect").
			Call([]reflect.Value{
				reflect.ValueOf(src),
				lanFrom[0],
				lanTo[0],
			})
	}

	var resultsInterface []interface{}
	for _, v := range results {
		resultsInterface = append(resultsInterface, v.Interface())
	}

	var result *TransResult
	var err error
	if resultsInterface[0] != nil {
		result = resultsInterface[0].(*TransResult)
	}
	if resultsInterface[1] != nil {
		err = resultsInterface[1].(error)
	}

	return result, err
}
