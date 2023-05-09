package restapi

import (
	"errors"
	"fmt"
	"go-okx/define"
	"net/http"
)

func (c *ApiConfig) GetPos(instType, instId, posId string) (res define.GetPosResult, err error) {

	requestUrl := define.GetPos + "?"

	if instType != "" {
		requestUrl = fmt.Sprintf("%s&instType=%s", requestUrl, instType)
	}

	if instId != "" {
		requestUrl = fmt.Sprintf("%s&instId=%s", requestUrl, instId)
	}

	if posId != "" {
		requestUrl = fmt.Sprintf("%s&posId=%s", requestUrl, posId)
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
