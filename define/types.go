package define

type TimeResult struct {
	Common
	Data []struct {
		Ts string `json:"ts"`
	} `json:"data"`
}

type Tickers struct {
	Data []Ticker
	Common
}

type Ticker struct {
	Last      string `json:"last" gorm:"column:last"`
	LastSz    string `json:"lastSz" gorm:"column:lastSz"`
	Open24h   string `json:"open24h" gorm:"column:open24h"`
	AskSz     string `json:"askSz" gorm:"column:askSz"`
	Low24h    string `json:"low24h" gorm:"column:low24h"`
	AskPx     string `json:"askPx" gorm:"column:askPx"`
	VolCcy24h string `json:"volCcy24h" gorm:"column:volCcy24h"`
	InstType  string `json:"instType" gorm:"column:instType"`
	InstID    string `json:"instId" gorm:"column:instId"`
	BidSz     string `json:"bidSz" gorm:"column:bidSz"`
	BidPx     string `json:"bidPx" gorm:"column:bidPx"`
	High24h   string `json:"high24h" gorm:"column:high24h"`
	SodUtc0   string `json:"sodUtc0" gorm:"column:sodUtc0"`
	Vol24h    string `json:"vol24h" gorm:"column:vol24h"`
	SodUtc8   string `json:"sodUtc8" gorm:"column:sodUtc8"`
	Ts        string `json:"ts" gorm:"column:ts"`
}

type Candles struct {
	Common
	Data [][]string `json:"data"`
}

type Order struct {
	InstId     string `json:"instId"`
	TdMode     string `json:"tdMode"`            // 交易模式
	ClOrdId    string `json:"clOrdId,omitempty"` // 自定义订单id
	Side       string `json:"side"`
	OrdType    string `json:"ordType"`
	Px         string `json:"px,omitempty"` // 委托价格
	Sz         string `json:"sz"`           // 委托数量
	ReduceOnly bool   `json:"reduceOnly,omitempty"`
	PosSide    string `json:"posSide"` // side为sell，posSide为short时表示卖出开空  side为buy，posSide为long时表示买入开多
	TgtCcy     string `json:"tgtCcy,omitempty"`
	BanAmend   bool   `json:"banAmend,omitempty"`
	Trigger
}

type Trigger struct {
	TpTriggerPx     string `json:"tpTriggerPx,omitempty"`     // 止盈触发价
	TpTriggerPxType string `json:"tpTriggerPxType,omitempty"` // 止盈触发价类型
	TpOrdPx         string `json:"tpOrdPx,omitempty"`         // 止盈委托价，为-1时执行市价止盈
	SlTriggerPx     string `json:"slTriggerPx,omitempty"`     // 止损触发价
	SlTriggerPxType string `json:"slTriggerPxType,omitempty"` // 止损触发类型
	SlOrdPx         string `json:"slOrdPx,omitempty"`         // 止损委托价，为-1时执行市价止损
}

type CancelAlgo struct {
	AlgoId string `json:"algoId"`
	InstId string `json:"instId"`
}

type SetPosMode struct {
	PosMode string `json:"posMode"`
}

type SetLeverage struct {
	InstId  string `json:"instId"`
	Lever   string `json:"lever"`
	MgnMode string `json:"mgnMode"`
}

type ClosePos struct {
	InstId  string `json:"instId"`
	PosSide string `json:"posSide"` // 持仓方向
	MgnMode string `json:"mgnMode"` // 保证金模式
	Ccy     string `json:"ccy"`
	AutoCxl bool   `json:"autoCxl"` // 当市价全平时，平仓单是否需要自动撤销,默认为false
}

type AlgoOrderResult struct {
	Common
	Data []struct {
		AlgoId string `json:"algoId"`
		SCode  string `json:"sCode"`
		SMsg   string `json:"sMsg"`
	}
}

