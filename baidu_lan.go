package translator_engine

var baiduInLanMap map[string]string
var baiduOutLanMap map[string]string

func (eng *BaiduTransEngine) LanCodeIn(code string) string {
	v, ok := baiduInLanMap[code]
	if ok {
		return v
	} else {
		return code
	}
}

func (eng *BaiduTransEngine) LanCodeOut(code string) string {
	v, ok := baiduOutLanMap[code]
	if ok {
		return v
	} else {
		return code
	}
}

func init() {
	baiduInLanMap = map[string]string{
		ZH:    "zh",
		ZH_TW: "cht",
		JA:    "jp",
		KO:    "kor",
		FR:    "fra",
		ES:    "spa",
		VI:    "vie",
		MS:    "may",
	}
	baiduOutLanMap = map[string]string{}
	switchLanCodeMap(baiduInLanMap, baiduOutLanMap)
}
