package app

import (
	"simple-message/helper"

	"github.com/streadway/amqp"
)

func NewRabbitmq() (*amqp.Connection, *amqp.Channel) {
	// 1. Buat koneksi RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helper.PanicIfError(err)
	//defer conn.Close()

	// 2. Buka channel
	ch, err := conn.Channel()
	helper.PanicIfError(err)
	//defer ch.Close()

	// 3. Deklarasikan exchange
	err = ch.ExchangeDeclare("my_exchange", "direct", true, false, false, false, nil)
	helper.PanicIfError(err)

	//4.declatation queque
	q, err := ch.QueueDeclare("my_queue", true, false, false, false, nil)
	helper.PanicIfError(err)

	// 5. Bind queue ke exchange
	err = ch.QueueBind(q.Name, "MY_KEY", "my_exchange", false, nil)
	helper.PanicIfError(err)

	return conn, ch
}
