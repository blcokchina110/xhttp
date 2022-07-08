package xhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	timeoutSecond = 10
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

//post并解析返回值
func PostWithInterface(url string, headers Values, body interface{}, data interface{}) error {
	bs, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resultBytes, err := post(url, headers, bs)
	if err != nil {
		return err
	}
	return json.Unmarshal(resultBytes, &data)
}

//获取response body数据，并解析
func GetParseData(url string, headers Values, data interface{}) error {
	bs, _, err := get(url, headers)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, &data)
}

//get
func Get(url string, headers Values) ([]byte, int, error) {
	return get(url, headers)
}

//get
//-1 failure
func get(urlAddr string, headers Values) ([]byte, int, error) {
	//check url
	if _, err := url.ParseRequestURI(urlAddr); err != nil {
		return nil, -1, err
	}

	req, err := http.NewRequest(http.MethodGet, urlAddr, nil)
	if err != nil {
		return nil, -1, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	//create http client
	client := httpClient()
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
func post(urlAddr string, headers Values, bs []byte) ([]byte, error) {
	//check url
	if _, err := url.ParseRequestURI(urlAddr); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, urlAddr, bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}
	//header set
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := httpClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

//创建http client
func httpClient() http.Client {
	transport := http.Transport{
		Dial:              dialTimeout,
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	return client
}

//超时处理
func dialTimeout(network, addr string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, addr, time.Second*time.Duration(timeoutSecond))
	if err != nil {
		return conn, err
	}

	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetKeepAlive(false)

	return tcpConn, err
}
