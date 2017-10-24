// Package goincheck is client for Coincheck Exchange API
package goincheck

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strconv"
	"time"
)

const defalutBaseURL = "https://coincheck.com"

var userAgent = fmt.Sprintf("CoinCheckGoClient/%s (%s)", version, runtime.Version())

// Client struct represents Coincheck API client.
type Client struct {
	accessKey       string
	secretAccessKey string

	BaseURL *url.URL

	HTTPClient *http.Client
}

// Ticker struct represents Coincheck Ticker API Response.
type Ticker struct {
	Last      float64 `json:"last"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
	Timestamp float64 `json:"timestamp"`
}

// Trade struct represents Coincheck Trade API Respoonse.
type Trade struct {
	ID        int     `json:"id"`
	Amount    string  `json:"amount"`
	Rate      float64 `json:"rate"`
	OrderType string  `json:"order_type"`
	CreatedAt string  `json:"created_at"`
}

// OrderBook struct represents Coincheck OrderBook API Response.
type OrderBook struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

type exchangeRateParam struct {
	OrderType string  `json:"order_type"`
	Pair      string  `json:"pair"`
	Amount    float64 `json:"amount"`
	Price     int     `json:"price"`
}

// ExchangeRate struct represents Coincheck ExchangeRate API Response.
type ExchangeRate struct {
	Success bool `json:"success"`
	Rate    int  `json:"rate"`
	Price   int  `json:"price"`
	Amount  int  `json:"amount"`
}

// Order struct represents Coincheck Order API Response.
type Order struct {
	Success      bool    `json:"success"`
	ID           int     `json:"id"`
	Pair         string  `json:"pair"`
	OrderType    string  `json:"order_type"`
	Amount       float64 `json:"amount"`
	Rate         int     `json:"rate"`
	StopLossRate int     `json:"stop_less_rate"`
	CreatedAt    string  `json:"created_at"`
	Error        string  `json:"error"`
}

type orderParam struct {
	Pair            string  `json:"pair"`
	OrderType       string  `json:"order_type"`
	Amount          float64 `json:"amount"`
	Rate            int     `json:"rate"`
	MarketBuyAmount int     `json:"market_buy_amount"`
	PositionID      int     `json:"position_id"`
	StopLossRate    int     `json:"stop_less_rate"`
}

// Balance struct represents Coincheck Balance API Response.
type Balance struct {
	Success      bool   `json:"success"`
	Jpy          string `json:"jpy"`
	Btc          string `json:"btc"`
	JpyReserved  string `json:"jpy_reserved"`
	BtcReserved  string `json:"btc_reserved"`
	JpyLendInUse string `json:"jpy_lend_in_use"`
	BtcLendInUse string `json:"btc_lend_in_use"`
	JpyLend      string `json:"jpy_lent"`
	BtcLend      string `json:"btc_lent"`
	JpyDebt      string `json:"jpy_debt"`
	BtcDebt      string `json:"btc_debt"`
}

// RatePair struct reporesents Coincheck RatePair API Response.
type RatePair struct {
	Rate string `json:"rate"`
}

// NewClient creats Client Struct.
func NewClient(key, secretKey string) (*Client, error) {
	if key == "" || secretKey == "" {
		return &Client{}, errors.New("key is missing")
	}

	baseurl, _ := url.Parse(defalutBaseURL)
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	cli := &Client{accessKey: key, secretAccessKey: secretKey, BaseURL: baseurl, HTTPClient: client}

	return cli, nil
}

func (cli *Client) GetTicker(ctx context.Context) (*Ticker, error) {
	req, err := cli.newRequest(ctx, http.MethodGet, "/api/ticker", []byte(""))
	if err != nil {
		return nil, err
	}

	res, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var ticker Ticker
	err = decodeBody(res, &ticker)
	if err != nil {
		return nil, err
	}

	return &ticker, nil
}

func (cli *Client) GetTrade(ctx context.Context) (*[]Trade, error) {
	req, err := cli.newRequest(ctx, http.MethodGet, "/api/trades", []byte(""))
	if err != nil {
		return nil, err
	}

	res, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var trades []Trade
	err = decodeBody(res, &trades)
	if err != nil {
		return nil, err
	}

	return &trades, nil
}

func (cli *Client) GetOrderBook(ctx context.Context) (*OrderBook, error) {
	req, err := cli.newRequest(ctx, http.MethodGet, "/api/order_books", []byte(""))
	if err != nil {
		return nil, err
	}

	res, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var orderbook OrderBook
	err = decodeBody(res, &orderbook)
	if err != nil {
		return nil, err
	}

	return &orderbook, nil
}

func (cli *Client) GetRatePair(ctx context.Context, pair Pair) (*RatePair, error) {
	endpoint := "/api/rate/" + string(pair)
	req, err := cli.newRequest(ctx, http.MethodGet, endpoint, []byte(""))
	if err != nil {
		return nil, err
	}

	res, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var ratePair RatePair
	err = decodeBody(res, &ratePair)
	if err != nil {
		return nil, err
	}

	return &ratePair, nil
}

func (cli *Client) GetExchangeRate(ctx context.Context) (*ExchangeRate, error) {
	req, err := cli.newRequest(ctx, http.MethodGet, "/api/exchange/orders/rate", []byte(""))
	if err != nil {
		return nil, err
	}

	res, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var rate ExchangeRate
	err = decodeBody(res, &rate)
	if err != nil {
		return nil, err
	}
	return &rate, nil
}

func (cli *Client) newOrder(ctx context.Context, param *orderParam) (*Order, error) {
	data, err := encodeBody(param)
	if err != nil {
		return nil, err
	}

	req, err := cli.newRequest(ctx, http.MethodPost, "/api/exchange/orders", data)
	if err != nil {
		return nil, err
	}

	res, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var order Order
	err = decodeBody(res, &order)
	if err != nil {
		return nil, err
	}

	if order.Success == false {
		return nil, errors.New(order.Error)
	}

	return &order, nil
}

func (cli *Client) OrderToBuy(ctx context.Context, rate int, amount float64) (*Order, error) {
	if rate < 0 {
		return nil, errors.New("rate is negative")
	}
	if amount < 0 {
		return nil, errors.New("amount is negative")
	}

	param := orderParam{OrderType: "buy", Pair: "btc_jpy", Rate: rate, Amount: amount}

	return cli.newOrder(ctx, &param)
}

func (cli *Client) OrderToSell(ctx context.Context, rate int, amount float64) (*Order, error) {
	if rate < 0 {
		return nil, errors.New("rate is negative")
	}
	if amount < 0 {
		return nil, errors.New("amount is negative")
	}

	param := orderParam{OrderType: "sell", Pair: "btc_jpy", Rate: rate, Amount: amount}

	return cli.newOrder(ctx, &param)
}

func (cli *Client) OrderToMarketBuy(ctx context.Context, yen int) (*Order, error) {
	if yen < 0 {
		return nil, errors.New("yen is negative")
	}

	param := orderParam{OrderType: "market_buy", Pair: "btc_jpy", MarketBuyAmount: yen}

	return cli.newOrder(ctx, &param)
}

func (cli *Client) OrderToMarketSell(ctx context.Context, amount float64) (*Order, error) {
	if amount < 0 {
		return nil, errors.New("amount is negative")
	}

	param := orderParam{OrderType: "market_sell", Pair: "btc_jpy", Amount: amount}

	return cli.newOrder(ctx, &param)
}

func (cli *Client) GetBalance(ctx context.Context) (*Balance, error) {
	req, err := cli.newRequest(ctx, http.MethodGet, "/api/accounts/balance", []byte(""))
	if err != nil {
		return nil, err
	}

	res, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var balance Balance
	err = decodeBody(res, &balance)
	if err != nil {
		return nil, err
	}

	return &balance, nil
}

func (cli *Client) newRequest(ctx context.Context, method, endpoint string, body []byte) (*http.Request, error) {
	u := *cli.BaseURL
	u.Path = path.Join(cli.BaseURL.Path, endpoint)
	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	headers := getHeaders(cli.accessKey, cli.secretAccessKey, cli.BaseURL.String()+endpoint, string(body))
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req = req.WithContext(ctx)

	return req, nil
}

func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}

func encodeBody(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}

func getHeaders(key, secret, uri, body string) map[string]string {
	currentTime := time.Now().UTC().Unix()
	nonce := strconv.Itoa(int(currentTime))
	message := nonce + uri + body
	signature := calcHmac256(message, secret)
	return map[string]string{
		"Content-Type":     "application/json",
		"ACCESS-KEY":       key,
		"ACCESS-NONCE":     nonce,
		"ACCESS-SIGNATURE": signature,
		"User-Agent":       userAgent,
	}
}

func calcHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
