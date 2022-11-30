package xhttp

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "http://ip-api.com/json/1111"

	bs, err := Get(url, nil)
	fmt.Println(string(bs), err)

	type data struct {
		Code int `json:"code"`
		Data struct {
			BlockNumber int64 `json:"blockNumber"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

 	var d data
	err = GetParseData(url, nil, &d)

	fmt.Println(d.Data.BlockNumber, err)
}
