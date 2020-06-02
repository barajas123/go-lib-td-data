package data

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	HistorySuffix   = "/pricehistory?apikey="
	periodType = "&periodType=ytd"
	frequencyType = "&frequencyType=daily"
	startDateSuffix = "&startDate="
	start = "1578321000000"
)

type HistoryProvider struct {
	Ticker string
	APIKey string
	StartDate string
	EndPoint  string
}

func NewHistoryProvider(apiKey, ticker string) *HistoryProvider {
	t := start // set to 01/06/2020
	endPoint := BaseUrl + ticker + HistorySuffix + apiKey + periodType + frequencyType + startDateSuffix + start
	return &HistoryProvider{
		Ticker: ticker,
		APIKey: apiKey,
		StartDate: t,
		EndPoint:  endPoint,
	}
}

// ChangeStartDate changes the start date for the api call , takes in unix time stamp since epoch
func (hp *HistoryProvider)ChangeStartDate(newTime string){
	hp.StartDate = newTime
	hp.EndPoint = BaseUrl + hp.Ticker + HistorySuffix + hp.APIKey + periodType + frequencyType + startDateSuffix + start
}

//GetData returns HistoryPayload type
func (h *HistoryProvider) GetData() *HistoryPayload {
	if h.EndPoint != "https://api.tdameritrade.com/v1/marketdata/SPXL/pricehistory?apikey=LUL87NNF5RJ3CONF4N3P1MNUUZSOTGNY&periodType=ytd&frequencyType=daily&startDate=1578321000000"{
		logrus.Infof("\nHAVE: %s\nWANT: %s",h.EndPoint,"https://api.tdameritrade.com/v1/marketdata/SPXL/pricehistory?apikey=LUL87NNF5RJ3CONF4N3P1MNUUZSOTGNY&periodType=ytd&frequencyType=daily&startDate=1578321000000")
	}
	var payload HistoryPayload
	resp, err := http.Get(h.EndPoint)
	if err != nil {
		logrus.Errorf("ERROR hitting endpoint %s", err.Error())
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("ERROR reading response body")
		return nil
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		logrus.Errorf("ERROR unmarshal data", err)
	}
	return &payload
}


type HistoryPayload struct{
	Candles []struct {
		Open     float64 `json:"open"`
		High     float64 `json:"high"`
		Low      float64 `json:"low"`
		Close    float64 `json:"close"`
		Volume   int `json:"volume"`
		Datetime int `json:"datetime"`
	} `json:"candles"`
	Symbol string `json:"symbol"`
	Empty bool `json:"empty"`
}
