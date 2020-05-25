package data

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	HistorySuffix   = "/pricehistory?apikey="
	startDateSuffix = "&startDate="
	debugDate       = "1590154200000"
)

type HistoryPayload struct {
	Candles []struct {
		Open     float64 `json:"open"`
		High     float64 `json:"high"`
		Low      float64 `json:"low"`
		Close    float64 `json:"close"`
		Volume   int     `json:"volume"`
		Datetime int     `json:"datetime"`
	} `json:"candles"`
	Symbol string `json:"symbol"`
	Empty  bool   `json:"empty"`
}

type HistoryProvider struct {
	StartDate time.Time
	EndPoint  string
}

func NewHistoryProvider(apiKey, ticker string) *HistoryProvider {
	t := time.Unix(1589808600000, 0)
	endPoint := BaseUrl + ticker + HistorySuffix + apiKey + startDateSuffix + debugDate
	return &HistoryProvider{
		StartDate: t,
		EndPoint:  endPoint,
	}
}
func (h *HistoryProvider) GetData() *HistoryPayload {
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
