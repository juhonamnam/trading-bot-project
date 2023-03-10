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

var running = true
var wsDisconnect = make(chan struct{})

func Start() {
	for {
		if running {
			initializeWebsocket()
		} else {
			return
		}
	}
}

func Stop() {
	running = false
	close(wsDisconnect)
}

func initializeWebsocket() {
	defer func() {
		if err := recover(); err != nil {
			logger.VBS.Error.Println(err)
		}
	}()

	logger.VBS.Info.Println("WS Connecting")
	c, res, err := websocket.DefaultDialer.Dial("wss://api.upbit.com/websocket/v1", nil)
	if err != nil {
		logger.VBS.Error.Println("WS Dial:", err)
		return
	}
	defer func() {
		c.Close()
		logger.VBS.Info.Println("WS Disconnected")
	}()
	logger.VBS.Info.Println("WS Connected")
	logger.VBS.Debug.Printf("%+v\n", res)

	readerDisconnect := make(chan struct{})

	go func() {
		// Reader
		defer close(readerDisconnect)
		for {
			ticker := tickerResponse{}
			err := c.ReadJSON(&ticker)
			if err != nil {
				logger.VBS.Error.Println("WS Read:", err)
				return
			}
			logger.VBS.Debug.Printf("%+v\n", ticker)
			processVB(&ticker)
		}
	}()

	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-wsDisconnect:
			return

		case <-readerDisconnect:
			return

		case <-ticker.C:
			// Writter
			err := c.WriteJSON([]interface{}{
				ticketField{Ticket: "juhonamnam-trading-bot-project"},
				typeField{Type: "ticker", Codes: []string{"KRW-BTC", "KRW-ETH", "KRW-EOS", "KRW-BCH"}, IsOnlySnapshot: true},
				formatField{Format: "SIMPLE"},
			})
			if err != nil {
				logger.VBS.Info.Println("WS Write:", err)
				return
			}
		}
	}

}
