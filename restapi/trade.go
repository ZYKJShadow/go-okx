package restapi

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-okx/define"
	"net/http"
)

// PostOrder 下单
func (c *ApiConfig) PostOrder(order define.Order) (res define.OrderResult, err error) {

	logrus.Infof("下单请求提交:\n产品id:%v\n交易模式:%v\n方向:%v\n订单类型:%v\n数量:%v\n计价方式:%v\n", order.InstId, order.TdMode, order.Side, order.OrdType, order.Sz, order.TgtCcy)

	err = c.SendRequest(define.OrderUrl, &order, &res, http.MethodPost, false)

	if err != nil {
		return
	}

	for i := range res.Data {
		if res.Data[i].SCode != "0" {
			err = errors.New(res.Data[i].SMsg)
			return
		}
	}

	if len(res.Data) == 0 {
		err = errors.New("empty result")
		return
	}

	return
}

// MakeMarketLongOrder 市价做多
func (c *ApiConfig) MakeMarketLongOrder(instId, sz, tpOrdPx, tpTriggerPx, slTriggerPx, slOrdPx string) (define.OrderResult, error) {

	order := define.Order{
		InstId:  instId,
		TdMode:  define.Cross,
		Side:    define.Buy,
		OrdType: define.Market,
		Sz:      sz,
		TgtCcy:  "",
		PosSide: define.MakeLong,
		Trigger: define.Trigger{
			TpTriggerPxType: define.Last,
			TpOrdPx:         tpOrdPx,
			TpTriggerPx:     tpTriggerPx,
			SlTriggerPx:     slTriggerPx,
			SlOrdPx:         slOrdPx,
			SlTriggerPxType: define.Last,
		},
	}

	return c.PostOrder(order)
}

// MakeMarketShortOrder 市价做空
func (c *ApiConfig) MakeMarketShortOrder(instId, sz, tpOrdPx, tpTriggerPx, slTriggerPx, slOrdPx string) (define.OrderResult, error) {

	order := define.Order{
		InstId:  instId,
		TdMode:  define.Cross,
		Side:    define.Buy,
		OrdType: define.Market,
		Sz:      sz,
		TgtCcy:  "",
		PosSide: define.MakeShort,
		Trigger: define.Trigger{
			TpTriggerPxType: define.Last,
			TpOrdPx:         tpOrdPx,
			TpTriggerPx:     tpTriggerPx,
			SlTriggerPx:     slTriggerPx,
			SlOrdPx:         slOrdPx,
			SlTriggerPxType: define.Last,
		},
	}

	return c.PostOrder(order)
}

// PostOrderAlgo 止盈止损等策略下单
func (c *ApiConfig) PostOrderAlgo(order define.Order) (res define.AlgoOrderResult, err error) {

	err = c.SendRequest(define.PostOrderAlgo, &order, &res, http.MethodPost, false)

	if err != nil {
		return
	}

	for i := range res.Data {
		if res.Data[i].SCode != "0" {
			err = errors.New(res.Data[i].SMsg)
			return
		}
	}

	if len(res.Data) == 0 {
		err = errors.New(res.Msg)
		return
	}

	return
}

// PostCancelOrderAlgos 撤销策略订单
func (c *ApiConfig) PostCancelOrderAlgos(cancelOrder define.CancelAlgo) (res define.Common, err error) {

	err = c.SendRequest(define.PostCancelOrderAlgos, &cancelOrder, &res, http.MethodPost, false)

	return
}

// PostClosePos 市价全平
func (c *ApiConfig) PostClosePos(closePos define.ClosePos) (res define.Common, err error) {

	err = c.SendRequest(define.PostClosePosUrl, &closePos, &res, http.MethodPost, false)

	if err != nil {
		return
	}

	if res.Code != "0" {
		err = errors.New(res.Msg)
	}

	return
}

// GetOrder 获取订单信息
func (c *ApiConfig) GetOrder(instId, ordId, clOrdId string) (res define.GetOrderResult, err error) {

	if instId == "" {
		err = errors.New("instId empty param")
		return
	}

	if ordId == "" {
		err = errors.New("ordId empty param")
		return
	}

	requestUrl := fmt.Sprintf("%s?instId=%s&ordId=%s", define.OrderUrl, instId, ordId)

	if clOrdId != "" {
		requestUrl = fmt.Sprintf("%s&clOrdId=%s", requestUrl, clOrdId)
	}

	err = c.SendRequest(requestUrl, nil, &res, http.MethodGet, false)

	if res.Code != "0" {
		err = errors.New(res.Msg)
		return
	}

	if len(res.Data) == 0 {
		err = errors.New("empty result")
		return
	}

	return
}
