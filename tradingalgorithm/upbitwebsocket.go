package tradingalgorithm

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/juhonamnam/trading-bot-project/logger"
)

type ticketField struct {
	Ticket string `json:"ticket"`
}

type typeField struct {
	Type           string   `json:"type"`
	Codes          []string `json:"codes"`
	IsOnlySnapshot bool     `json:"isOnlySnapshot,omitempty"`
	IsOnlyRealtime bool     `json:"isOnlyRealtime,omitempty"`
}

type formatField struct {
	Format string `json:"format"`
}

type tickerResponse struct {
	Code         string  `json:"cd"`
	TradingPrice float64 `json:"tp"`
	Timestamp    int64   `json:"tms"`
}

func Start() {
	for {
		initializeWebsocket()
	}
}

func initializeWebsocket() {
	defer func() {
		if err := recover(); err != nil {
			logger.VBS.Error(err)
		}
	}()

	c, _, err := websocket.DefaultDialer.Dial("wss://api.upbit.com/websocket/v1", nil)
	if err != nil {
		logger.VBS.Error("WS Dial:", err)
		return
	}
	defer c.Close()

	disconnect := make(chan struct{})

	go func() {
		defer close(disconnect)
		for {
			ticker := tickerResponse{}
			err := c.ReadJSON(&ticker)
			if err != nil {
				logger.VBS.Error("WS Read:", err)
				return
			}
			processVB(&ticker)
		}
	}()

	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-disconnect:
			return

		case <-ticker.C:
			err := c.WriteJSON([]interface{}{
				ticketField{Ticket: "juhonamnam-trading-bot-project"},
				typeField{Type: "ticker", Codes: []string{"KRW-BTC", "KRW-ETH", "KRW-EOS", "KRW-BCH"}, IsOnlySnapshot: true},
				formatField{Format: "SIMPLE"},
			})
			if err != nil {
				logger.VBS.Info("WS Write:", err)
				return
			}
		}
	}

}
