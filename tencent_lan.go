package translator_engine

var tencentInLanMap map[string]string
var tencentOutLanMap map[string]string

func (eng *TencentTransEngine) LanCodeIn(code string) string {
	return code
}

func (eng *TencentTransEngine) LanCodeOut(code string) string {
	return code
}

func init() {
	tencentInLanMap = map[string]string{}
	tencentOutLanMap = map[string]string{}
	switchLanCodeMap(tencentInLanMap, tencentOutLanMap)
}
