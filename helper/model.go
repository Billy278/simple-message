package helper

import (
	"simple-message/model/domain"
	"simple-message/model/web"
)

func ResponseMessage(message domain.Massage) web.ResponseMessage {
	return web.ResponseMessage{
		Id:       message.Id,
		Sender:   message.Sender,
		Date:     message.Date,
		Message:  message.Message,
		Receiver: message.Receiver,
		Jumlah:   message.Jumlah,
	}
}

func ResponsesMessage(Allmessage []domain.Massage) []web.ResponseMessage {
	var responsemessage []web.ResponseMessage
	for _, message := range Allmessage {
		responsemessage = append(responsemessage, ResponseMessage(message))
	}
	return responsemessage
}

func ResponseSenderMessage(message domain.Massage) web.ResponseSenderMessage {
	return web.ResponseSenderMessage{
		Message: message.Message,
		Date:    message.Date,
	}
}

func ResponsesSenderMessage(Allmessage []domain.Massage) []web.ResponseSenderMessage {
	var responsemessage []web.ResponseSenderMessage
	for _, message := range Allmessage {
		responsemessage = append(responsemessage, ResponseSenderMessage(message))
	}
	return responsemessage
}
