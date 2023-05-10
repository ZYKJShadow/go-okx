package restapi

import (
	"errors"
	"fmt"
	"go-okx/define"
	"net/http"
	"net/url"
)

func (c *ApiConfig) GetPos(instType, instId, posId string) (res define.GetPosResult, err error) {

	requestUrl := fmt.Sprintf("%s?instType=%s&instId=%s&posId=%s", define.GetPos, instType, instId, posId)
	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, false)
	if err != nil {
		return
	}

	if res.Code != "0" {
		err = errors.New(res.Msg)
		return
	}

	return
}

func (c *ApiConfig) SetPosMode(setPos define.SetPosMode) (err error) {

	var res define.Common

	err = c.SendRequest(define.PostSetPosMode, &setPos, &res, http.MethodPost, false)

	if res.Code != "0" {
		err = errors.New(res.Msg)
	}

	return
}

func (c *ApiConfig) SetLeverage(setLev define.SetLeverage) (err error) {

	var res define.Common

	err = c.SendRequest(define.PostSetLeverage, &setLev, &res, http.MethodPost, false)

	if err != nil {
		return
	}

	if res.Code != "0" {
		err = errors.New(res.Msg)
	}

	return
}

func (c *ApiConfig) GetBalance(ccy []string) (res define.BalanceResult, err error) {

	requestUrl := define.GetBalanceUrl

	if len(ccy) > 0 {
		requestUrl += "?"
		for i := 0; i < len(ccy); i++ {
			coin := ccy[i]
			if i == len(ccy)-1 {
				requestUrl += coin
				break
			}
			requestUrl += coin + ","
		}
	}

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, false)

	if err != nil {
		return
	}

	if res.Code != "0" {
		err = errors.New(res.Msg)
	}

	return
}

func (c *ApiConfig) GetPositionsHistory(instType string, instId string, mgmMode string, t string, posId string, after string, before string, limit string) (res define.BalanceResult, err error) {

	params := url.Values{}

	if instType != "" {
		params.Add("instType", instType)
	}

	if instId != "" {
		params.Add("instId", instId)
	}

	if mgmMode != "" {
		params.Add("mgmMode", mgmMode)
	}

	if t != "" {
		params.Add("type", t)
	}

	if posId != "" {
		params.Add("posId", posId)
	}

	if after != "" {
		params.Add("after", after)
	}

	if before != "" {
		params.Add("before", before)
	}

	if limit != "" {
		params.Add("limit", limit)
	}

	requestUrl := fmt.Sprintf("%s?%s", define.GetPosHistory, params.Encode())

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, false)

	if err != nil {
		return
	}

	if res.Code != "0" {
		err = errors.New(res.Msg)
	}

	return
}
