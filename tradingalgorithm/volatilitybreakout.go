package tradingalgorithm

import (
	"fmt"
	"math"
	"time"

	"github.com/juhonamnam/trading-bot-project/env"
	"github.com/juhonamnam/trading-bot-project/logger"
)

var resetTimestamp int64 = time.Now().UnixMilli() / 86400000 * 86400000
var targetPrices = map[string]float64{}
var buyPrices = map[string]float64{}
var haltFlag = false

func processVB(ticker *tickerResponse) {
	if haltFlag {
		return
	}

	if ticker.Timestamp > resetTimestamp {
		logger.VBS.Info.Println("Reset Time")
		haltFlag = true
		defer func() { haltFlag = false }()
		for _, m := range markets {
			setTargetPrice(m)
			time.Sleep(time.Duration(1) * time.Second)
		}

		resetTimestamp += 86400000
		// New Target Price
	}

	if buyPrices[ticker.Code] != 0 {
		return
	}

	if ticker.TradingPrice >= targetPrices[ticker.Code] {
		// Buy
		logger.VBS.Info.Printf("Buy %s, Current Price: %.0f\n", ticker.Code, ticker.TradingPrice)
		buyMessage := fmt.Sprintf("<u><i>%s 매수!!</i></u>\n현재가: %.0f", ticker.Code, ticker.TradingPrice)
		sendMessage(buyMessage, env.VBSChatId)
		buyPrices[ticker.Code] = ticker.TradingPrice
	}
}

func setTargetPrice(market string) {
	res, err := getCandles(market)

	pfx := fmt.Sprintf("setTargetPrice(%s):", market)
	if err != nil {
		logger.VBS.Error.Println(pfx, err)
		time.Sleep(time.Duration(10) * time.Second)
		setTargetPrice(market)
		return
	}
	if (*res)[0].Timestamp < resetTimestamp {
		logger.VBS.Error.Println(pfx, "Recent Data Not Out Yet")
		time.Sleep(time.Duration(10) * time.Second)
		setTargetPrice(market)
		return
	}

	var volatility float64 = 0
	for _, c := range (*res)[1:] {
		volatility += math.Abs(c.TradePrice-c.OpeningPrice) / (c.HighPrice - c.LowPrice)
	}
	k := 1 - volatility/20

	targetPrices[market] = (*res)[0].OpeningPrice + k*((*res)[1].HighPrice-(*res)[1].LowPrice)

	if buyPrices[market] != 0 {
		// Sell
		interest := 100 * ((*res)[0].TradePrice - buyPrices[market]) / buyPrices[market]
		logger.VBS.Info.Printf("Sell %s, Interest: %.2f%%\n", market, interest)
		sellMessage := fmt.Sprintf("<u><i>%s 매도!!</i></u>\n매수가: %.0f\n매도가: %.0f\n수익률: <b>%.2f%%</b>", market, buyPrices[market], (*res)[0].TradePrice, interest)
		sendMessage(sellMessage, env.VBSChatId)
		buyPrices[market] = 0
	}
	logger.VBS.Info.Printf(pfx+"Target Price %.0f, k-value: %f", targetPrices[market], k)
}
