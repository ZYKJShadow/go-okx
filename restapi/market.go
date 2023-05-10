package restapi

import (
	"errors"
	"fmt"
	"go-okx/define"
	"net/http"
)

func (c *ApiConfig) GetAllTickers(instType string) (res define.Tickers, err error) {

	if instType == "" {
		err = errors.New("instType empty param")
		return
	}

	err = c.SendRequest(fmt.Sprintf("%s?instType=%s", define.GetTickersUrl, instType), nil, &res, http.MethodGet, true)

	return
}

func (c *ApiConfig) GetTicker(instId string) (res define.Tickers, err error) {

	if instId == "" {
		err = errors.New("instId empty param")
		return
	}

	err = c.SendRequest(fmt.Sprintf("%s?instId=%s", define.GetTickerUrl, instId), nil, &res, http.MethodGet, true)

	if err != nil {
		return
	}

	if res.Code != "0" {
		err = errors.New(res.Msg)
		return
	}

	return
}

func (c *ApiConfig) GetHistoryCandles(instId, bar, after, before, limit string) (res define.Candles, err error) {

	if instId == "" {
		err = errors.New("instId empty param")
		return
	}

	requestUrl := fmt.Sprintf("%s?instId=%s", define.GetHisCandlesUrl, instId)

	if bar != "" {
		requestUrl = fmt.Sprintf("%s&bar=%s", requestUrl, bar)
	}

	if after != "" {
		requestUrl = fmt.Sprintf("%s&after=%s", requestUrl, after)
	}

	if before != "" {
		requestUrl = fmt.Sprintf("%s&before=%s", requestUrl, before)
	}

	if limit != "" {
		requestUrl = fmt.Sprintf("%s&limit=%s", requestUrl, limit)
	}

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, true)

	return
}

func (c *ApiConfig) GetHistoryIndexCandles(instId, bar, after, before, limit string) (res define.Candles, err error) {

	if instId == "" {
		err = errors.New("instId empty param")
		return
	}

	requestUrl := fmt.Sprintf("%s?instId=%s", define.GetHistoryIndexCandlesUrl, instId)

	if bar != "" {
		requestUrl = fmt.Sprintf("%s&bar=%s", requestUrl, bar)
	}

	if after != "" {
		requestUrl = fmt.Sprintf("%s&after=%s", requestUrl, after)
	}

	if before != "" {
		requestUrl = fmt.Sprintf("%s&before=%s", requestUrl, before)
	}

	if limit != "" {
		requestUrl = fmt.Sprintf("%s&limit=%s", requestUrl, limit)
	}

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, true)

	return
}

func (c *ApiConfig) GetIndexCandles(instId, bar, after, before, limit string) (res define.Candles, err error) {

	if instId == "" {
		err = errors.New("instId empty param")
		return
	}

	requestUrl := fmt.Sprintf("%s?instId=%s", define.GetIndexCandlesUrl, instId)

	if bar != "" {
		requestUrl = fmt.Sprintf("%s&bar=%s", requestUrl, bar)
	}

	if after != "" {
		requestUrl = fmt.Sprintf("%s&after=%s", requestUrl, after)
	}

	if before != "" {
		requestUrl = fmt.Sprintf("%s&before=%s", requestUrl, before)
	}

	if limit != "" {
		requestUrl = fmt.Sprintf("%s&limit=%s", requestUrl, limit)
	}

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, true)

	return
}

// GetCandles 接口限速1次/50ms
func (c *ApiConfig) GetCandles(instId, bar, after, before, limit string) (res define.Candles, err error) {

	if instId == "" {
		err = errors.New("instId empty param")
		return
	}

	requestUrl := fmt.Sprintf("%s?instId=%s", define.GetCandlesUrl, instId)

	if bar != "" {
		requestUrl = fmt.Sprintf("%s&bar=%s", requestUrl, bar)
	}

	if after != "" {
		requestUrl = fmt.Sprintf("%s&after=%s", requestUrl, after)
	}

	if before != "" {
		requestUrl = fmt.Sprintf("%s&before=%s", requestUrl, before)
	}

	if limit != "" {
		requestUrl = fmt.Sprintf("%s&limit=%s", requestUrl, limit)
	}

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, true)

	return
}
