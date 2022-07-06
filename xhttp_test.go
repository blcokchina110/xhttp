package xhttp

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "api.github.com/repos/ethereum/go-ethereum/releases/latest"

	bs, code, err := Get(url, nil)

	fmt.Println(string(bs), code, err)

	// headers := make(Values)
	// headers.Set("Content-Type", "application/json")

	// bs, code, err := Get(url, headers)
	// fmt.Println(string(bs), code, err)

	// type data struct {
	// 	RS       int    `json:"rs"`
	// 	Code     int    `json:"code"`
	// 	Address  string `json:"address"`
	// 	IsDomain int    `json:"isDomain"`
	// }
	// var d data
	// err = GetParseData(url, nil, &d)
	// fmt.Println(d)
}
