package web

import "time"

type ResponseMessage struct {
	Id       int       `json:"id"`
	Sender   string    `json:"sender"`
	Date     time.Time `json:"date"`
	Message  string    `json:"message"`
	Receiver string    `json:"receiver"`
	Jumlah   int       `json:"jumlah"`
}
