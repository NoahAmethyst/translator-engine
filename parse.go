package translator_engine

import (
	"encoding/json"
	"io"
	"net/http"
)

func doJsonRequest(cli *http.Client, req *http.Request, data any) (err error) {
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &data)
	return
}
