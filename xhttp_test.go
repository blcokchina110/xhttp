package xhttp

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://ip.cn/api/index?ip=&type=0"

	bs, code, err := Get(url, nil)
	fmt.Println(string(bs), code, err)

	type data struct {
		RS       int    `json:"rs"`
		Code     int    `json:"code"`
		Address  string `json:"address"`
		IsDomain int    `json:"isDomain"`
	}
	var d data
	err = GetParseData(url, nil, &d)
	fmt.Println(d)
}
