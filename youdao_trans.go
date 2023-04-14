package translator_engine

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"unicode/utf8"
)

type youdaoResp struct {
	Translation []string `json:"translation"`
	Query       string   `json:"query"`
}

var youdaoScene map[Scene]string

func encrypt(signStr string) string {
	if !utf8.ValidString(signStr) {
		fmt.Printf("encrypt:%q\n", signStr)
	}

	hashAlgorithm := sha256.New()
	hashAlgorithm.Write([]byte(signStr))
	return hex.EncodeToString(hashAlgorithm.Sum(nil))
}

func truncate(q string) string {
	if q == "" {
		return ""
	}
	size := utf8.RuneCountInString(q)
	if size <= 20 {
		return q
	}
	r := q[:10] + strconv.Itoa(size) + q[size-10:]
	if !utf8.ValidString(r) {
		fmt.Printf("truncate:%q\n", r)
	}
	return r
}

func (eng *YoudaoTransEngine) TransText(src, from, to string, scene ...Scene) (data *TransResult, err error) {
	from = eng.LanCodeIn(from)
	to = eng.LanCodeIn(to)
	return eng.TransTextDirect(src, from, to, scene...)
}

func (eng *YoudaoTransEngine) TransTextDirect(src, from, to string, scene ...Scene) (data *TransResult, err error) {
	respData := make(map[string]interface{})

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	salt := uuid.New().String()
	currTime := strconv.Itoa(int(time.Now().Unix()))

	var buffer bytes.Buffer
	buffer.WriteString(eng.key)
	buffer.WriteString(truncate(src))
	buffer.WriteString(salt)
	buffer.WriteString(currTime)
	buffer.WriteString(eng.secret)

	signStr := encrypt(buffer.String())

	req := url.Values{}
	req.Set("from", from)
	req.Set("to", to)
	req.Set("signType", "v3")
	req.Set("curtime", currTime)
	req.Set("appKey", eng.key)
	req.Set("q", src)
	req.Set("salt", salt)
	req.Set("sign", signStr)

	var _scene string
	if len(scene) > 0 {
		_scene = youdaoScene[scene[0]]
	}

	if len(_scene) > 0 {
		req.Set("domain", _scene)
	}

	ctx := context.Background()

	err = eng.requester.build(ctx, http.MethodPost, youdaoApi, headers, req)
	if err != nil {
		return
	}

	if err = doJsonRequest(eng.client, eng.requester.req, &respData); err != nil {
		return
	}

	if v, ok := respData["errorCode"]; ok {
		if _errMsg, ok := errMsg[v.(string)]; ok {
			err = errors.New(_errMsg)
			return
		}
	}

	var respStruct youdaoResp
	_bytes, err := json.Marshal(respData)
	if err != nil {
		return
	}
	err = json.Unmarshal(_bytes, &respStruct)
	if err != nil {
		return
	}

	if len(respStruct.Translation) == 0 {
		err = errors.New("empty translation")
	}

	data = &TransResult{
		From: from,
		To:   to,
		Src:  respStruct.Query,
		Dst:  respStruct.Translation[0],
	}

	return
}

func init() {
	youdaoScene = map[Scene]string{
		Medicine:    "medicine",
		Communicate: "computers",
		Finance:     "finance",
	}
}
