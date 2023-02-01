package tradingalgorithm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/juhonamnam/trading-bot-project/logger"
)

type candle struct {
	OpeningPrice float64 `json:"opening_price"`
	HighPrice    float64 `json:"high_price"`
	LowPrice     float64 `json:"low_price"`
	TradePrice   float64 `json:"trade_price"`
	Timestamp    int64   `json:"timestamp"`
}

var httpClient = http.Client{Timeout: time.Duration(10) * time.Second}

func getCandles(market string) (*[]candle, error) {
	url := fmt.Sprintf("https://api.upbit.com/v1/candles/days?market=%s&count=21", market)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var candles []candle
	err = json.Unmarshal(body, &candles)

	if err != nil {
		return nil, err
	}

	logger.VBS.Debug.Printf("getCandles: %+v\n", candles)

	return &candles, nil
}
