package helper

import (
	"time"
	"udp-be/models"
)

func GetConnInitMessage() models.MetricsData {
	return models.MetricsData{
		Version:     1,
		MessageType: 1,
		SeqNum:      0,
		Payload:     []uint8{},
		TimeStamp:   uint64(time.Now().UnixNano()),
	}
}

func GetHeartBeatMessage() models.MetricsData {
	return models.MetricsData{
		Version:     1,
		MessageType: 2,
		SeqNum:      0,
		Payload:     []uint8{},
		TimeStamp:   uint64(time.Now().UnixNano()),
	}
}
