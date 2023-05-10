package restapi

import (
	"fmt"
	"go-okx/define"
	"testing"
)

const (
	ApiKey    = ""
	SecretKey = ""
	Password  = ""
	Proxy     = ""
	Simulate  = true
)

func TestGetOrder(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	order := define.Order{InstId: "BTC-USDT-SWAP", TdMode: define.Cross, Side: define.Buy, PosSide: define.MakeLong, OrdType: define.Market, Sz: "100"}

	res, err := c.PostOrder(order)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := c.GetOrder("BTC-USDT-SWAP", res.Data[0].OrdId, res.Data[0].ClOrdId)
	if err != nil {
		t.Error(err)
		return
	}

	if resp.Code != "0" {
		t.Error(resp.Msg)
		return
	}

	orderRes := resp.Data[0]

	fmt.Printf("产品id:%s\n收益:%s\n订单状态:%s\n杠杆:%s\n持仓方向:%s\n成交价格:%s\n", orderRes.InstId, orderRes.Pnl, orderRes.State, orderRes.Lever, orderRes.PosSide, orderRes.AvgPx)

}

// 币币交易
func TestCashPostOrder(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	// 买入价值100刀的BTC
	order := define.Order{InstId: "BTC-USDT", TdMode: define.Cash, Side: define.Buy, OrdType: define.Market, Sz: "100", TgtCcy: define.QuoteCcy}

	res, err := c.PostOrder(order)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res)

}

// 合约交易
func TestContractPostOrder(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	order := define.Order{InstId: "BTC-USDT-SWAP", TdMode: define.Cross, Side: define.Buy, PosSide: define.MakeLong, OrdType: define.Market, Sz: "100"}

	res, err := c.PostOrder(order)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res)
}

// 市价全平
func TestPostClosePos(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	order := define.Order{InstId: "BTC-USDT-SWAP", TdMode: define.Cross, Side: define.Buy, PosSide: define.MakeLong, OrdType: define.Market, Sz: "100"}

	_, err := c.PostOrder(order)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := c.PostClosePos(define.ClosePos{InstId: "BTC-USDT-SWAP", MgnMode: define.Cross, PosSide: define.MakeLong, Ccy: "USDT", AutoCxl: true})
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res)
}

// 策略下单
func TestPostAlgoOrder(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	// 双向止盈止损
	order := define.Order{
		InstId:  "BTC-USDT-SWAP",
		TdMode:  define.Cross,
		Side:    define.Buy,
		PosSide: define.MakeLong,
		OrdType: define.Oco,
		Sz:      "100",
		Trigger: define.Trigger{
			TpOrdPx:         "30000",
			SlOrdPx:         "10000",
			TpTriggerPx:     "1",
			SlTriggerPx:     "100000",
			TpTriggerPxType: define.Last,
			SlTriggerPxType: define.Last,
		},
	}

	res, err := c.PostOrderAlgo(order)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res)
}

// 获取系统时间
func TestGetSystemTime(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, false)

	res, err := c.GetSystemTime()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

}

// 获取所有产品行情信息
func TestGetAllTickers(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, false)

	// 永续合约产品行情信息
	res, err := c.GetAllTickers(define.SWAP)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

	// 币币产品行情信息
	res, err = c.GetAllTickers(define.SPOT)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

	// 交割合约产品行情信息
	res, err = c.GetAllTickers(define.FUTURES)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

	// 期权合约产品行情信息
	res, err = c.GetAllTickers(define.OPTION)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

}

// 获取单个产品行情信息
func TestGetTicker(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	// 永续合约产品行情信息
	res, err := c.GetTicker("BTC-USDT-SWAP")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

	// 币币产品行情信息
	res, err = c.GetTicker("BTC-USDT")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

	// 交割合约产品行情信息
	res, err = c.GetTicker("BTC-USD-210326")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

	// 期权合约产品行情信息
	res, err = c.GetTicker("BTC-USD-210326-40000-C")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

}

// 获取产品K线数据
func TestGetCandles(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	// 永续合约产品行情信息
	res, err := c.GetCandles("BTC-USDT-SWAP", define.H, "", "", "100")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)
}

// 获取产品历史K线数据
func TestGetHistoryCandles(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	// 永续合约产品行情信息
	res, err := c.GetHistoryCandles("BTC-USDT-SWAP", define.H, "1683687600000", "", "1")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)
}

// 获取指数K线数据
func TestGetIndexCandles(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	// 永续合约产品行情信息
	res, err := c.GetIndexCandles("BTC-USDT", define.H, "", "", "100")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

}

func TestGetHistoryIndexCandles(t *testing.T) {

	c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)

	// 永续合约产品行情信息
	res, err := c.GetHistoryIndexCandles("BTC-USDT", define.H, "1683687600000", "", "1")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", res.Data)

}
