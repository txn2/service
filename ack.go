package service

// Ack
type Ack struct {
	Version     int         `json:"ack_version"`
	Uuid        string      `json:"ack_uuid"`
	RequestUuid string      `json:"req_uuid"`
	DateTime    string      `json:"date_time"`
	Success     bool        `json:"success"`
	ServerCode  int         `json:"server_code"`
	Location    string      `json:"location"`
	PayloadType string      `json:"payload_type"`
	Payload     interface{} `json:"payload"`
}

// SetPayload
func (a *Ack) SetPayload(payload interface{}) {
	a.Payload = payload
}

// SetPayloadType
func (a *Ack) SetPayloadType(payloadType string) {
	a.PayloadType = payloadType
}
