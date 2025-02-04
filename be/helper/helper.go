package helper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"udp-be/models"
)

func ConvertToBytes(data models.MetricsData) ([]byte, error) {
	// func (p *Package) ToBytes() ([]byte, error) {
	var buffer bytes.Buffer

	// Write fields in BigEndian format
	if err := binary.Write(&buffer, binary.BigEndian, data.Version); err != nil {
		return nil, err
	}
	if err := binary.Write(&buffer, binary.BigEndian, data.MessageType); err != nil {
		return nil, err
	}
	if err := binary.Write(&buffer, binary.BigEndian, data.SeqNum); err != nil {
		return nil, err
	}
	if err := binary.Write(&buffer, binary.BigEndian, data.TimeStamp); err != nil {
		return nil, err
	}
	if _, err := buffer.Write(data.Payload); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func ConvertToPackage(bytes []byte, n int) (models.MetricsData, error) {
	if len(bytes) < 14 {
		return models.MetricsData{}, fmt.Errorf("invalid packet: insufficient length")
	}

	version := bytes[0]
	messageType := bytes[1]
	seqNum := binary.BigEndian.Uint32(bytes[2:6])
	timestamp := binary.BigEndian.Uint64(bytes[6:14])
	payload := bytes[14:n]

	return models.MetricsData{
		Version:     version,
		MessageType: messageType,
		SeqNum:      seqNum,
		TimeStamp:   timestamp,
		Payload:     payload,
	}, nil
}

func ConvertToMetricsPayload(bytes []byte, n int) (models.MetricsPayload, error) {
	if len(bytes) < 11 {
		return models.MetricsPayload{}, fmt.Errorf("invalid packet: insufficient length")
	}
	status_code := bytes[0:2]
	response_time := bytes[2:10]
	method := bytes[10]
	route := bytes[11:n]

	return models.MetricsPayload{
		StatusCode:   binary.BigEndian.Uint16(status_code),
		ResponseTime: math.Float64frombits(binary.BigEndian.Uint64(response_time)),
		Method:       method,
		Route:        string(route),
	}, nil
}
