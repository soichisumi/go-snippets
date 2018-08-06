package library

import (
	"net/url"
	"github.com/gorilla/websocket"
	"log"
	"fmt"
	"time"
)

type CallApi struct {
	ID      int    `json:"id"`
	Command string `json:"command"`
}

type GetTx struct {
	ID      int    `json:"id"`
	Command string `json:"command"`
	Transaction		string 	`json:"transaction"`
}

type ProposedTx struct {
	ID      int      `json:"id"`
	Command string   `json:"command"`
	Streams []string `json:"streams"`
}

// gorilla-websocket
func RunGori() {
	u := url.URL{Scheme: "ws", Host: "rippled-ws-testnet-develop.ginco.company:6006", Path: "/"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("fatal: %s", err.Error())
	}
	fmt.Println("connected.")
	defer conn.Close()

	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", msg)
		}
	}()

	// tick ledger_current
	// test tx : 9BC76BB6D1824028D2A3DCBE030EF835195B8041298414677C05FCD53F4AD55C
	//for {
	//	select {
	//	case <-done:
	//		return
	//	case <-ticker.C:
	//		conn.WriteJSON(&CallApi{
	//			ID:      2,
	//			Command: "ledger_current",
	//		})
	//	}
	//}

	// tx
	//conn.WriteJSON(&GetTx{
	//	ID: 		0,
	//	Command: 	"tx",
	//	Transaction: "9BC76BB6D1824028D2A3DCBE030EF835195B8041298414677C05FCD53F4AD55C",
	//})
	//conn.WriteJSON(&GetTx{
	//	ID: 		1,
	//	Command: 	"ledger_current",
	//})
	// subscribe
	// Streams: 	[]string{"transactions_proposed", "ledger"},
	conn.WriteJSON(&ProposedTx{
		ID: 		0,
		Command: 	"subscribe",
		Streams: 	[]string{"transactions_proposed"},
	})
	<-done
}
