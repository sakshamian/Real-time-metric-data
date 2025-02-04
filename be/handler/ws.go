package handler

import (
	"log"
	"net/http"
	"sync"
	"time"
	"udp-be/db"
	"udp-be/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.MetricsPayload)
var mutex = &sync.Mutex{}
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
	}
}

func SendUpdates() {
	for {
		metric := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(metric)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func PollDatabase() {
	var lastTimestamp time.Time

	for {
		var latestMetric models.MetricsPayload
		err := db.DB.Table("message").Order("created_at DESC").First(&latestMetric).Error
		if err == nil && latestMetric.CreatedAt.After(lastTimestamp) {
			lastTimestamp = latestMetric.CreatedAt
			broadcast <- latestMetric
		}
		time.Sleep(1 * time.Second)
	}
}
