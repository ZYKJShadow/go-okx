# OKX API Go SDK
This repository provides a Go interface to the OKX API. It includes the necessary components to interact with the API endpoints for placing orders, retrieving order status, getting market data, and more.
## Overview
The SDK includes the following features:

Placing Orders: You can place spot, futures, and option orders using the SDK. Both limit and market orders are supported.

Retrieving Order Status: After placing an order, you can retrieve its status from the OKX API.

Retrieving Market Data: The SDK provides functions to retrieve market data such as tickers, candlestick data, and system time.

Closing Positions: You can close open positions using the SDK.

Placing Advanced Orders: The SDK supports advanced order types like OCO (One Cancels the Other) orders.

## Installation
```go
go get -u github.com/ZYKJShadow/go-okx
```
## Usage
Before using the SDK, you need to initialize the API configuration with your OKX API credentials:
```go
c := NewApiConfig(ApiKey, SecretKey, Password, Proxy, Simulate)
```
Here, ApiKey, SecretKey, and Password are your OKX API credentials. Proxy is an optional HTTP proxy, and Simulate is a flag that indicates whether to use the OKX API's simulation mode.

With the API configuration initialized, you can use the SDK's functions to interact with the OKX API. For example, you can place a spot order as follows:
```go
order := define.Order{InstId: "BTC-USDT", TdMode: define.Cash, Side: define.Buy, OrdType: define.Market, Sz: "100", TgtCcy: define.QuoteCcy}
res, err := c.PostOrder(order)
```
## Support
If you encounter any issues or have any questions about the SDK, please open an issue in the GitHub repository.