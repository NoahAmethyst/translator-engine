package translator_engine

import "sync"

var once sync.Once

const (
	// ZH 中文
	ZH = "zh"
	// ZH_TW 中文繁体
	ZH_TW = "zh-TW"
	// EN 英文
	EN = "en"
	// JA 日文
	JA = "ja"
	// KO 韩文
	KO = "ko"
	// FR 法文
	FR = "fr"
	// ES 西班牙文
	ES = "es"
	// PT 葡萄牙文
	PT = "pt"
	// IT 意大利文
	IT = "it"
	// RU 俄文
	RU = "ru"
	// DE 德语
	DE = "de"
	// TR 土耳其语
	TR = "tr"
	// VI 越南语
	VI = "vi"
	// ID 印尼
	ID = "id"
	// TH 泰语
	TH = "th"
	// MS 马来语
	MS = "ms"
	// AUTO 自动识别，仅对from有效
	AUTO = "auto"
)

func switchLanCodeMap(in, out map[string]string) {
	once.Do(func() {
		if out == nil {
			out = map[string]string{}
		}
	})
	for k, v := range in {
		out[v] = k
	}

}
