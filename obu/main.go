package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/Admiral-Simo/toll-calculator/types"
	"github.com/gorilla/websocket"
)

const (
	sendInterval = time.Second
	wsEndpoint   = "ws://127.0.0.1:30000/ws"
	latRange     = 0.5
	longRange    = 0.5
)

func genLatLong() (float64, float64) {
	f := rand.Float64()*longRange - longRange/2 // means range from -latRange to latRange
	s := rand.Float64()*longRange - longRange/2 // means range from -longRange to longRange
	return f, s
}

func sendFromObu(id int, conn *websocket.Conn) {
	lat, long := genLatLong()
	obu := &types.OBUData{
		OBUID: id,
		Lat:   lat,
		Long:  long,
	}
	// send the obu right here
	if err := conn.WriteJSON(obu); err != nil {
		log.Fatal(err)
	}
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	obuIDS := generateOBUIDS(20)
	for {
		for _, id := range obuIDS {
			sendFromObu(id, conn)
		}
		time.Sleep(sendInterval)
	}
}
