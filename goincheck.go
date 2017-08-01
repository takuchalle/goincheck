package goincheck

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const defalutBaseURL = "https://coincheck.com"

type Client struct {
	accessKey       string
	secretAccessKey string

	BaseURL *url.URL

	HTTPClient *http.Client
}

type Ticker struct {
	Last      float64 `json:"last"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
	Timestamp float64 `json:"timestamp"`
}

type Trade struct {
	ID        int     `json:"id"`
	Amount    string  `json:"amount"`
	Rate      float64 `json:"rate"`
	OrderType string  `json:"order_type"`
	CreatedAt string  `json:"created_at"`
}

type OrderBook struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

func NewClient(key, secretKey string) (*Client, error) {
	if key == "" || secretKey == "" {
		return &Client{}, errors.New("key is missing")
	}

	baseurl, _ := url.Parse(defalutBaseURL)
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	cli := &Client{accessKey: key, secretAccessKey: secretKey, BaseURL: baseurl, HTTPClient: client}

	return cli, nil
}

func (cli *Client) GetTicker() (ticker Ticker, err error) {
	cli.httpGetRequest("/api/ticker", &ticker)

	return ticker, nil
}

func (cli *Client) GetTrade() (trades []Trade, err error) {
	cli.httpGetRequest("/api/trades", &trades)

	return trades, nil
}

func (cli *Client) GetOrderBook() (orderbook OrderBook, err error) {
	cli.httpGetRequest("/api/order_books", &orderbook)

	return orderbook, nil
}

func (cli *Client) httpGetRequest(endpoint string, data interface{}) error {
	resp, _ := http.Get(cli.BaseURL.String() + endpoint)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, data)

	return nil
}
