package websocketapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ZYKJShadow/recws"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
	"go-okx/define"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type WS struct {
	sync.RWMutex

	wsURL string

	ctx    context.Context
	cancel context.CancelFunc
	conn   recws.RecConn

	subscriptions map[string]interface{}

	klineCallback func(kline define.WSKline, bar string)
}

func NewWS(wsURL string) *WS {
	ws := &WS{
		wsURL:         wsURL,
		subscriptions: make(map[string]interface{}),
	}
	ws.ctx, ws.cancel = context.WithCancel(context.Background())
	ws.conn = recws.RecConn{
		KeepAliveTimeout: 10 * time.Second,
	}
	ws.conn.SubscribeHandler = ws.subscribeHandler
	return ws
}

func (ws *WS) SetProxy(proxyURL string) (err error) {
	if proxyURL == "" {
		return
	}
	var purl *url.URL
	purl, err = url.Parse(proxyURL)
	if err != nil {
		return
	}
	logx.Infof("[ws][%s] proxy url:%s", proxyURL, purl)
	ws.conn.Proxy = http.ProxyURL(purl)
	return
}

func (ws *WS) Start() {
	logx.Infof("wsURL: %v", ws.wsURL)
	ws.conn.Dial(ws.wsURL, nil)
	go ws.run()
}

func (ws *WS) SetKlineCallBack(callback func(kline define.WSKline, bar string)) {
	ws.klineCallback = callback
}

func (ws *WS) Subscribe(id string, op interface{}) error {
	ws.Lock()
	defer ws.Unlock()
	ws.subscriptions[id] = op
	return ws.sendWSMessage(op)
}

// Unsubscribe 取消订阅
func (ws *WS) Unsubscribe(id string) error {
	ws.Lock()
	defer ws.Unlock()

	if _, ok := ws.subscriptions[id]; ok {
		delete(ws.subscriptions, id)
	}
	return nil
}

func (ws *WS) sendWSMessage(msg interface{}) error {
	return ws.conn.WriteJSON(msg)
}

func (ws *WS) PostOrder(order define.Order) error {

	generateUUID := uuid.New().String()

	msg := define.WebSocketMsg{
		Id:   generateUUID,
		Op:   "order",
		Args: []define.WebSocketArg{{order}},
	}

	return ws.sendWSMessage(msg)
}

func (ws *WS) SubscribeKline(instId, interval string) error {

	type Arg struct {
		Channel string `json:"channel"`
		InstId  string `json:"instId"`
	}

	type Sub struct {
		Op   string `json:"op"`
		Args []Arg  `json:"args"`
	}

	return ws.Subscribe(instId, Sub{
		Op: "subscribe",
		Args: []Arg{
			{
				Channel: fmt.Sprintf("%v%v", define.CandleChannel, interval),
				InstId:  instId,
			},
		},
	})

}

func (ws *WS) run() {
	ctx := context.Background()
	for {
		select {
		case <-ctx.Done():
			go ws.conn.Close()
			logx.Infof("Websocket closed %s", ws.conn.GetURL())
			return
		default:
			_, msg, err := ws.conn.ReadMessage()
			if err != nil {
				logx.Errorf("Read error: %v", err)
				time.Sleep(time.Millisecond * 200)
				continue
			}
			ws.handleMsg(msg)
		}
	}
}

func (ws *WS) subscribeHandler() error {
	for _, v := range ws.subscriptions {
		//log.Printf("sub: %#v", v)
		err := ws.sendWSMessage(v)
		if err != nil {
			logx.Errorf("%v", err)
			return err
		}
	}
	return nil
}

func (ws *WS) handleMsg(msg []byte) {

	var wsCommon define.WSCommon
	err := json.Unmarshal(msg, &wsCommon)
	if err != nil {
		logx.Error(err)
		return
	}

	switch wsCommon.Event {
	case "error":
		logx.Error(wsCommon.Msg)
		return
	case "subscribe":
		logx.Info("subscribe successfully")
		return
	default:
		ret := gjson.ParseBytes(msg)
		channel := ret.Get("arg").Get("channel").String()
		if strings.Contains(channel, "candle") {
			var kline define.WSKline
			err = json.Unmarshal(msg, &kline)
			if err != nil {
				logx.Error(err)
				return
			}
			switch channel {
			case "candle5m":
				ws.klineCallback(kline, define.FiveMin)
			case "candle15m":
				ws.klineCallback(kline, define.FifteenMin)
			case "candle30m":
				ws.klineCallback(kline, define.ThirtyMin)
			case "candle1H":
				ws.klineCallback(kline, define.H)
			case "candle4H":
				ws.klineCallback(kline, define.FourH)
			default:
				logx.Errorf("not select channel %s", channel)
			}
		}
	}

}
