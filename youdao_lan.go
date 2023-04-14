package translator_engine

var youdaoInLanMap map[string]string
var youdaoOutLanMap map[string]string

func (eng *YoudaoTransEngine) LanCodeIn(code string) string {
	v, ok := youdaoInLanMap[code]
	if ok {
		return v
	} else {
		return code
	}
}

func (eng *YoudaoTransEngine) LanCodeOut(code string) string {
	v, ok := youdaoOutLanMap[code]
	if ok {
		return v
	} else {
		return code
	}
}

func init() {
	youdaoInLanMap = map[string]string{
		ZH:    "zh-CHS",
		ZH_TW: "zh-CHT",
	}
	aliOutLanMap = map[string]string{}
	switchLanCodeMap(aliInLanMap, youdaoOutLanMap)
}
