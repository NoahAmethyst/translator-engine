package translator_engine

var aliInLanMap map[string]string
var aliOutLanMap map[string]string

func (eng *AliTransEngine) LanCodeIn(code string) string {
	return code
}

func (eng *AliTransEngine) LanCodeOut(code string) string {
	return code
}

func init() {
	aliInLanMap = map[string]string{}
	aliOutLanMap = map[string]string{}
	switchLanCodeMap(aliInLanMap, aliOutLanMap)
}
