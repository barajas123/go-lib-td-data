package data

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)
const(
	// TD endpoint
	BaseUrl = "https://api.tdameritrade.com/v1/marketdata/"
	QuoteSuffix  = "/data?apikey="
)



type QuoteProvider struct{
	Endpoint string
}

func NewQuoteProvider(apiKey,ticker string) *QuoteProvider{
	return &QuoteProvider{
		Endpoint: BaseUrl + ticker + QuoteSuffix + apiKey,
	}
}

func (q *QuoteProvider) GetData() *Quote {
	var quote Quote
	resp, err  := http.Get(q.Endpoint)
	if err != nil{
		logrus.Error(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		logrus.Error("ERROR reading body : %d",err)
		return nil
	}

	err = json.Unmarshal(body,&quote)
	if err != nil{
		logrus.Error("ERROR unmarshal data")
		return nil
	}

	return &quote
}


// TODO add support for different tickers, not just SPXL
type SimpleTDQuote struct {
	BidPrice float64 `json:"bidPrice"`
	BidSize  int `json:"bidSize"`
	AskPrice float64 `json:"askPrice"`
	AskSize  int `json:"askSize"`
	TotalVolume   int `json:"totalVolume"`
}


type Quote struct {
	Symbol struct{
		AssetType                          string  `json:"assetType"`
		AssetMainType                      string  `json:"assetMainType"`
		Cusip                              string  `json:"cusip"`
		AssetSubType                       string  `json:"assetSubType"`
		Symbol                             string  `json:"symbol"`
		Description                        string  `json:"description"`
		BidPrice                           float64 `json:"bidPrice"`
		BidSize                            float64 `json:"bidSize"`
		BidID                              string  `json:"bidId"`
		AskPrice                           float64 `json:"askPrice"`
		AskSize                            float64 `json:"askSize"`
		AskID                              string  `json:"askId"`
		LastPrice                          float64 `json:"lastPrice"`
		LastSize                           float64 `json:"lastSize"`
		LastID                             string  `json:"lastId"`
		OpenPrice                          float64 `json:"openPrice"`
		HighPrice                          float64 `json:"highPrice"`
		LowPrice                           float64 `json:"lowPrice"`
		BidTick                            string  `json:"bidTick"`
		ClosePrice                         float64 `json:"closePrice"`
		NetChange                          float64 `json:"netChange"`
		TotalVolume                        float64 `json:"totalVolume"`
		QuoteTimeInLong                    int64   `json:"quoteTimeInLong"`
		TradeTimeInLong                    int64   `json:"tradeTimeInLong"`
		Mark                               float64 `json:"mark"`
		Exchange                           string  `json:"exchange"`
		ExchangeName                       string  `json:"exchangeName"`
		Marginable                         bool    `json:"marginable"`
		Shortable                          bool    `json:"shortable"`
		Volatility                         float64 `json:"volatility"`
		Digits                             int     `json:"digits"`
		Five2WkHigh                        float64 `json:"52WkHigh"`
		Five2WkLow                         float64 `json:"52WkLow"`
		NAV                                float64 `json:"nAV"`
		PeRatio                            float64 `json:"peRatio"`
		DivAmount                          float64 `json:"divAmount"`
		DivYield                           float64 `json:"divYield"`
		DivDate                            string  `json:"divDate"`
		SecurityStatus                     string  `json:"securityStatus"`
		RegularMarketLastPrice             float64 `json:"regularMarketLastPrice"`
		RegularMarketLastSize              int     `json:"regularMarketLastSize"`
		RegularMarketNetChange             float64 `json:"regularMarketNetChange"`
		RegularMarketTradeTimeInLong       int64   `json:"regularMarketTradeTimeInLong"`
		NetPercentChangeInDouble           float64 `json:"netPercentChangeInDouble"`
		MarkChangeInDouble                 float64 `json:"markChangeInDouble"`
		MarkPercentChangeInDouble          float64 `json:"markPercentChangeInDouble"`
		RegularMarketPercentChangeInDouble float64 `json:"regularMarketPercentChangeInDouble"`
		Delayed                            bool    `json:"delayed"`
	} `json:"SPXL"`

}

