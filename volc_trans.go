package translator_engine

import (
	"encoding/json"
	"errors"
)

type volcTransReq struct {
	SourceLanguage string   `json:"SourceLanguage"`
	TargetLanguage string   `json:"TargetLanguage"`
	TextList       []string `json:"TextList"`
	Options        struct {
		Category string `json:"Category"`
	} `json:"Options"`
}

type volcTransResp struct {
	TranslationList []struct {
		Translation            string `json:"Translation"`
		DetectedSourceLanguage string `json:"DetectedSourceLanguage"`
	} `json:"TranslationList"`
	ResponseMetadata struct {
		RequestId string `json:"RequestId"`
		Action    string `json:"Action"`
		Version   string `json:"Version"`
		Service   string `json:"Service"`
		Region    string `json:"Region"`
		Error     *struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error"`
	} `json:"ResponseMetadata"`
}

func (eng *VolcTransEngine) TransText(src, from, to string, _ ...Scene) (data *TransResult, err error) {
	from = eng.LanCodeIn(from)
	to = eng.LanCodeIn(to)
	return eng.TransTextDirect(src, from, to)
}

func (eng *VolcTransEngine) TransTextDirect(src, from, to string, _ ...Scene) (data *TransResult, err error) {

	req := volcTransReq{
		TextList:       []string{src},
		TargetLanguage: to,
	}

	if from != AUTO {
		req.SourceLanguage = from
	}

	reqbytes, _err := json.Marshal(req)
	if _err != nil {
		err = _err
		return
	}

	resp, _, _err := eng.client.Json("TranslateText", nil, string(reqbytes))

	if _err != nil {
		err = _err
		return
	}

	var respStruct volcTransResp

	_err = json.Unmarshal(resp, &respStruct)
	if _err != nil {
		err = _err
		return
	}

	if respStruct.ResponseMetadata.Error != nil {
		err = errors.New(respStruct.ResponseMetadata.Error.Message)
		return

	}

	data = &TransResult{
		From: eng.LanCodeIn(respStruct.TranslationList[0].DetectedSourceLanguage),
		To:   eng.LanCodeOut(to),
		Src:  src,
		Dst:  respStruct.TranslationList[0].Translation,
	}
	return
}
