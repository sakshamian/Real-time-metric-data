package models

import "time"

type MetricsData struct {
	Version     uint8
	MessageType uint8
	SeqNum      uint32
	Payload     []uint8
	TimeStamp   uint64
}

type MetricsPayload struct {
	StatusCode   uint16    `json:"status_code"`
	ResponseTime float64   `json:"response_time"`
	Method       uint8     `json:"method"`
	Route        string    `json:"route"`
	CreatedAt    time.Time `json:"created_at"`
}

type MetricsTable struct {
	ID          uint `gorm:"primary_key"`
	Version     uint8
	MessageType uint8
	SeqNum      uint32
	Payload     string
	TimeStamp   uint64
}

type Message struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	StatusCode   int       `gorm:"not null"`
	ResponseTime float64   `gorm:"not null"`
	Method       int       `gorm:"not null"`
	Route        string    `gorm:"type:text;not null"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (Message) TableName() string {
	return "message"
}
