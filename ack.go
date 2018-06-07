package service

import (
	"fmt"
	"time"
)

// Ack
type Ack struct {
	Version     int         `json:"ack_version"`
	Agent       string      `json:"agent"`
	Uuid        string      `json:"ack_uuid"`
	RequestUuid string      `json:"req_uuid"`
	DateTime    string      `json:"date_time"`
	Success     bool        `json:"success"`
	ServerCode  int         `json:"server_code"`
	Location    string      `json:"location"`
	PayloadType string      `json:"payload_type"`
	Payload     interface{} `json:"payload"`
	Duration    string      `json:"duration"`
	instTime    time.Time
}

// StartTimer
func (a *Ack) StartTimer() {
	a.instTime = time.Now()
}

// SetPayload
func (a *Ack) SetPayload(payload interface{}) {
	a.Duration = fmt.Sprintf("%s", time.Since(a.instTime))
	a.Payload = payload
}

// SetPayloadType
func (a *Ack) SetPayloadType(payloadType string) {
	a.PayloadType = payloadType
}
