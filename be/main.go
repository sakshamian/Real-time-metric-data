package main

import (
	"fmt"
	"time"
	"udp-be/db"
	"udp-be/handler"
	"udp-be/helper"
	messageserver "udp-be/message_server"
	"udp-be/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// connect to udp client
	messageserver.ConnectToUdp()

	// load env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	// connect to sql DB
	db.Connect()

	initMessage, err := helper.ConvertToBytes(helper.GetConnInitMessage())
	if err != nil {
		fmt.Println(err)
	}

	// connection init
	_, err = messageserver.UdpClient.Write(initMessage)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	fmt.Println("Connection init message sent to server:")

	// heartbeat message
	go func() {
		heartbeatMessage, err := helper.ConvertToBytes(helper.GetHeartBeatMessage())
		if err != nil {
			fmt.Println(err)
		}

		for {
			_, err := messageserver.UdpClient.Write(heartbeatMessage)

			if err != nil {
				fmt.Println("Error sending heartbeat:", err)
			} else {
				fmt.Println("Heartbeat sent to client.")
			}
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			receivedMessage := make([]byte, 1024)
			messageserver.UdpClient.SetReadDeadline(time.Now().Add(1 * time.Second))
			n, _, err := messageserver.UdpClient.ReadFromUDP(receivedMessage)
			if err != nil {
				fmt.Println("Error reading from server:", err)
				continue
			}

			dataReceived, err := helper.ConvertToPackage(receivedMessage, n)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Received from server:", dataReceived)

			payloadReceived, err := helper.ConvertToMetricsPayload(dataReceived.Payload, n)
			if err != nil {
				fmt.Println(err)
			}

			// Write to DB
			handler.WriteToDatabase(payloadReceived)

			time.Sleep(1 * time.Second)
		}
	}()

	// metrics route
	router := gin.Default()

	router.Use(middleware.CorsMiddleware())

	router.GET("/metrics", handler.HandleMetrics)
	router.GET("/ws", handler.HandleConnections)

	go handler.SendUpdates()
	go handler.PollDatabase()

	router.Run(":8080")
}
