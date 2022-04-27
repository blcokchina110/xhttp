package xhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//post
func Post(url string, headers Values, bs []byte) ([]byte, error) {
	return post(url, headers, bs)
}

//post body可以带任意结构体
func PostJson(url string, headers Values, body interface{}) ([]byte, error) {
	bs, err := json.Marshal(body)
	if err != nil || bs == nil {
		return nil, err
	}

	return post(url, headers, bs)
}

//获取response body数据，并解析
func GetParseData(url string, headers Values, data interface{}) error {
	bs, status, err := get(url, headers)
	if err == nil && status == 200 && bs != nil {
		if err := json.Unmarshal(bs, &data); err != nil {
			return err
		}
	}
	return err
}

//get
func Get(url string, headers Values) ([]byte, int, error) {
	return get(url, headers)
}

//get
//-1 failure
func get(url string, headers Values) ([]byte, int, error) {
	transport := http.Transport{
		Dial:              dialTimeout,
		DisableKeepAlives: true,
	}

	client := http.Client{
		Transport: &transport,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, -1, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}

	return body, resp.StatusCode, nil
}

//post
func post(url string, headers Values, bs []byte) ([]byte, error) {
	//
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}

	//
	transport := http.Transport{
		Dial:              dialTimeout,
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	//
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

//超时处理
func dialTimeout(network, addr string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, addr, time.Second*10)
	if err != nil {
		return conn, err
	}

	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetKeepAlive(false)

	return tcpConn, err
}
