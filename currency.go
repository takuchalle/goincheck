// Package goincheck is client for Coincheck Exchange API
package goincheck

// Pair shows exchange pairs.
type Pair string

const (
	// BtcJpy is pair of BTC(Bitcoint) and Japanese yen.
	BtcJpy Pair = "btc_jpy"

	// EthJpy is pair of ETH(Ethereum) and Japanese yen.
	EthJpy Pair = "eth_jpy"

	// EtcJpy is pair of ETC(Ethereum Classic) and Japanese yen.
	EtcJpy Pair = "etc_jpy"

	// DapJpy is pair of DAP and Japanese yen.
	DapJpy Pair = "dao_jpy"

	// LskJpy is pair of Lsk(Lisk) and Japanese yen.
	LskJpy Pair = "lsk_jpy"

	// FctJpy is pair of FCT(Factom) and Japanese yen.
	FctJpy Pair = "fct_jpy"

	// XmrJpy is pair of XMR(Manero) and Japanese yen.
	XmrJpy Pair = "xmr_jpy"

	// RepJpy is pair of REP(Augur) and Japanese yen.
	RepJpy Pair = "rep_jpy"

	// XrpJpy is pair of XRP(Ripple) and Japanese yen.
	XrpJpy Pair = "xrp_jpy"

	// ZecJpy is pair of ZEC(Zcash) and Japanese yen.
	ZecJpy Pair = "zec_jpy"

	// EthBtc is pair of ETH(Ethereum) and BTC(Bitcoin).
	EthBtc Pair = "eth_btc"

	// EtcBtc is pair of ETC(Ethereum Classic) and BTC(Bitcoin).
	EtcBtc Pair = "etc_btc"

	// LskBtc is pair of LSK(Lisk) and BTC(Bitcoin).
	LskBtc Pair = "lsk_btc"

	// FctBtc is pair of FCT(Factom) and BTC(Bitcoin).
	FctBtc Pair = "fct_btc"

	// XmrBtc is pair of XMR(Monero) and BTC(Bitcoin).
	XmrBtc Pair = "xmr_btc"

	// RerBtc is pair of REP(Augur) and BTC(Bitcoin).
	RerBtc Pair = "rep_btc"

	// XrpBtc is pair of XRP(Ripple) and BTC(Bitcoin).
	XrpBtc Pair = "xrp_btc"

	// ZecBtc is pair of ZEC(Zcash) and BTC(Bitcoin).
	ZecBtc Pair = "zec_btc"
)
