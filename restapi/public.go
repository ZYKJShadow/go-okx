package restapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"go-okx/define"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (c *ApiConfig) GetSystemTime() (res define.TimeResult, err error) {
	err = c.SendRequest(define.GetTime, nil, &res, http.MethodGet, true)
	if err != nil {
		return
	}
	if res.Code != "0" {
		err = errors.New(res.Msg)
	}
	return
}

func (c *ApiConfig) SendRequest(requestUrl string, msg interface{}, res interface{}, method string, public bool) (err error) {

	globalUrl := define.RestGlobalUrl
	if c.Simulate {
		globalUrl = define.RestSimulateUrl
	}

	client := &http.Client{
		Timeout: time.Duration(1) * time.Second,
	}

	if c.Proxy != "" {
		u, _ := url.Parse(c.Proxy)
		t := &http.Transport{
			MaxIdleConns:    10,
			MaxConnsPerHost: 10,
			IdleConnTimeout: time.Duration(3) * time.Second,
			Proxy:           http.ProxyURL(u),
		}
		client.Transport = t
	}

	var body []byte = nil
	if msg != nil {
		if body, err = json.Marshal(msg); err != nil {
			return
		}
	}

	r, err := http.NewRequest(method, globalUrl+requestUrl, strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	if method == http.MethodPost {
		r.Header.Set("content-type", "application/json")
	}

	if !public {
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		r.Header.Set("OK-ACCESS-KEY", c.ApiKey)
		r.Header.Set("OK-ACCESS-SIGN", getAccessSign(method, requestUrl, string(body), c.SecretKey, timestamp))
		r.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
		r.Header.Set("OK-ACCESS-PASSPHRASE", c.Password)
	}

	if c.Simulate {
		//"x-simulated-trading: 1"  模拟盘
		r.Header.Set("x-simulated-trading", "1")
	}

	rsp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	var ret []byte
	ret, err = io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	if res != nil {
		err = json.Unmarshal(ret, &res)
	}

	return
}

func getAccessSign(method, requestPath, body, secretKey, timestamp string) string {
	// OK-ACCESS-SIGN的请求头是对timestamp + method + requestPath + body字符串（+表示字符串连接），以及SecretKey，使用HMAC SHA256方法加密，通过Base-64编码输出而得到的。
	return base64.StdEncoding.EncodeToString(hmacSha256(secretKey, timestamp+method+requestPath+body))
}

func hmacSha256(key, data string) []byte {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return h.Sum(nil)
}
