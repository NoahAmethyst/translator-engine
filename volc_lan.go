package translator_engine

var volcInLanMap map[string]string
var volcOutLanMap map[string]string

func (eng *VolcTransEngine) LanCodeIn(code string) string {
	return code
}

func (eng *VolcTransEngine) LanCodeOut(code string) string {
	return code
}

func init() {
	volcInLanMap = map[string]string{
		ZH: "zh-Hant",
	}
	volcOutLanMap = map[string]string{}
	switchLanCodeMap(aliInLanMap, aliOutLanMap)
}
