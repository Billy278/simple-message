package web

import "time"

type RequestMessage struct {
	Sender   string    `json:"sender"`
	Date     time.Time `validate:"required" json:"date"`
	Message  string    `json:"message"`
	Receiver string    `validate:"required" json:"receiver"`
}
