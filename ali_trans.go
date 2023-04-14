package translator_engine

import (
	alimt_cli "github.com/alibabacloud-go/alimt-20181012/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

var aliScene map[Scene]string

func (eng *AliTransEngine) TransText(src, from, to string, scene ...Scene) (data *TransResult, err error) {
	from = eng.LanCodeIn(from)
	to = eng.LanCodeIn(to)
	return eng.TransTextDirect(src, from, to, scene...)
}

func (eng *AliTransEngine) TransTextDirect(src, from, to string, scene ...Scene) (data *TransResult, err error) {
	var _scene string
	if len(scene) > 0 {
		_scene = aliScene[scene[0]]
	}

	if len(_scene) > 0 {
		req := &alimt_cli.TranslateGeneralRequest{
			FormatType:     tea.String("text"),
			SourceLanguage: tea.String(from),
			SourceText:     tea.String(src),
			TargetLanguage: tea.String(to),
		}

		resp, _err := eng.client.TranslateGeneral(req)
		if err != nil {
			err = _err
			return
		}
		data = &TransResult{
			From: from,
			To:   to,
			Src:  src,
			Dst:  *resp.Body.Data.Translated,
		}
	} else {
		req := &alimt_cli.TranslateRequest{
			FormatType:     tea.String("text"),
			Scene:          tea.String(_scene),
			SourceLanguage: tea.String(from),
			SourceText:     tea.String(src),
			TargetLanguage: tea.String(to),
		}

		resp, _err := eng.client.Translate(req)
		if err != nil {
			err = _err
			return
		}
		data = &TransResult{
			From: eng.LanCodeOut(from),
			To:   eng.LanCodeOut(to),
			Src:  src,
			Dst:  *resp.Body.Data.Translated,
		}
	}

	return
}

func init() {
	aliScene = map[Scene]string{
		Title:       "title",
		Desc:        "description",
		Communicate: "communication",
		Medicine:    "medical",
		Social:      "social",
		Finance:     "finance",
	}
}
