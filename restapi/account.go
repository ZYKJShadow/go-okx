package restapi

import (
	"errors"
	"fmt"
	"github.com/ZYKJShadow/go-okx/define"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

// GetPos 查看持仓信息
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

// SetPosMode 设置持仓模式
func (c *ApiConfig) SetPosMode(setPos define.SetPosMode) (err error) {

	var res define.Common

	err = c.SendRequest(define.PostSetPosMode, &setPos, &res, http.MethodPost, false)

	if res.Code != "0" {
		err = errors.New(res.Msg)
	}

	return
}

// SetLeverage 设置杠杆倍数
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

// GetBalance 查询账户余额
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

// GetPositionsHistory 查询历史持仓信息
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

// GetAccountPosRisk 查询账户持仓风险
func (c *ApiConfig) GetAccountPosRisk(instId string) (res define.AccountPosRiskResult, err error) {

	requestUrl := fmt.Sprintf("%s?instId=%s", define.GetAccountPosRisk, instId)

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, false)

	if err != nil {
		logrus.Error(err)
		return
	}

	if res.Code != "0" {
		err = errors.New(res.Msg)
	}

	return
}

// GetAccountBills 账单流水查询（近七天）
func (c *ApiConfig) GetAccountBills(instType, ccy, mgmMode, ctType, billType, subType, after, before, begin, end, limit string) (res define.AccountBillsResult, err error) {

	requestUrl := fmt.Sprintf("%s?instType=%s&ccy=%s&mgmMode=%s&ctType=%s&billType=%s&subType=%s&after=%s&before=%s&begin=%s&end=%s&limit=%s", define.GetAccountBills, instType, ccy, mgmMode, ctType, billType, subType, after, before, begin, end, limit)
	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, false)
	if err != nil {
		logrus.Error(err)
		return
	}

	if res.Code != "0" {
		err = errors.New(res.Msg)
	}

	return
}

// GetAccountBillsArchive 账单流水查询（近三个月）
func (c *ApiConfig) GetAccountBillsArchive(instType, ccy, mgmMode, ctType, billType, subType, after, before, begin, end, limit string) (res define.AccountBillsResult, err error) {

	requestUrl := fmt.Sprintf("%s?instType=%s&ccy=%s&mgmMode=%s&ctType=%s&billType=%s&subType=%s&after=%s&before=%s&begin=%s&end=%s&limit=%s", define.GetAccountBillsArchive, instType, ccy, mgmMode, ctType, billType, subType, after, before, begin, end, limit)
	err = c.SendRequest(requestUrl, nil, res, http.MethodGet, false)
	if err != nil {
		logrus.Error(err)
		return
	}

	return

}

// GetAccountConfig 查询账户配置
func (c *ApiConfig) GetAccountConfig() (res define.AccountConfigResult, err error) {

	err = c.SendRequest(define.GetAccountConfig, nil, &res, http.MethodGet, false)

	if err != nil {
		return
	}

	if res.Code != "0" {
		err = errors.New(res.Msg)
	}

	return

}

// GetMaxSize 获取最大可买卖/开仓数量
func (c *ApiConfig) GetMaxSize(instId, tdMode, ccy, px, leverage string, upSpotOffset bool) (res define.AccountMaxSizeResult, err error) {

	if instId == "" {
		err = errors.New("instId is required")
		return
	}

	if tdMode == "" {
		err = errors.New("tdMode is required")
		return
	}

	if ccy == "" {
		err = errors.New("ccy is required")
		return
	}

	requestUrl := fmt.Sprintf("%s?instId=%s&tdMode=%s&ccy=%s&px=%s&leverage=%s&upSpotOffset=%t", define.GetAccountMaxSize, instId, tdMode, ccy, px, leverage, upSpotOffset)

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, false)

	return

}

// GetMaxAvailSize 获取最大可用数量
func (c *ApiConfig) GetMaxAvailSize(instId, tdMode, ccy, quickMgnType string, upSpotOffset, reduceOnly bool) (res define.AccountMaxAvailSizeResult, err error) {

	if instId == "" {
		err = errors.New("instId is required")
		return
	}

	if tdMode == "" {
		err = errors.New("tdMode is required")
		return
	}

	requestUrl := fmt.Sprintf("%s?instId=%s&tdMode=%s&ccy=%s&quickMgnType=%s&upSpotOffset=%t&reduceOnly=%t", define.GetAccountMaxAvailSize, instId, tdMode, ccy, quickMgnType, upSpotOffset, reduceOnly)

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, false)

	return

}

func (c *ApiConfig) PostMarginBalance(mb define.MarginBalance) (res define.AccountMarginBalanceResult, err error) {

	if mb.InstId == "" {
		err = errors.New("instId is required")
		return
	}

	if mb.PosSide == "" {
		err = errors.New("posSide is required")
		return
	}

	if mb.Type == "" {
		err = errors.New("marginType is required")
		return
	}

	if mb.Amt == "" {
		err = errors.New("amt is required")
		return
	}

	err = c.SendRequest(define.PostAccountMarginBalance, &mb, &res, http.MethodPost, false)

	return
}

// GetMaxLoan 获取最大可借
func (c *ApiConfig) GetMaxLoan(instId, mgnMode, mgnCcy string) (res define.AccountMaxLoanResult, err error) {

	if instId == "" {
		err = errors.New("instId is required")
		return
	}

	if mgnMode == "" {
		err = errors.New("mgnMode is required")
		return
	}

	if mgnCcy == "" {
		err = errors.New("mgnCcy is required")
		return
	}

	requestUrl := fmt.Sprintf("%s?instId=%s&mgnMode=%s&mgnCcy=%s", define.GetAccountMaxLoan, instId, mgnMode, mgnCcy)

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, false)

	return
}