type OrderResult struct {
	Common
	Data []struct {
		ClOrdId string `json:"clOrdId,omitempty"`
		OrdId   string `json:"ordId"`
		Tag     string `json:"tag"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	}
}

type GetPosResult struct {
	Common
	Data []struct {
		Adl         string `json:"adl"`
		AvailPos    string `json:"availPos"` // 可平仓数量
		AvgPx       string `json:"avgPx"`    // 开仓均价
		CTime       string `json:"cTime"`
		Ccy         string `json:"ccy"`
		DeltaBS     string `json:"deltaBS"`
		DeltaPA     string `json:"deltaPA"`
		GammaBS     string `json:"gammaBS"`
		GammaPA     string `json:"gammaPA"`
		Imr         string `json:"imr"`
		InstId      string `json:"instId"`
		InstType    string `json:"instType"`
		Interest    string `json:"interest"`
		Last        string `json:"last"`
		UsdPx       string `json:"usdPx"`
		Lever       string `json:"lever"`
		Liab        string `json:"liab"`
		LiabCcy     string `json:"liabCcy"`
		LiqPx       string `json:"liqPx"` // 预估强平价
		MarkPx      string `json:"markPx"`
		Margin      string `json:"margin"` // 保证金余额，可增减，仅适用于逐仓
		MgnMode     string `json:"mgnMode"`
		MgnRatio    string `json:"mgnRatio"` // 保证金率
		Mmr         string `json:"mmr"`
		NotionalUsd string `json:"notionalUsd"`
		OptVal      string `json:"optVal"`
		PTime       string `json:"pTime"`
		Pos         string `json:"pos"`
		PosCcy      string `json:"posCcy"`
		PosId       string `json:"posId"`
		PosSide     string `json:"posSide"`
		ThetaBS     string `json:"thetaBS"`
		ThetaPA     string `json:"thetaPA"`
		TradeId     string `json:"tradeId"`
		QuoteBal    string `json:"quoteBal"`
		BaseBal     string `json:"baseBal"`
		UTime       string `json:"uTime"`
		Upl         string `json:"upl"`
		UplRatio    string `json:"uplRatio"`
		VegaBS      string `json:"vegaBS"`
		VegaPA      string `json:"vegaPA"`
	} `json:"data"`
}

type GetOrderResult struct {
	Common
	Data []struct {
		InstType        string `json:"instType"` // 产品类型
		InstId          string `json:"instId"`
		Ccy             string `json:"ccy"` // 保证金币种，仅适用于单币种保证金模式下的全仓币币杠杆订单
		OrdId           string `json:"ordId"`
		ClOrdId         string `json:"clOrdId,omitempty"` //客户自定义订单ID
		Tag             string `json:"tag,omitempty"`
		Px              string `json:"px"`  // 委托价格
		Sz              string `json:"sz"`  // 委托数量
		Pnl             string `json:"pnl"` // 收益
		OrdType         string `json:"ordType"`
		Side            string `json:"side"`
		PosSide         string `json:"posSide"`
		TdMode          string `json:"tdMode"`
		AccFillSz       string `json:"accFillSz"`
		FillPx          string `json:"fillPx"`
		TradeId         string `json:"tradeId"`
		FillSz          string `json:"fillSz"`
		FillTime        string `json:"fillTime"`
		Source          string `json:"source"`
		State           string `json:"state"`           // 订单状态
		AvgPx           string `json:"avgPx"`           // 成交均价，如果成交数量为0，该字段也为0
		Lever           string `json:"lever"`           // 杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
		TpTriggerPx     string `json:"tpTriggerPx"`     // 止盈触发价
		TpTriggerPxType string `json:"tpTriggerPxType"` // 止盈触发价类型
		TpOrdPx         string `json:"tpOrdPx"`         // 止盈委托价
		SlTriggerPx     string `json:"slTriggerPx"`     // 止损触发价
		SlTriggerPxType string `json:"slTriggerPxType"`
		SlOrdPx         string `json:"slOrdPx"`
		FeeCcy          string `json:"feeCcy"` // 交易手续费币种
		Fee             string `json:"fee"`    // 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为负数。如： -0.01
		RebateCcy       string `json:"rebateCcy"`
		Rebate          string `json:"rebate"`
		TgtCcy          string `json:"tgtCcy"`
		Category        string `json:"category"` // 订单种类
		UTime           string `json:"uTime"`    // 订单状态更新时间，Unix时间戳的毫秒数格式，如：1597026383085
		CTime           string `json:"cTime"`    // 订单创建时间，Unix时间戳的毫秒数格式， 如 ：1597026383085
	} `json:"data"`
}

type AccountPosRiskResult struct {
	Common
	Data []struct {
		AdjEq   string `json:"adjEq"`
		BalData []struct {
			Ccy   string `json:"ccy"`
			DisEq string `json:"disEq"`
			Eq    string `json:"eq"`
		} `json:"balData"`
		PosData []struct {
			BaseBal     string `json:"baseBal"`
			Ccy         string `json:"ccy"`
			InstId      string `json:"instId"`
			InstType    string `json:"instType"`
			MgnMode     string `json:"mgnMode"`
			NotionalCcy string `json:"notionalCcy"`
			NotionalUsd string `json:"notionalUsd"`
			Pos         string `json:"pos"`
			PosCcy      string `json:"posCcy"`
			PosId       string `json:"posId"`
			PosSide     string `json:"posSide"`
			QuoteBal    string `json:"quoteBal"`
		} `json:"posData"`
		Ts string `json:"ts"`
	} `json:"data"`
}

type BalanceResult struct {
	Common
	Data []struct {
		AdjEq   string `json:"adjEq"`
		Details []struct {
			AvailBal      string `json:"availBal"`
			AvailEq       string `json:"availEq"`
			CashBal       string `json:"cashBal"`
			Ccy           string `json:"ccy"`
			CrossLiab     string `json:"crossLiab"`
			DisEq         string `json:"disEq"`
			Eq            string `json:"eq"`
			EqUsd         string `json:"eqUsd"`
			FrozenBal     string `json:"frozenBal"`
			Interest      string `json:"interest"`
			IsoEq         string `json:"isoEq"`
			IsoLiab       string `json:"isoLiab"`
			IsoUpl        string `json:"isoUpl"`
			Liab          string `json:"liab"`
			MaxLoan       string `json:"maxLoan"`
			MgnRatio      string `json:"mgnRatio"`
			NotionalLever string `json:"notionalLever"`
			OrdFrozen     string `json:"ordFrozen"`
			Twap          string `json:"twap"`
			UTime         string `json:"uTime"`
			Upl           string `json:"upl"`
			UplLiab       string `json:"uplLiab"`
			StgyEq        string `json:"stgyEq"`
			SpotInUseAmt  string `json:"spotInUseAmt"`
		} `json:"details"`
		Imr         string `json:"imr"`
		IsoEq       string `json:"isoEq"`
		MgnRatio    string `json:"mgnRatio"`
		Mmr         string `json:"mmr"`
		NotionalUsd string `json:"notionalUsd"`
		OrdFroz     string `json:"ordFroz"`
		TotalEq     string `json:"totalEq"`
		UTime       string `json:"uTime"`
	} `json:"data"`
}

type WSCommon struct {
	Event string `json:"event"`
	Code  string `json:"code,omitempty"`
	Msg   string `json:"msg,omitempty"`
}

type WSKline struct {
	Data [][]string `json:"data"`
}

type Positions struct {
	AvgPx   float64
	MgnMode string
}
