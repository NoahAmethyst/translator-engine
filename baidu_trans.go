package translator_engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type baiduResp struct {
	From        string      `json:"from"`
	To          string      `json:"to"`
	TransResult []TransResp `json:"trans_result"`
}

type TransResp struct {
	Dst string `json:"dst"`
	Src string `json:"src"`
}

func (eng *BaiduTransEngine) getAccessToken() (accessToken string, err error) {
	if eng.authRequest.expireAt > time.Now().Unix() {
		accessToken = eng.authRequest.accessToken
		return
	}
	ctx := context.Background()
	data := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", eng.key, eng.secret)
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	err = eng.authRequest.build(ctx, http.MethodPost, "https://aip.baidubce.com/oauth/2.0/token", headers, data)

	if err != nil {
		return
	}

	accessTokenObj := make(map[string]interface{})

	if err = doJsonRequest(eng.client, eng.authRequest.req, &accessTokenObj); err != nil {
		return
	}
	if v, ok := accessTokenObj["error_description"]; ok {
		err = errors.New(v.(string))
		return
	}
	accessToken = accessTokenObj["access_token"].(string)
	expireAt := int64(accessTokenObj["expires_in"].(float64))
	eng.authRequest.accessToken = accessToken
	eng.authRequest.expireAt = expireAt
	return
}

func (eng *BaiduTransEngine) TransText(src, from, to string, _ ...Scene) (data *TransResult, err error) {
	from = eng.LanCodeIn(from)
	to = eng.LanCodeIn(to)
	return eng.TransTextDirect(src, from, to)
}

func (eng *BaiduTransEngine) TransTextDirect(src, from, to string, _ ...Scene) (data *TransResult, err error) {
	respData := make(map[string]interface{})
	ctx := context.Background()
	req := map[string]string{
		"q":    src,
		"from": from,
		"to":   to,
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	accessToken, err := eng.getAccessToken()
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/mt/texttrans/v1?access_token=%s", accessToken)

	err = eng.requester.build(ctx, http.MethodPost, url, headers, req)

	err = doJsonRequest(eng.client, eng.requester.req, &respData)
	if err != nil {
		return
	}
	if v, ok := respData["error_description"]; ok {
		err = errors.New(v.(string))
		return
	}

	if v, ok := respData["error_msg"]; ok {
		err = errors.New(v.(string))
		return
	}

	result, ok := respData["result"]
	if !ok {
		err = errors.New(fmt.Sprintf("result colum not found in response data:%+v", respData))
		return
	}

	var respStruct baiduResp

	bytes, err := json.Marshal(result)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &respStruct)
	if err != nil {
		return
	}

	if len(respStruct.TransResult) == 0 {
		err = errors.New("empty translation")
	}
	data = &TransResult{
		From: eng.LanCodeOut(respStruct.From),
		To:   eng.LanCodeOut(respStruct.To),
		Src:  respStruct.TransResult[0].Src,
		Dst:  respStruct.TransResult[0].Dst,
	}
	return
}
