package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-message/helper"
	"simple-message/model/domain"
)

type RepositoryUserImpl struct {
}

func NewRepositoryUserImpl() RepositoryUser {
	return &RepositoryUserImpl{}
}

func (repository_user *RepositoryUserImpl) Register(ctx context.Context, DB *sql.DB, user domain.User) {
	sql := "Insert Into user(name,username,password) Values(?,?,?)"
	result, err := DB.ExecContext(ctx, sql, user.Name, user.Username, user.Password)
	helper.PanicIfError(err)
	id, _ := result.LastInsertId()
	user.Id = int(id)

}
func (repository_user *RepositoryUserImpl) FindByUser(ctx context.Context, DB *sql.DB, username string) (domain.User, error) {
	sql := "Select id,name,username, password from user Where username=?"
	rows, _ := DB.QueryContext(ctx, sql, username)
	defer rows.Close()
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("NOT FOUND")
	}
}
