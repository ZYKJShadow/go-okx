package define

type Common struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

const (
	TimeMaxDiff = 2000 // 时间误差最大值，大于此值不能够下单
)

// 实盘
const (
	RestGlobalUrl      = "https://aws.okx.com"
	SocketPubGlobalUrl = "wss://wsaws.okx.com:8443/ws/v5/public"
	SocketPriGlobalUrl = "wss://wsaws.okx.com:8443/ws/v5/private"
)

// 模拟盘
const (
	RestSimulateUrl = "https://www.okx.com"
	SocketSimPubUrl = "wss://wspap.okx.com:8443/ws/v5/public?brokerId=9999"
	SocketSimPriUrl = "wss://wspap.okx.com:8443/ws/v5/private?brokerId=9999"
)

// account url
const (
	GetPos          = "/api/v5/account/positions"
	PostSetPosMode  = "/api/v5/account/set-position-mode"
	PostSetLeverage = "/api/v5/account/set-leverage"
)

// public url
const (
	GetTime = "/api/v5/public/time"
)

// market url
const (
	GetTickersUrl    = "/api/v5/market/GetTickersUrl"
	GetCandlesUrl    = "/api/v5/market/candles"
	GetHisCandlesUrl = "/api/v5/market/history-candles"
	GetTickerUrl     = "/api/v5/market/ticker"
)

// trade url
const (
	OrderUrl             = "/api/v5/trade/order" // 直接下单
	PostClosePosUrl      = "/api/v5/trade/close-position"
	PostOrderAlgo        = "/api/v5/trade/order-algo"   // 包含止盈止损的下单
	PostCancelOrderAlgos = "/api/v5/trade/cancel-algos" // 撤销策略订单
)

// 时间粒度
const (
	Min        = "1m"
	ThreeMin   = "3m"
	FiveMin    = "5m"
	FifteenMin = "15m"
	ThirtyMin  = "30m"
	H          = "1H"
	TwoH       = "2H"
	FourH      = "4H"
	Day        = "1D"
	TwoDay     = "2D"
	ThreeDay   = "3D"
	Week       = "1W"
	Mon        = "1M"
	Year       = "1Y"
)

// 毫秒
const (
	FiveMill    = 5 * 60 * 1000
	FifteenMill = 15 * 60 * 1000
	ThirtyMill  = 30 * 60 * 1000
	HMill       = 60 * 60 * 1000
	FourHMill   = 4 * HMill
)

// 产品类型
const (
	SPOT    = "SPOT"    // 币币
	SWAP    = "SWAP"    // 永续合约
	FUTURES = "FUTURES" // 交割合约
	OPTION  = "OPTION"  // 期权
)

type Trend int

// 趋势
const (
	Up   Trend = 1
	Down Trend = -1
)

type Signal int

// 信号
const (
	Wait Signal = iota
	Long
	Short
)

type NetworkMode int

// 网络模式
const (
	HttpMode NetworkMode = iota + 1
	SocketMode
)

// 普通委托订单类型
const (
	Market   = "market"
	Limit    = "limit"
	PostOnly = "post_only"
)

// 条件委托订单类型
const (
	Condition     = "condition"       // 单向止盈止损
	Oco           = "oco"             // 双向止盈止损
	Plan          = "trigger"         // 计划委托
	MoveOrderStop = "move_order_stop" // 移动止盈止损
	Iceberg       = "iceberg"         // 冰山委托
	Twap          = "twap"            // 时间加权
)

// 交易模式
const (
	Isolated = "isolated" // 逐仓
	Cross    = "cross"    // 全仓
	Cash     = "cash"
)

// 方向
const (
	Buy       = "buy"
	Sell      = "sell"
	MakeLong  = "long"
	MakeShort = "short"
)

// 市价单委托数量的类型
const (
	BaseCcy  = "base_ccy"  // 交易货币
	QuoteCcy = "quote_ccy" // 计价货币
)

// 平仓价格类型
const (
	Last  = "last"  // 最新价格
	Index = "index" // 指数价格
	Mark  = "mark"  // 标记价格
)

// websocket
const (
	CandleChannel = "candle"
)

// 持仓模式
const (
	LongShortMode = "long_short_mode"
	NetMode       = "net_mode"
)

type WebSocketMsg struct {
	Id   string         `json:"id,omitempty"`
	Op   string         `json:"op"`
	Args []WebSocketArg `json:"args"`
}

type WebSocketArg struct {
	Order Order `json:"order,omitempty"`
}
