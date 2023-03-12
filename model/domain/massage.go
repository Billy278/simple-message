package domain

import "time"

type Massage struct {
	Id        int
	Sender    string
	Date      time.Time
	Message   string
	Receiver  string
	Sender_id int
	Jumlah    int
}
