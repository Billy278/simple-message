package domain

import "time"

type Massage struct {
	Id       int
	Sender   string
	Date     time.Time
	Message  string
	Receiver string
	Jumlah   int
}
