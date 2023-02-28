package web

import "time"

type ResponseSenderMessage struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}
