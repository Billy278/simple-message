package repository

import (
	"context"
	"database/sql"
	"simple-message/helper"
	"simple-message/model/domain"
)

type RepositoryMessageImpl struct {
}

func NewRepositoryMessageImpl() RepositoryMessage {
	return &RepositoryMessageImpl{}
}

func (repository_message *RepositoryMessageImpl) Create(ctx context.Context, DB *sql.DB, message domain.Massage) domain.Massage {
	sql := "Insert Into message(sender,date,message,receiver,sender_id) Values (?,?,?,?,?)"
	result, err := DB.ExecContext(ctx, sql, message.Sender, message.Date, message.Message, message.Receiver, message.Sender_id)
	helper.PanicIfError(err)
	id, _ := result.LastInsertId()
	message.Id = int(id)
	return message
}
func (repository_message *RepositoryMessageImpl) CreateTest(ctx context.Context, DB *sql.DB, message domain.Massage) domain.Massage {
	sql := "Insert Into message_test(sender,date,message,receiver,sender_id) Values (?,?,?,?,?)"
	result, err := DB.ExecContext(ctx, sql, message.Sender, message.Date, message.Message, message.Receiver, message.Sender_id)
	helper.PanicIfError(err)
	id, _ := result.LastInsertId()
	message.Id = int(id)
	return message
}
func (repository_message *RepositoryMessageImpl) SelectPartSender(ctx context.Context, DB *sql.DB, sender string, receiver string, offside int, limit int) []domain.Massage {
	sql := "Select message,date from message where receiver=? && sender=? LIMIT ? OFFSET ?"
	rows, err := DB.QueryContext(ctx, sql, receiver, sender, limit, offside)
	helper.PanicIfError(err)
	defer rows.Close()
	var messageAll []domain.Massage

	for rows.Next() {
		message := domain.Massage{}
		rows.Scan(&message.Message, &message.Date)
		messageAll = append(messageAll, message)
	}
	return messageAll
}
func (repository_message *RepositoryMessageImpl) SelectAllSenderWithLastMessage(ctx context.Context, DB *sql.DB, receiver string) []domain.Massage {
	//sql := "Select id,sender,date,message,receiver from message Where receiver=? && id > ?"
	sql := "select max(m.id), m.sender,m.message, m.receiver, count(*)-1 as jumlah from message m left join last_read lr on m.sender_id = lr.sender_id where m.id >= lr.message_id && m.receiver=? group by m.sender_id"
	rows, err := DB.QueryContext(ctx, sql, receiver)
	helper.PanicIfError(err)
	defer rows.Close()
	//Select *from message where receiver='bima' && id >3 limit 2 offset 1;
	var messageAll []domain.Massage
	for rows.Next() {
		message := domain.Massage{}
		//err := rows.Scan(&message.Id, &message.Sender, &message.Date, &message.Message, &message.Receiver)
		err := rows.Scan(&message.Id, &message.Sender, &message.Message, &message.Receiver, &message.Jumlah)
		helper.PanicIfError(err)
		messageAll = append(messageAll, message)
	}
	return messageAll
}
